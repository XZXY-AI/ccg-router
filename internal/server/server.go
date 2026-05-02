// Package server wires the HTTP surface:
//
//	POST /v1/messages           -> Anthropic-compatible (Claude Code)
//	POST /v1/chat/completions   -> OpenAI-compatible    (Codex CLI)
//	GET  /api/usage/summary     -> usage summary (local panel)
//	GET  /api/usage/window      -> last-5h rolling window
//	GET  /healthz               -> liveness
package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ccg-labs/ccg-router/internal/adapter"
	"github.com/ccg-labs/ccg-router/internal/dispatch"
	"github.com/ccg-labs/ccg-router/internal/ledger"
	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/ccg-labs/ccg-router/internal/router"
	"github.com/ccg-labs/ccg-router/internal/ui"
	"github.com/ccg-labs/ccg-router/internal/upstream"
)

type Deps struct {
	Pool       *upstream.Pool
	Engine     *router.Engine
	Ledger     *ledger.Ledger
	Dispatcher *dispatch.Dispatcher // optional; default created if nil
}

type Server struct {
	deps Deps
}

func New(d Deps) *Server {
	if d.Dispatcher == nil {
		d.Dispatcher = dispatch.New()
	}
	return &Server{deps: d}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/messages", s.handleAnthropic)
	mux.HandleFunc("POST /v1/chat/completions", s.handleOpenAI)
	mux.HandleFunc("GET /api/usage/summary", s.handleUsageSummary)
	mux.HandleFunc("GET /api/usage/window", s.handleUsageWindow)
	mux.Handle("GET /ui/", ui.Handler())
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`ok`))
	})
	return mux
}

func (s *Server) handleAnthropic(w http.ResponseWriter, r *http.Request) {
	nr, err := adapter.ParseAnthropicRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if nr.Stream {
		http.Error(w, "streaming not supported in ccg-router v0.1; set stream=false or wait for v0.2", http.StatusNotImplemented)
		return
	}
	s.proxy(w, r, nr, "/v1/messages", "anthropic")
}

func (s *Server) handleOpenAI(w http.ResponseWriter, r *http.Request) {
	nr, err := adapter.ParseOpenAIRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if nr.Stream {
		http.Error(w, "streaming not supported in ccg-router v0.1; set stream=false or wait for v0.2", http.StatusNotImplemented)
		return
	}
	s.proxy(w, r, nr, "/v1/chat/completions", "openai")
}

func (s *Server) proxy(
	w http.ResponseWriter, r *http.Request,
	nr normal.NormalRequest, path string, protocol string,
) {
	pool := filterProtocol(s.deps.Pool.Enabled(), protocol)
	pick, _, err := s.deps.Engine.Pick(r.Context(), nr, pool, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	status, hdr, body, err := s.deps.Dispatcher.Do(r.Context(), pick, nr, path, r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// Best-effort ledger insert; never fail the response for ledger errors.
	_ = s.deps.Ledger.Record(r.Context(), ledger.Entry{
		Timestamp:  time.Now().UTC(),
		SourceCLI:  string(nr.SourceCLI),
		UpstreamID: pick.ID,
		Model:      nr.Model,
	})

	ct := hdr.Get("Content-Type")
	if ct != "" {
		w.Header().Set("Content-Type", ct)
	}
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func filterProtocol(pool []upstream.Upstream, protocol string) []upstream.Upstream {
	out := make([]upstream.Upstream, 0, len(pool))
	for _, u := range pool {
		if u.Protocol == protocol {
			out = append(out, u)
		}
	}
	return out
}

func (s *Server) handleUsageSummary(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC()
	sum, err := s.deps.Ledger.Summary(r.Context(), now.Add(-24*time.Hour), now)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(sum)
}

func (s *Server) handleUsageWindow(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC()
	entries, err := s.deps.Ledger.Window(r.Context(), now.Add(-5*time.Hour), now)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(entries)
}
