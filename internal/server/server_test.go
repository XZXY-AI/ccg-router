package server

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/ccg-labs/ccg-router/internal/config"
	"github.com/ccg-labs/ccg-router/internal/ledger"
	"github.com/ccg-labs/ccg-router/internal/router"
	"github.com/ccg-labs/ccg-router/internal/upstream"
	"github.com/stretchr/testify/require"
)

func TestServer_RoutesAnthropicRequestToUpstream(t *testing.T) {
	up := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"x","type":"message","role":"assistant","model":"claude-sonnet-4-7","content":[{"type":"text","text":"hi"}],"usage":{"input_tokens":3,"output_tokens":2}}`))
	}))
	defer up.Close()

	l, srv := testServer(t, up.URL)

	body := []byte(`{"model":"claude-sonnet-4-7","max_tokens":64,"messages":[{"role":"user","content":"hi"}]}`)
	before := time.Now().Add(-time.Second)
	resp, err := http.Post(srv.URL+"/v1/messages", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()
	out, _ := io.ReadAll(resp.Body)
	require.Equal(t, 200, resp.StatusCode, string(out))
	require.Contains(t, string(out), `"type":"message"`)

	entries, err := l.Window(context.Background(), before, time.Now().Add(time.Second))
	require.NoError(t, err)
	require.Len(t, entries, 1, "one ledger row should be written per upstream request")
	require.Equal(t, "claude-code", entries[0].SourceCLI)
	require.Equal(t, "ant", entries[0].UpstreamID)
}

func TestServer_StreamingReturns501(t *testing.T) {
	up := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("streaming requests must be rejected before upstream dispatch")
	}))
	defer up.Close()

	_, srv := testServer(t, up.URL)

	body := []byte(`{"model":"claude-sonnet-4-7","max_tokens":64,"stream":true,"messages":[{"role":"user","content":"hi"}]}`)
	resp, err := http.Post(srv.URL+"/v1/messages", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusNotImplemented, resp.StatusCode)
}

func testServer(t *testing.T, upstreamURL string) (*ledger.Ledger, *httptest.Server) {
	t.Helper()
	pool, err := upstream.NewPool(config.Config{
		Upstreams: []config.Upstream{{
			ID: "ant", Protocol: "anthropic", BaseURL: upstreamURL,
			AuthHeader: "x-api-key: test", Enabled: true,
		}},
	}, nil)
	require.NoError(t, err)

	eng, err := router.New("prefer-cheaper")
	require.NoError(t, err)

	l, err := ledger.Open(filepath.Join(t.TempDir(), "ledger.db"))
	require.NoError(t, err)
	t.Cleanup(func() { _ = l.Close() })

	s := New(Deps{Pool: pool, Engine: eng, Ledger: l})
	ts := httptest.NewServer(s.Handler())
	t.Cleanup(ts.Close)
	return l, ts
}
