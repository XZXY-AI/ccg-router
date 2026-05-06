# ccg-router doctor Failed

`ccg-router doctor` checks local config, ledger, registry, upstreams, and local routing setup.

```bash
ccg-router doctor
```

## Common Causes

- The config file does not exist. Run `ccg-router init`.
- The ledger path is not writable.
- The registry URL or Ed25519 public key is invalid.
- An upstream references an environment variable that is not set.
- The daemon is configured to bind outside loopback without an auth token.

## What To Share In An Issue

Include the full `ccg-router doctor` output, your OS, the ccg-router version, and the install method. Remove API keys, auth headers, provider tokens, and private upstream URLs before posting.

## Related Pages

- Public page: <https://xzxy-ai.github.io/ccg-router/errors/doctor-failed/>
- Configuration: [../configuration.md](../configuration.md)
- Preset registry: [../preset-registry.md](../preset-registry.md)
