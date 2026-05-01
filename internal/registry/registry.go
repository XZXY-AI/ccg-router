// Package registry loads a signed, externally-hosted preset registry.
//
// Contract: the registry provider publishes two files at the same prefix,
// `registry.json` and `registry.sig` (base64 Ed25519 signature over the
// raw registry.json bytes). Any mismatch rejects the load.
package registry

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Preset struct {
	ID          string            `json:"id"`
	Protocol    string            `json:"protocol"` // "anthropic" | "openai"
	BaseURL     string            `json:"base_url"`
	AuthHeader  string            `json:"auth_header"`
	Recommended bool              `json:"recommended"`
	Notes       string            `json:"notes,omitempty"`
	ModelMap    map[string]string `json:"model_map,omitempty"`
}

type SignedRegistry struct {
	Version int      `json:"version"`
	Presets []Preset `json:"presets"`
}

type Options struct {
	URL       string
	PublicKey string // base64 Ed25519 public key
	Timeout   time.Duration
}

func Load(ctx context.Context, opts Options) (SignedRegistry, error) {
	if opts.URL == "" {
		return SignedRegistry{}, errors.New("registry url required")
	}
	if opts.PublicKey == "" {
		return SignedRegistry{}, errors.New("registry public_key required")
	}
	timeout := opts.Timeout
	if timeout == 0 {
		timeout = 5 * time.Second
	}
	client := &http.Client{Timeout: timeout}

	body, err := httpGet(ctx, client, opts.URL)
	if err != nil {
		return SignedRegistry{}, fmt.Errorf("fetch registry: %w", err)
	}
	sigURL := strings.TrimSuffix(opts.URL, ".json") + ".sig"
	sigBody, err := httpGet(ctx, client, sigURL)
	if err != nil {
		return SignedRegistry{}, fmt.Errorf("fetch signature: %w", err)
	}

	// opts.PublicKey is base64 of the DER-encoded SPKI public key (what
	// `openssl pkey -pubout -outform DER | openssl base64 -A` produces).
	// We accept a raw 32-byte key as a fallback so tests and power users
	// can skip the SPKI wrapper.
	pub, err := parseEd25519PublicKey(opts.PublicKey)
	if err != nil {
		return SignedRegistry{}, err
	}
	sig, err := base64.StdEncoding.DecodeString(strings.TrimSpace(string(sigBody)))
	if err != nil || len(sig) != ed25519.SignatureSize {
		return SignedRegistry{}, errors.New("invalid signature encoding")
	}
	if !ed25519.Verify(pub, body, sig) {
		return SignedRegistry{}, errors.New("signature verification failed")
	}

	var reg SignedRegistry
	if err := json.Unmarshal(body, &reg); err != nil {
		return SignedRegistry{}, fmt.Errorf("parse registry: %w", err)
	}
	for i, p := range reg.Presets {
		if p.ID == "" || p.BaseURL == "" || p.Protocol == "" {
			return SignedRegistry{}, fmt.Errorf("preset[%d] missing required fields", i)
		}
	}
	return reg, nil
}

func httpGet(ctx context.Context, c *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("http %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

// parseEd25519PublicKey accepts either:
//   - base64 of the 32-byte raw public key (tests, power users), or
//   - base64 of the DER-encoded SPKI wrapper (what openssl emits with
//     `-pubout -outform DER`).
func parseEd25519PublicKey(encoded string) (ed25519.PublicKey, error) {
	raw, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, errors.New("public_key is not valid base64")
	}
	if len(raw) == ed25519.PublicKeySize {
		return ed25519.PublicKey(raw), nil
	}
	pubAny, err := x509.ParsePKIXPublicKey(raw)
	if err != nil {
		return nil, fmt.Errorf("public_key: not raw 32-byte and not DER SPKI: %w", err)
	}
	pub, ok := pubAny.(ed25519.PublicKey)
	if !ok {
		return nil, errors.New("public_key SPKI is not Ed25519")
	}
	return pub, nil
}
