package router

import (
	"context"

	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/ccg-labs/ccg-router/internal/upstream"
)

type preferCapable struct{}

func (preferCapable) Pick(_ context.Context, _ normal.NormalRequest,
	pool []upstream.Upstream, unhealthy map[string]bool) (upstream.Upstream, Decision, error) {
	healthy := filterHealthy(pool, unhealthy)
	if len(healthy) == 0 {
		return upstream.Upstream{}, Decision{}, errNoUpstream
	}
	// v0.1 heuristic: last enabled is assumed "more capable";
	// user orders them in ccg.toml.
	names := make([]string, len(healthy))
	for i, u := range healthy {
		names[i] = u.ID
	}
	return healthy[len(healthy)-1],
		Decision{Reason: "last-enabled wins under prefer-capable", Candidates: names}, nil
}
