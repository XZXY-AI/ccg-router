package adapter

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/XZXY-AI/ccg-router/internal/normal"
	"github.com/stretchr/testify/require"
)

func TestParseOpenAIRequest_MinimalChatCompletion(t *testing.T) {
	body := []byte(`{
	  "model": "gpt-5",
	  "max_tokens": 256,
	  "messages":[
	    {"role":"system","content":"be concise"},
	    {"role":"user","content":"hello"}
	  ]
	}`)
	req := httptest.NewRequest("POST", "/v1/chat/completions",
		bytes.NewReader(body))
	nr, err := ParseOpenAIRequest(req)
	require.NoError(t, err)
	require.Equal(t, normal.SourceCodex, nr.SourceCLI)
	require.Equal(t, "gpt-5", nr.Model)
	require.Equal(t, 256, nr.MaxTokens)
	require.Len(t, nr.Messages, 2)
	require.Equal(t, "system", nr.Messages[0].Role)
}

func TestParseOpenAIRequest_DefaultsMaxTokens(t *testing.T) {
	body := []byte(`{
	  "model": "gpt-5",
	  "messages":[{"role":"user","content":"x"}]
	}`)
	req := httptest.NewRequest("POST", "/v1/chat/completions",
		bytes.NewReader(body))
	nr, err := ParseOpenAIRequest(req)
	require.NoError(t, err)
	require.Equal(t, 4096, nr.MaxTokens, "defaults when absent")
}
