# Preset Registry

## What a Registry Is

A registry is a signed JSON file that describes upstream presets. The provider publishes:

- `registry.json`: the preset list
- `registry.sig`: base64 Ed25519 signature over the raw `registry.json` bytes

`ccg-router` rejects the registry if the signature does not verify.

## Subscribe

```toml
[registry]
enabled = true
url = "https://presets.ccg-labs.dev/registry.json"
public_key = "PASTE_BASE64_ED25519_PUBLIC_KEY"
```

Then run:

```bash
ccg-router doctor
```

Expect `[ok] registry` when the URL, key, and signature chain are valid.

## Self-Host

Generate an Ed25519 key pair:

```bash
openssl genpkey -algorithm Ed25519 -out registry.key
openssl pkey -in registry.key -pubout -out registry.pub
openssl pkey -in registry.key -pubout -outform DER | openssl base64 -A
```

Create `registry.json`:

```json
{
  "version": 1,
  "presets": [
    {
      "id": "example-openai",
      "protocol": "openai",
      "base_url": "https://example.invalid",
      "auth_header": "Authorization: Bearer ${EXAMPLE_API_KEY}",
      "recommended": false,
      "model_map": {}
    }
  ]
}
```

Sign it:

```bash
#!/usr/bin/env bash
set -euo pipefail
: "${REGISTRY_KEY:?path to Ed25519 private key required}"
openssl pkeyutl -sign -rawin -inkey "$REGISTRY_KEY" -in registry.json \
  | openssl base64 -A > registry.sig
```

Serve both files over HTTPS.
