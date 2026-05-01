package normal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalRequest_ZeroValueIsInvalid(t *testing.T) {
	var r NormalRequest
	err := r.Validate()
	require.Error(t, err)
}

func TestNormalRequest_MinimalValid(t *testing.T) {
	r := NormalRequest{
		SourceCLI: SourceClaudeCode,
		Model:     "claude-sonnet-4-7",
		Messages: []Message{
			{Role: "user", Content: "hi"},
		},
		MaxTokens: 256,
	}
	require.NoError(t, r.Validate())
}

func TestNormalRequest_TagsParsedFromPrefix(t *testing.T) {
	r := NormalRequest{
		SourceCLI: SourceCodex,
		Model:     "gpt-5",
		Messages: []Message{
			{Role: "user", Content: "@long please summarize..."},
		},
		MaxTokens: 1024,
	}
	tags := r.InferTags()
	require.Contains(t, tags, "long")
}
