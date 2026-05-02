# FAQ

## Does this work with Claude Max?

No. `ccg-router` is a local API router. It expects API-compatible endpoints and local CLI environment variables.

## Do I need an API key?

Yes, for direct official upstreams. Put keys in environment variables such as `ANTHROPIC_API_KEY` or `OPENAI_API_KEY`; do not paste secrets into public files.

## Is my data sent anywhere?

Requests are sent only to the upstream you configure and select. The local ledger stores no API keys and no prompts. Registry fetches do not include keys or ledger data.

## How is this different from `claude-code-router`?

`claude-code-router` focuses on Claude Code routing. `ccg-router` exposes both Anthropic-compatible and OpenAI-compatible local endpoints so Claude Code and Codex CLI can share one local routing layer.

## Can Claude Code and Codex CLI share the same local router?

Yes. Claude Code can point Anthropic-compatible traffic at `http://127.0.0.1:17180`, while Codex CLI can point OpenAI-compatible traffic at the same local daemon.

## Can I use this as an OpenAI-compatible proxy?

Yes, for non-streaming `/v1/chat/completions` requests in `v0.1`. Configure an OpenAI-compatible upstream and set `OPENAI_BASE_URL=http://127.0.0.1:17180`.

## Can I use this as an Anthropic-compatible proxy?

Yes, for non-streaming `/v1/messages` requests in `v0.1`. Configure an Anthropic-compatible upstream and set `ANTHROPIC_BASE_URL=http://127.0.0.1:17180`.

## Does `ccg-router` track Claude Code usage locally?

It records request metadata in a local SQLite ledger. It does not store API keys and does not store prompts.

## Does `ccg-router` support streaming?

Not yet. `v0.1` is a non-streaming public preview. Streaming passthrough is planned for `v0.2`.

## Where do I get a preset registry?

You can self-host one with `docs/preset-registry.md`, subscribe to a signed registry you trust, or leave registry support disabled and configure upstreams manually.
