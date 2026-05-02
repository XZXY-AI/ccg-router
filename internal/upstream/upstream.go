// Package upstream holds the resolved pool of Upstream endpoints.
package upstream

import (
	"fmt"
	"os"
	"strings"

	"github.com/XZXY-AI/ccg-router/internal/config"
)

type Upstream struct {
	ID                 string
	Protocol           string
	BaseURL            string
	ResolvedAuthHeader string
	ModelMap           map[string]string
	Enabled            bool
}

type Pool struct {
	byID map[string]Upstream
	ids  []string // insertion order
}

// NewPool resolves ${ENV} references in auth headers from the provided
// map (or from os.Getenv when a key is absent in the map).
func NewPool(cfg config.Config, env map[string]string) (*Pool, error) {
	p := &Pool{byID: make(map[string]Upstream)}
	lookup := func(k string) string {
		if v, ok := env[k]; ok {
			return v
		}
		return os.Getenv(k)
	}
	for _, u := range cfg.Upstreams {
		auth := expandEnv(u.AuthHeader, lookup)
		if _, exists := p.byID[u.ID]; exists {
			return nil, fmt.Errorf("duplicate upstream id %q", u.ID)
		}
		p.byID[u.ID] = Upstream{
			ID:                 u.ID,
			Protocol:           u.Protocol,
			BaseURL:            u.BaseURL,
			ResolvedAuthHeader: auth,
			ModelMap:           u.ModelMap,
			Enabled:            u.Enabled,
		}
		p.ids = append(p.ids, u.ID)
	}
	return p, nil
}

func (p *Pool) Get(id string) (Upstream, bool) {
	u, ok := p.byID[id]
	return u, ok
}

func (p *Pool) Enabled() []Upstream {
	out := make([]Upstream, 0, len(p.ids))
	for _, id := range p.ids {
		u := p.byID[id]
		if u.Enabled {
			out = append(out, u)
		}
	}
	return out
}

func expandEnv(s string, lookup func(string) string) string {
	for {
		i := strings.Index(s, "${")
		if i < 0 {
			return s
		}
		j := strings.Index(s[i:], "}")
		if j < 0 {
			return s
		}
		key := s[i+2 : i+j]
		s = s[:i] + lookup(key) + s[i+j+1:]
	}
}
