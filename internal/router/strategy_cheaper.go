package router

import (
	"context"

	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/ccg-labs/ccg-router/internal/upstream"
)

type preferCheaper struct{}

func (preferCheaper) Pick(_ context.Context, _ normal.NormalRequest,
	pool []upstream.Upstream, unhealthy map[string]bool) (upstream.Upstream, Decision, error) {
	healthy := filterHealthy(pool, unhealthy)
	if len(healthy) == 0 {
		return upstream.Upstream{}, Decision{}, errNoUpstream
	}
	names := make([]string, len(healthy))
	for i, u := range healthy {
		names[i] = u.ID
	}
	return healthy[0], Decision{Reason: "first-enabled wins under prefer-cheaper", Candidates: names}, nil
}
