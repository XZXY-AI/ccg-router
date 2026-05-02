// Package config loads and validates ccg.toml.
package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type Upstream struct {
	ID         string            `toml:"id"`
	Protocol   string            `toml:"protocol"` // "anthropic" | "openai"
	BaseURL    string            `toml:"base_url"`
	AuthHeader string            `toml:"auth_header"` // template w/ ${ENV}
	ModelMap   map[string]string `toml:"model_map"`   // optional
	Enabled    bool              `toml:"enabled"`
}

type Registry struct {
	URL       string `toml:"url"`
	PublicKey string `toml:"public_key"` // base64 Ed25519
	Enabled   bool   `toml:"enabled"`
}

type Config struct {
	Listen    string     `toml:"listen"`
	Strategy  string     `toml:"strategy"`
	AuthToken string     `toml:"auth_token"`
	Upstreams []Upstream `toml:"upstream"`
	Registry  Registry   `toml:"registry"`
}

var validStrategies = map[string]bool{
	"prefer-cheaper": true,
	"prefer-capable": true,
	"round-robin":    true,
}

func Defaults() Config {
	return Config{
		Listen:   "127.0.0.1:17180",
		Strategy: "prefer-cheaper",
	}
}

func Load(path string) (Config, error) {
	c := Defaults()
	b, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return c, nil
	}
	if err != nil {
		return Config{}, err
	}
	if err := toml.Unmarshal(b, &c); err != nil {
		return Config{}, fmt.Errorf("parse %s: %w", path, err)
	}
	if !validStrategies[c.Strategy] {
		return Config{}, fmt.Errorf("unknown strategy %q", c.Strategy)
	}
	if !isLoopbackListen(c.Listen) && c.AuthToken == "" {
		return Config{}, fmt.Errorf("auth_token is required when listen is not loopback")
	}
	for i, u := range c.Upstreams {
		if u.ID == "" || u.Protocol == "" || u.BaseURL == "" {
			return Config{}, fmt.Errorf("upstream[%d] missing required fields", i)
		}
		if u.Protocol != "anthropic" && u.Protocol != "openai" {
			return Config{}, fmt.Errorf("upstream[%d] protocol must be anthropic or openai", i)
		}
	}
	return c, nil
}

func isLoopbackListen(listen string) bool {
	return strings.HasPrefix(listen, "127.") ||
		strings.HasPrefix(listen, "localhost:") ||
		strings.HasPrefix(listen, "[::1]:") ||
		strings.HasPrefix(listen, "::1:")
}
