package dispatch

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/ccg-labs/ccg-router/internal/upstream"
	"github.com/stretchr/testify/require"
)

func TestDispatcher_ForwardsAndReturnsResponse(t *testing.T) {
	up := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer test", r.Header.Get("Authorization"))
		body, _ := io.ReadAll(r.Body)
		require.Contains(t, string(body), "hello")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"x","object":"chat.completion","choices":[]}`))
	}))
	defer up.Close()

	d := New()
	u := upstream.Upstream{
		ID: "t", Protocol: "openai", BaseURL: up.URL,
		ResolvedAuthHeader: "Authorization: Bearer test", Enabled: true,
	}
	nr := normal.NormalRequest{
		SourceCLI: normal.SourceCodex, Model: "gpt-5", MaxTokens: 32,
		Messages: []normal.Message{{Role: "user", Content: "hello"}},
		Extra:    map[string][]byte{"raw": []byte(`{"model":"gpt-5","messages":[{"role":"user","content":"hello"}]}`)},
	}
	status, hdr, body, err := d.Do(context.Background(), u, nr, "/v1/chat/completions", http.Header{})
	require.NoError(t, err)
	require.Equal(t, 200, status)
	require.Equal(t, "application/json", hdr.Get("Content-Type"))
	require.True(t, strings.Contains(string(body), "chat.completion"))
}

func TestDispatcher_RetriesOn5xxThenSucceeds(t *testing.T) {
	calls := 0
	up := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls < 2 {
			w.WriteHeader(502)
			return
		}
		_, _ = w.Write([]byte(`ok`))
	}))
	defer up.Close()
	d := New()
	u := upstream.Upstream{ID: "t", Protocol: "openai", BaseURL: up.URL,
		ResolvedAuthHeader: "Authorization: Bearer x", Enabled: true}
	nr := normal.NormalRequest{
		SourceCLI: normal.SourceCodex, Model: "m", MaxTokens: 8,
		Messages: []normal.Message{{Role: "user", Content: "x"}},
		Extra:    map[string][]byte{"raw": []byte(`{}`)},
	}
	status, _, body, err := d.Do(context.Background(), u, nr, "/v1/chat/completions", http.Header{})
	require.NoError(t, err)
	require.Equal(t, 200, status)
	require.Equal(t, "ok", string(body))
	require.GreaterOrEqual(t, calls, 2)
}

func TestDispatcher_ForwardsAnthropicVersion(t *testing.T) {
	var seen string
	up := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = r.Header.Get("anthropic-version")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer up.Close()
	d := New()
	u := upstream.Upstream{ID: "a", Protocol: "anthropic", BaseURL: up.URL,
		ResolvedAuthHeader: "x-api-key: k", Enabled: true}
	nr := normal.NormalRequest{
		SourceCLI: normal.SourceClaudeCode, Model: "claude-sonnet-4-7", MaxTokens: 8,
		Messages: []normal.Message{{Role: "user", Content: "x"}},
		Extra:    map[string][]byte{"raw": []byte(`{}`)},
	}
	in := http.Header{}
	in.Set("anthropic-version", "2023-06-01")
	in.Set("anthropic-beta", "tools-2024-04-04")
	_, _, _, err := d.Do(context.Background(), u, nr, "/v1/messages", in)
	require.NoError(t, err)
	require.Equal(t, "2023-06-01", seen)
}
