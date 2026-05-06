# Local Usage Ledger

`ccg-router` records request metadata in a local SQLite usage ledger.

The ledger is designed for local AI coding CLI workflows:

- Claude Code usage tracking
- Codex CLI usage tracking
- Upstream routing visibility
- Local-only request metadata

The ledger does not store provider API keys and does not store prompts.

Use `ccg-router doctor` to check local config, ledger, and registry health.
