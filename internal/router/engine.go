// Package router picks an upstream for a NormalRequest given a strategy.
package router

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/ccg-labs/ccg-router/internal/upstream"
)

type Decision struct {
	StrategyUsed string
	Reason       string
	Candidates   []string
}

type Strategy interface {
	Pick(ctx context.Context, nr normal.NormalRequest,
		pool []upstream.Upstream,
		unhealthy map[string]bool) (upstream.Upstream, Decision, error)
}

type Engine struct {
	name string
	s    Strategy
	// rrCursor is per-Engine so separate engines (tests, multi-tenant)
	// do not interfere with each other. Only the roundRobin strategy
	// actually reads it.
	rrCursor atomic.Uint64
}

func New(name string) (*Engine, error) {
	e := &Engine{name: name}
	switch name {
	case "prefer-cheaper":
		e.s = &preferCheaper{}
	case "prefer-capable":
		e.s = &preferCapable{}
	case "round-robin":
		e.s = &roundRobin{engine: e}
	default:
		return nil, fmt.Errorf("unknown strategy %q", name)
	}
	return e, nil
}

func (e *Engine) Pick(ctx context.Context, nr normal.NormalRequest,
	pool []upstream.Upstream,
	unhealthy map[string]bool) (upstream.Upstream, Decision, error) {
	u, d, err := e.s.Pick(ctx, nr, pool, unhealthy)
	d.StrategyUsed = e.name
	return u, d, err
}

var errNoUpstream = errors.New("no healthy upstream available")

func filterHealthy(pool []upstream.Upstream,
	unhealthy map[string]bool) []upstream.Upstream {
	out := make([]upstream.Upstream, 0, len(pool))
	for _, u := range pool {
		if unhealthy[u.ID] {
			continue
		}
		out = append(out, u)
	}
	return out
}
