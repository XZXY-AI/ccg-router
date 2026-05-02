package router

import (
	"context"

	"github.com/XZXY-AI/ccg-router/internal/normal"
	"github.com/XZXY-AI/ccg-router/internal/upstream"
)

type roundRobin struct {
	engine *Engine // borrow the engine's per-instance counter
}

func (rr roundRobin) Pick(_ context.Context, _ normal.NormalRequest,
	pool []upstream.Upstream, unhealthy map[string]bool) (upstream.Upstream, Decision, error) {
	healthy := filterHealthy(pool, unhealthy)
	if len(healthy) == 0 {
		return upstream.Upstream{}, Decision{}, errNoUpstream
	}
	idx := int(rr.engine.rrCursor.Add(1)-1) % len(healthy)
	names := make([]string, len(healthy))
	for i, u := range healthy {
		names[i] = u.ID
	}
	return healthy[idx], Decision{Reason: "round-robin rotation", Candidates: names}, nil
}
