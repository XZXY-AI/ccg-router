# Configuration

`ccg-router init` writes `~/.ccg/ccg.toml`. Set `CCG_CONFIG` to use a different path.

## Fields

| Field | Description |
|---|---|
| `listen` | Local address the daemon binds to. Defaults to `127.0.0.1:17180`. |
| `strategy` | Routing strategy: `prefer-cheaper`, `prefer-capable`, or `round-robin`. |
| `auth_token` | Optional bearer token for local HTTP requests. Required when `listen` is not loopback. |
| `[[upstream]]` | One upstream endpoint definition. Order matters for routing. |
| `upstream.id` | Stable identifier used in logs, decisions, and ledger rows. |
| `upstream.protocol` | Wire protocol expected by the upstream: `anthropic` or `openai`. |
| `upstream.base_url` | Upstream API base URL. Do not include the request path. |
| `upstream.auth_header` | Header template. `${ENV_NAME}` values are read from the environment. |
| `upstream.enabled` | Whether the upstream can be selected by the router. |
| `upstream.model_map` | Optional model name rewrite map, applied by dispatch in a later task. |
| `[registry]` | Optional signed preset registry subscription. |
| `registry.enabled` | Whether to load the registry. |
| `registry.url` | HTTPS URL for `registry.json`. |
| `registry.public_key` | Base64 Ed25519 public key, raw 32-byte or DER-SPKI encoded. |

## Example

```toml
listen = "127.0.0.1:17180"
strategy = "prefer-cheaper"
auth_token = ""

[[upstream]]
id = "anthropic-direct"
protocol = "anthropic"
base_url = "https://api.anthropic.com"
auth_header = "x-api-key: ${ANTHROPIC_API_KEY}"
enabled = true

[[upstream]]
id = "openai-direct"
protocol = "openai"
base_url = "https://api.openai.com"
auth_header = "Authorization: Bearer ${OPENAI_API_KEY}"
enabled = true

[registry]
enabled = false
url = "https://example.com/ccg-router/registry.json"
public_key = "PASTE_BASE64_ED25519_PUBLIC_KEY"
```

`ccg-router` does **not** ship any third-party endpoint. You may subscribe to a community-maintained signed registry, or point to your own. See `docs/preset-registry.md`.
