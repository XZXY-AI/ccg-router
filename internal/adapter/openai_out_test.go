package adapter

import (
	"encoding/json"
	"testing"

	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/stretchr/testify/require"
)

func TestFormatOpenAIResponse_UsesRawPassthroughWhenPresent(t *testing.T) {
	raw := []byte(`{"id":"x","object":"chat.completion","choices":[{"index":0,"message":{"role":"assistant","content":"hi"},"finish_reason":"stop"}]}`)
	resp := normal.NormalResponse{Raw: raw}
	out, err := FormatOpenAIResponse(resp)
	require.NoError(t, err)
	require.JSONEq(t, string(raw), string(out))
}

func TestFormatOpenAIResponse_Synthesizes(t *testing.T) {
	resp := normal.NormalResponse{
		Model:        "gpt-5",
		Content:      "hello",
		InputTokens:  3,
		OutputTokens: 2,
		FinishReason: "stop",
	}
	out, err := FormatOpenAIResponse(resp)
	require.NoError(t, err)
	var got map[string]any
	require.NoError(t, json.Unmarshal(out, &got))
	require.Equal(t, "chat.completion", got["object"])
}
