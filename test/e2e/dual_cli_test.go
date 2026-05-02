package e2e

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/ccg-labs/ccg-router/internal/config"
	"github.com/ccg-labs/ccg-router/internal/ledger"
	"github.com/ccg-labs/ccg-router/internal/router"
	"github.com/ccg-labs/ccg-router/internal/server"
	"github.com/ccg-labs/ccg-router/internal/upstream"
	"github.com/stretchr/testify/require"
)

func TestDualCLI_EndToEnd(t *testing.T) {
	antMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"type":"message","model":"claude-sonnet-4-7","content":[{"type":"text","text":"ant"}],"usage":{"input_tokens":1,"output_tokens":1}}`))
	}))
	oaiMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"object":"chat.completion","choices":[{"message":{"role":"assistant","content":"oai"}}]}`))
	}))
	defer antMock.Close()
	defer oaiMock.Close()

	pool, err := upstream.NewPool(config.Config{
		Upstreams: []config.Upstream{
			{ID: "ant", Protocol: "anthropic", BaseURL: antMock.URL,
				AuthHeader: "x-api-key: x", Enabled: true},
			{ID: "oai", Protocol: "openai", BaseURL: oaiMock.URL,
				AuthHeader: "Authorization: Bearer y", Enabled: true},
		},
	}, nil)
	require.NoError(t, err)

	eng, err := router.New("prefer-cheaper")
	require.NoError(t, err)
	l, err := ledger.Open(filepath.Join(t.TempDir(), "ledger.db"))
	require.NoError(t, err)
	defer l.Close()

	s := httptest.NewServer(server.New(server.Deps{
		Pool: pool, Engine: eng, Ledger: l,
	}).Handler())
	defer s.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "POST", s.URL+"/v1/messages",
		bytes.NewReader([]byte(`{"model":"claude-sonnet-4-7","max_tokens":8,"messages":[{"role":"user","content":"hi"}]}`)))
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	b, _ := io.ReadAll(resp.Body)
	require.Contains(t, string(b), "ant", fmt.Sprintf("body=%s", string(b)))
	_ = resp.Body.Close()

	req, _ = http.NewRequestWithContext(ctx, "POST", s.URL+"/v1/chat/completions",
		bytes.NewReader([]byte(`{"model":"gpt-5","max_tokens":8,"messages":[{"role":"user","content":"hi"}]}`)))
	resp, err = http.DefaultClient.Do(req)
	require.NoError(t, err)
	b, _ = io.ReadAll(resp.Body)
	require.Contains(t, string(b), "oai")
	_ = resp.Body.Close()

	sum, err := l.Summary(ctx, time.Now().Add(-time.Hour), time.Now().Add(time.Hour))
	require.NoError(t, err)
	require.Equal(t, 2, sum.Requests)
}
