package ledger

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLedger_RecordAndQuery(t *testing.T) {
	dir := t.TempDir()
	l, err := Open(filepath.Join(dir, "ledger.db"))
	require.NoError(t, err)
	t.Cleanup(func() { _ = l.Close() })

	now := time.Now().UTC()
	require.NoError(t, l.Record(context.Background(), Entry{
		Timestamp:    now,
		SourceCLI:    "claude-code",
		UpstreamID:   "anthropic-direct",
		Model:        "claude-sonnet-4-7",
		InputTokens:  123,
		OutputTokens: 45,
		USDCost:      0.001,
		Fallback:     false,
	}))

	entries, err := l.Window(context.Background(), now.Add(-time.Hour), now.Add(time.Hour))
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, 123, entries[0].InputTokens)
}

func TestLedger_SummaryAggregates(t *testing.T) {
	dir := t.TempDir()
	l, err := Open(filepath.Join(dir, "ledger.db"))
	require.NoError(t, err)
	defer l.Close()
	ctx := context.Background()
	now := time.Now().UTC()
	for i := 0; i < 3; i++ {
		require.NoError(t, l.Record(ctx, Entry{
			Timestamp: now, SourceCLI: "codex", UpstreamID: "o",
			Model: "gpt-5", InputTokens: 10, OutputTokens: 5, USDCost: 0.0005,
		}))
	}
	sum, err := l.Summary(ctx, now.Add(-time.Hour), now.Add(time.Hour))
	require.NoError(t, err)
	require.Equal(t, 3, sum.Requests)
	require.InDelta(t, 0.0015, sum.USDCost, 1e-6)
}
