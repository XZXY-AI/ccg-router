# Claude Code Router

`ccg-router` can be used as a local Claude Code router for Anthropic-compatible requests.

Set Claude Code to use the local daemon:

```bash
export ANTHROPIC_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

Use this when you want:

- A local Anthropic-compatible proxy for Claude Code.
- One shared routing config across Claude Code and Codex CLI.
- Local request metadata in a SQLite usage ledger.

`v0.1` supports non-streaming requests. Streaming passthrough is planned for `v0.2`.
