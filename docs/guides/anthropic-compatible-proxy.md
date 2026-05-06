# Anthropic-Compatible Proxy

`ccg-router` exposes an Anthropic-compatible local endpoint for `/v1/messages`.

This is useful for Claude Code workflows that support `ANTHROPIC_BASE_URL`.

```bash
export ANTHROPIC_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

The same daemon can also expose an OpenAI-compatible endpoint for Codex CLI. Provider keys stay in local config or environment variables.
