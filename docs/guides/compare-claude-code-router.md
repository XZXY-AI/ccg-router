# ccg-router vs claude-code-router

Both projects can be useful for Claude Code routing.

`claude-code-router` focuses on Claude Code routing. `ccg-router` focuses on one local routing layer for Claude Code and Codex CLI together.

| Need | ccg-router |
|---|---|
| Claude Code router | Yes |
| Codex CLI router | Yes |
| OpenAI-compatible local endpoint | Yes |
| Anthropic-compatible local endpoint | Yes |
| Local SQLite usage ledger | Yes |

Choose `ccg-router` when you use Claude Code and Codex CLI side by side and want one local router with shared upstreams and local usage metadata.
