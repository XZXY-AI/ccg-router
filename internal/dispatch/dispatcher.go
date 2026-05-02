// Package dispatch forwards a NormalRequest to a chosen Upstream.
//
// v0.1 covers non-streaming forwarding here. Streaming (SSE) is
// handled by a dedicated helper later so the happy-path non-stream code
// stays readable and free of flusher plumbing.
package dispatch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ccg-labs/ccg-router/internal/normal"
	"github.com/ccg-labs/ccg-router/internal/upstream"
)

// passthroughHeaders lists request headers that must be forwarded
// verbatim to the upstream.
var passthroughHeaders = []string{
	"anthropic-version",
	"anthropic-beta",
	"openai-beta",
	"openai-organization",
	"openai-project",
	"user-agent",
}

type Dispatcher struct {
	client *http.Client
}

func New() *Dispatcher {
	return &Dispatcher{client: &http.Client{Timeout: 60 * time.Second}}
}

// Do forwards nr.Extra["raw"] to upstream base URL + path, setting the
// resolved auth header and returning status, response headers, and body.
func (d *Dispatcher) Do(ctx context.Context, u upstream.Upstream,
	nr normal.NormalRequest, path string,
	inHdr http.Header) (int, http.Header, []byte, error) {
	raw := nr.Extra["raw"]
	if raw == nil {
		return 0, nil, nil, fmt.Errorf("missing raw body in NormalRequest")
	}
	raw = rewriteModel(raw, u.ModelMap)
	target := strings.TrimRight(u.BaseURL, "/") + path

	var lastErr error
	backoff := 250 * time.Millisecond
	const maxAttempts = 3
	for attempt := 0; attempt < maxAttempts; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewReader(raw))
		if err != nil {
			return 0, nil, nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		applyAuth(req, u.ResolvedAuthHeader)
		for _, h := range passthroughHeaders {
			if v := inHdr.Get(h); v != "" {
				req.Header.Set(h, v)
			}
		}

		resp, err := d.client.Do(req)
		if err != nil {
			lastErr = err
			if attempt < maxAttempts-1 {
				time.Sleep(backoff)
				backoff *= 2
			}
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if resp.StatusCode >= 500 && attempt < maxAttempts-1 {
			lastErr = fmt.Errorf("upstream %d", resp.StatusCode)
			time.Sleep(backoff)
			backoff *= 2
			continue
		}
		return resp.StatusCode, resp.Header.Clone(), body, nil
	}
	return 0, nil, nil, fmt.Errorf("dispatch exhausted retries: %w", lastErr)
}

func applyAuth(req *http.Request, header string) {
	if header == "" {
		return
	}
	if i := strings.Index(header, ":"); i > 0 {
		key := strings.TrimSpace(header[:i])
		val := strings.TrimSpace(header[i+1:])
		req.Header.Set(key, val)
	}
}

func rewriteModel(raw []byte, modelMap map[string]string) []byte {
	if len(modelMap) == 0 {
		return raw
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(raw, &obj); err != nil {
		return raw
	}
	modelRaw, ok := obj["model"]
	if !ok {
		return raw
	}
	var model string
	if err := json.Unmarshal(modelRaw, &model); err != nil {
		return raw
	}
	if mapped, ok := modelMap[model]; ok && mapped != "" {
		b, _ := json.Marshal(mapped)
		obj["model"] = b
		if out, err := json.Marshal(obj); err == nil {
			return out
		}
	}
	return raw
}
