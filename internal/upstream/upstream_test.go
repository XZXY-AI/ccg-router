package upstream

import (
	"testing"

	"github.com/ccg-labs/ccg-router/internal/config"
	"github.com/stretchr/testify/require"
)

func TestPool_FromConfig(t *testing.T) {
	cfg := config.Config{
		Upstreams: []config.Upstream{
			{ID: "a", Protocol: "anthropic", BaseURL: "https://api.anthropic.com",
				AuthHeader: "x-api-key: ${ANTHROPIC_API_KEY}", Enabled: true},
			{ID: "o", Protocol: "openai", BaseURL: "https://api.openai.com",
				AuthHeader: "Authorization: Bearer ${OPENAI_API_KEY}", Enabled: false},
		},
	}
	p, err := NewPool(cfg, map[string]string{
		"ANTHROPIC_API_KEY": "sk-ant-xxx",
		"OPENAI_API_KEY":    "sk-openai-yyy",
	})
	require.NoError(t, err)

	a, ok := p.Get("a")
	require.True(t, ok)
	require.Equal(t, "x-api-key: sk-ant-xxx", a.ResolvedAuthHeader)

	// disabled upstream still resolvable but flagged
	o, _ := p.Get("o")
	require.False(t, o.Enabled)

	// Enabled() filters
	en := p.Enabled()
	require.Len(t, en, 1)
	require.Equal(t, "a", en[0].ID)
}
