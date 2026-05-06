# OpenAI-Compatible Proxy

`ccg-router` exposes an OpenAI-compatible local endpoint for `/v1/chat/completions`.

This is useful for AI coding CLI tools that support `OPENAI_BASE_URL` or OpenAI-compatible endpoint configuration.

```bash
export OPENAI_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

The same daemon can also expose an Anthropic-compatible endpoint for Claude Code. Provider keys stay in local config or environment variables.
