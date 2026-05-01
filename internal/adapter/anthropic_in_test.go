package adapter

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/stretchr/testify/require"
)

func TestParseAnthropicRequest_MinimalMessage(t *testing.T) {
	body := []byte(`{
	  "model": "claude-sonnet-4-7",
	  "max_tokens": 512,
	  "messages": [{"role":"user","content":"hello"}]
	}`)
	req := httptest.NewRequest("POST", "/v1/messages",
		bytes.NewReader(body))

	nr, err := ParseAnthropicRequest(req)
	require.NoError(t, err)
	require.Equal(t, normal.SourceClaudeCode, nr.SourceCLI)
	require.Equal(t, "claude-sonnet-4-7", nr.Model)
	require.Equal(t, 512, nr.MaxTokens)
	require.Len(t, nr.Messages, 1)
	require.Equal(t, "user", nr.Messages[0].Role)
	require.Contains(t, nr.Messages[0].Content, "hello")
}

func TestParseAnthropicRequest_StreamFlag(t *testing.T) {
	body := []byte(`{
	  "model": "claude-sonnet-4-7",
	  "max_tokens": 32,
	  "stream": true,
	  "messages": [{"role":"user","content":"x"}]
	}`)
	req := httptest.NewRequest("POST", "/v1/messages",
		bytes.NewReader(body))

	nr, err := ParseAnthropicRequest(req)
	require.NoError(t, err)
	require.True(t, nr.Stream)
}
