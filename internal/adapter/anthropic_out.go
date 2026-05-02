package adapter

import (
	"encoding/json"
	"time"

	"github.com/XZXY-AI/ccg-router/internal/normal"
)

func FormatAnthropicResponse(r normal.NormalResponse) ([]byte, error) {
	if len(r.Raw) > 0 {
		return r.Raw, nil
	}
	out := map[string]any{
		"id":    "ccg_" + time.Now().UTC().Format("20060102T150405.000000000"),
		"type":  "message",
		"role":  "assistant",
		"model": r.Model,
		"content": []map[string]any{
			{"type": "text", "text": r.Content},
		},
		"stop_reason": r.FinishReason,
		"usage": map[string]any{
			"input_tokens":  r.InputTokens,
			"output_tokens": r.OutputTokens,
		},
	}
	return json.Marshal(out)
}
