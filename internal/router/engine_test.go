package router

import (
	"context"
	"testing"

	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/ccg-labs/ccg-router/internal/upstream"
	"github.com/stretchr/testify/require"
)

func mkPool(ids ...string) []upstream.Upstream {
	out := make([]upstream.Upstream, len(ids))
	for i, id := range ids {
		protocol := "anthropic"
		if id[0] == 'o' {
			protocol = "openai"
		}
		out[i] = upstream.Upstream{ID: id, Protocol: protocol, Enabled: true}
	}
	return out
}

func TestEngine_PreferCheaper_PicksFirstEnabled(t *testing.T) {
	eng, err := New("prefer-cheaper")
	require.NoError(t, err)
	pool := mkPool("cheap", "expensive")
	nr := normal.NormalRequest{SourceCLI: normal.SourceClaudeCode, Model: "x"}
	pick, _, err := eng.Pick(context.Background(), nr, pool, nil)
	require.NoError(t, err)
	require.Equal(t, "cheap", pick.ID)
}

func TestEngine_RoundRobin_Rotates(t *testing.T) {
	eng, err := New("round-robin")
	require.NoError(t, err)
	pool := mkPool("a", "b", "c")
	nr := normal.NormalRequest{SourceCLI: normal.SourceClaudeCode, Model: "x"}

	seen := make([]string, 0, 6)
	for i := 0; i < 6; i++ {
		pick, _, err := eng.Pick(context.Background(), nr, pool, nil)
		require.NoError(t, err)
		seen = append(seen, pick.ID)
	}
	require.Equal(t, []string{"a", "b", "c", "a", "b", "c"}, seen)
}

func TestEngine_SkipsUnhealthyUpstreams(t *testing.T) {
	eng, err := New("prefer-cheaper")
	require.NoError(t, err)
	pool := mkPool("a", "b")
	nr := normal.NormalRequest{SourceCLI: normal.SourceClaudeCode, Model: "x"}
	unhealthy := map[string]bool{"a": true}
	pick, _, err := eng.Pick(context.Background(), nr, pool, unhealthy)
	require.NoError(t, err)
	require.Equal(t, "b", pick.ID)
}

func TestEngine_ErrorsWhenAllUnhealthy(t *testing.T) {
	eng, err := New("prefer-cheaper")
	require.NoError(t, err)
	pool := mkPool("a")
	unhealthy := map[string]bool{"a": true}
	_, _, err = eng.Pick(context.Background(),
		normal.NormalRequest{SourceCLI: normal.SourceClaudeCode, Model: "x"},
		pool, unhealthy)
	require.Error(t, err)
}
