# ANTHROPIC_BASE_URL Not Working with Claude Code

If Claude Code is not reaching ccg-router, first verify that the local daemon is running and that `ANTHROPIC_BASE_URL` points to the local router.

```bash
export ANTHROPIC_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

## Checklist

- Run `ccg-router doctor` and fix config or ledger errors.
- Confirm your upstream has `protocol = "anthropic"`.
- Confirm the upstream auth header references an environment variable that exists.
- Use non-streaming requests with v0.1. Streaming passthrough is planned for v0.2.
- Make sure no shell profile overrides `ANTHROPIC_BASE_URL` after you set it.
- Launch Claude Code from the same shell session where `ANTHROPIC_BASE_URL` is exported.

## Related Pages

- Public page: <https://xzxy-ai.github.io/ccg-router/errors/anthropic-base-url-not-working/>
- Local router setup: <https://xzxy-ai.github.io/ccg-router/anthropic-base-url/>
- Claude Code router guide: <https://xzxy-ai.github.io/ccg-router/claude-code-router/>
