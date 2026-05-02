# ccg-router

[中文](README.zh-CN.md) · **English**

> A unified local router for Claude Code and Codex CLI.
> One config, smart fallback, real usage insights.

![CI](https://github.com/ccg-labs/ccg-router/actions/workflows/ci.yml/badge.svg)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)

![demo](docs/demo.gif)

## Why ccg-router?

| Tool | What it does | Difference |
|---|---|---|
| `claude-code-router` | Routes Claude Code traffic | Single CLI focus |
| Manual switching | Change shell env vars by hand | Slow, inconsistent, no ledger |
| `ccg-router` | Local routing layer for Claude Code and Codex CLI | One config, shared routing, local usage ledger |

## Quickstart

```bash
ccg-router init
export ANTHROPIC_BASE_URL=http://127.0.0.1:17180
export OPENAI_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

Open `http://127.0.0.1:17180/ui/`.

## How it works

Claude Code sends Anthropic-compatible requests to `127.0.0.1:17180`. Codex CLI sends OpenAI-compatible requests to the same daemon. `ccg-router` normalizes the request, selects an upstream, forwards the raw body, and records a local ledger row.

## Features

- Anthropic-compatible `/v1/messages`
- OpenAI-compatible `/v1/chat/completions`
- Three routing strategies
- Local SQLite usage ledger
- Signed preset registry loader
- Read-only local UI

## Configuration

See `docs/configuration.md`.

## Routing strategies

See `docs/routing-strategies.md`.

## Preset registry

See `docs/preset-registry.md`.

## Roadmap

- v0.1: local daemon, routing, ledger, registry verification, UI
- v0.2: streaming passthrough, model map dispatch, more CLI adapters
- Later: encrypted ledger, plugin hooks, deeper usage analytics

## FAQ

See `docs/faq.md`.

## Community

See the hub:   https://github.com/ccg-labs/awesome-ai-coding-cli
Join Discord:  https://discord.gg/ccg-labs

## Contributing

Run `make test` before opening a PR. Keep public docs focused on local routing, official direct upstream examples, privacy, and reproducible behavior.

## License

Apache-2.0
