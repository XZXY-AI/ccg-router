# FAQ

## Does this work with Claude Max?

No. `ccg-router` is a local API router. It expects API-compatible endpoints and local CLI environment variables.

## Do I need an API key?

Yes, for direct official upstreams. Put keys in environment variables such as `ANTHROPIC_API_KEY` or `OPENAI_API_KEY`; do not paste secrets into public files.

## Is my data sent anywhere?

Requests are sent only to the upstream you configure and select. The local ledger stores no API keys and no prompts. Registry fetches do not include keys or ledger data.

## How is this different from `claude-code-router`?

`claude-code-router` focuses on Claude Code routing. `ccg-router` exposes both Anthropic-compatible and OpenAI-compatible local endpoints so Claude Code and Codex CLI can share one local routing layer.

## Where do I get a preset registry?

You can self-host one with `docs/preset-registry.md`, subscribe to a signed registry you trust, or leave registry support disabled and configure upstreams manually.
