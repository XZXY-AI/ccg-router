package adapter

import (
	"encoding/json"
	"testing"

	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/stretchr/testify/require"
)

func TestFormatAnthropicResponse_UsesRawPassthroughWhenPresent(t *testing.T) {
	raw := []byte(`{"id":"x","type":"message","role":"assistant","content":[{"type":"text","text":"hi"}]}`)
	resp := normal.NormalResponse{Raw: raw}
	out, err := FormatAnthropicResponse(resp)
	require.NoError(t, err)
	require.JSONEq(t, string(raw), string(out))
}

func TestFormatAnthropicResponse_SynthesizesWhenRawMissing(t *testing.T) {
	resp := normal.NormalResponse{
		Model:        "claude-sonnet-4-7",
		Content:      "hello",
		InputTokens:  3,
		OutputTokens: 2,
		FinishReason: "end_turn",
	}
	out, err := FormatAnthropicResponse(resp)
	require.NoError(t, err)
	var got map[string]any
	require.NoError(t, json.Unmarshal(out, &got))
	require.Equal(t, "message", got["type"])
	require.Equal(t, "claude-sonnet-4-7", got["model"])
}
