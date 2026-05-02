// Package adapter converts wire protocols to/from the normal IR.
package adapter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/XZXY-AI/ccg-router/internal/normal"
)

type anthropicMessage struct {
	Role    string          `json:"role"`
	Content json.RawMessage `json:"content"`
}

type anthropicReq struct {
	Model     string             `json:"model"`
	MaxTokens int                `json:"max_tokens"`
	Stream    bool               `json:"stream"`
	Messages  []anthropicMessage `json:"messages"`
	System    string             `json:"system,omitempty"`
}

// ParseAnthropicRequest reads an Anthropic /v1/messages body into NormalRequest.
func ParseAnthropicRequest(r *http.Request) (normal.NormalRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return normal.NormalRequest{}, fmt.Errorf("read body: %w", err)
	}
	var a anthropicReq
	if err := json.Unmarshal(body, &a); err != nil {
		return normal.NormalRequest{}, fmt.Errorf("parse anthropic request: %w", err)
	}
	nr := normal.NormalRequest{
		SourceCLI: normal.SourceClaudeCode,
		Model:     a.Model,
		MaxTokens: a.MaxTokens,
		Stream:    a.Stream,
		Extra:     map[string][]byte{"raw": body},
	}
	if a.System != "" {
		nr.Messages = append(nr.Messages, normal.Message{
			Role:    "system",
			Content: a.System,
		})
	}
	for _, m := range a.Messages {
		// v0.1: best-effort text extraction; non-text parts preserved in Raw.
		var text string
		if len(m.Content) > 0 && m.Content[0] == '"' {
			_ = json.Unmarshal(m.Content, &text)
		} else {
			text = string(m.Content)
		}
		nr.Messages = append(nr.Messages, normal.Message{
			Role:    m.Role,
			Content: text,
			Raw:     m.Content,
		})
	}
	if err := nr.Validate(); err != nil {
		return normal.NormalRequest{}, err
	}
	return nr, nil
}
