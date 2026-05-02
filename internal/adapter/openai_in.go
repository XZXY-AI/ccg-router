package adapter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/XZXY-AI/ccg-router/internal/normal"
)

type openaiMessage struct {
	Role    string          `json:"role"`
	Content json.RawMessage `json:"content"`
}

type openaiReq struct {
	Model     string          `json:"model"`
	MaxTokens int             `json:"max_tokens"`
	Stream    bool            `json:"stream"`
	Messages  []openaiMessage `json:"messages"`
}

func ParseOpenAIRequest(r *http.Request) (normal.NormalRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return normal.NormalRequest{}, fmt.Errorf("read body: %w", err)
	}
	var o openaiReq
	if err := json.Unmarshal(body, &o); err != nil {
		return normal.NormalRequest{}, fmt.Errorf("parse openai request: %w", err)
	}
	if o.MaxTokens == 0 {
		o.MaxTokens = 4096
	}
	nr := normal.NormalRequest{
		SourceCLI: normal.SourceCodex,
		Model:     o.Model,
		MaxTokens: o.MaxTokens,
		Stream:    o.Stream,
		Extra:     map[string][]byte{"raw": body},
	}
	for _, m := range o.Messages {
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
