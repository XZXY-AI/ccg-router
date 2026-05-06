# OPENAI_BASE_URL Not Working with Codex CLI

If Codex CLI is not reaching ccg-router, verify that the local daemon is running and that `OPENAI_BASE_URL` points to the local router.

```bash
export OPENAI_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

## Checklist

- Run `ccg-router doctor`.
- Confirm your upstream has `protocol = "openai"`.
- Confirm the auth header uses `Authorization: Bearer ${OPENAI_API_KEY}` or your provider's required shape.
- Use non-streaming requests with v0.1.
- Check that your shell exports `OPENAI_BASE_URL` in the same session that launches Codex CLI.
- If another shell profile sets `OPENAI_BASE_URL`, remove the conflicting value or export the local router URL last.

## Related Pages

- Public page: <https://xzxy-ai.github.io/ccg-router/errors/openai-base-url-not-working/>
- Local router setup: <https://xzxy-ai.github.io/ccg-router/openai-base-url/>
- Codex CLI router guide: <https://xzxy-ai.github.io/ccg-router/codex-cli-router/>
