package adapter

import (
	"encoding/json"
	"time"

	"github.com/XZXY-AI/ccg-router/internal/normal"
)

func FormatOpenAIResponse(r normal.NormalResponse) ([]byte, error) {
	if len(r.Raw) > 0 {
		return r.Raw, nil
	}
	out := map[string]any{
		"id":     "ccg-" + time.Now().UTC().Format("20060102T150405.000000000"),
		"object": "chat.completion",
		"model":  r.Model,
		"choices": []map[string]any{{
			"index": 0,
			"message": map[string]any{
				"role":    "assistant",
				"content": r.Content,
			},
			"finish_reason": r.FinishReason,
		}},
		"usage": map[string]any{
			"prompt_tokens":     r.InputTokens,
			"completion_tokens": r.OutputTokens,
			"total_tokens":      r.InputTokens + r.OutputTokens,
		},
	}
	return json.Marshal(out)
}
