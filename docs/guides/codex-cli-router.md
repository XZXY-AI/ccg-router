# Codex CLI Router

`ccg-router` can be used as a local Codex CLI router for OpenAI-compatible requests.

Set Codex CLI to use the local daemon:

```bash
export OPENAI_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

Use this when you want:

- A local OpenAI-compatible proxy for Codex CLI.
- One shared routing config across Codex CLI and Claude Code.
- Local request metadata in a SQLite usage ledger.

`v0.1` supports non-streaming requests. Streaming passthrough is planned for `v0.2`.
