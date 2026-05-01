// Package normal defines the internal request/response IR that all
// inbound protocol adapters produce and all outbound adapters consume.
package normal

import (
	"errors"
	"regexp"
	"strings"
)

type SourceCLI string

const (
	SourceClaudeCode SourceCLI = "claude-code"
	SourceCodex      SourceCLI = "codex"
	SourceUnknown    SourceCLI = "unknown"
)

type Message struct {
	Role    string // "system" | "user" | "assistant" | "tool"
	Content string
	// Tool-call / multi-part payloads are carried as JSON in Raw for v0.1.
	Raw []byte
}

type NormalRequest struct {
	SourceCLI SourceCLI
	Model     string
	Messages  []Message
	MaxTokens int
	Stream    bool
	// Arbitrary passthrough fields (temperature, tools, tool_choice, ...)
	// kept as raw JSON so adapters can round-trip without lossy parsing.
	Extra map[string][]byte
}

type NormalResponse struct {
	Model        string
	Content      string
	InputTokens  int
	OutputTokens int
	FinishReason string
	Raw          []byte // full upstream body for adapters to reshape
}

var tagRe = regexp.MustCompile(`(?i)@([a-z]+)`)

func (r *NormalRequest) Validate() error {
	if r.SourceCLI == "" {
		return errors.New("source_cli is required")
	}
	if r.Model == "" {
		return errors.New("model is required")
	}
	if len(r.Messages) == 0 {
		return errors.New("at least one message is required")
	}
	if r.MaxTokens <= 0 {
		return errors.New("max_tokens must be positive")
	}
	return nil
}

func (r *NormalRequest) InferTags() []string {
	var tags []string
	for _, m := range r.Messages {
		if m.Role != "user" {
			continue
		}
		for _, match := range tagRe.FindAllStringSubmatch(m.Content, -1) {
			tags = append(tags, strings.ToLower(match[1]))
		}
	}
	return tags
}
