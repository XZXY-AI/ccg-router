# ccg-router

[中文](README.zh-CN.md) · **English**

> One local router for Claude Code and Codex CLI.
> Share upstreams, switch strategies, and keep a local usage ledger without juggling shell env vars.

![CI](https://github.com/XZXY-AI/ccg-router/actions/workflows/ci.yml/badge.svg)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)
![GitHub stars](https://img.shields.io/github/stars/XZXY-AI/ccg-router?style=social)

![demo](docs/demo.gif)

`ccg-router` runs on `127.0.0.1`, accepts Anthropic-compatible and OpenAI-compatible requests, routes them to the upstreams you configure, and stores usage metadata in a local SQLite ledger. Your provider keys stay in your local config or environment.

If this saves you from switching CLI env vars by hand, starring the repo helps more AI coding users find it.

## Why ccg-router?

| Tool | What it does | Difference |
|---|---|---|
| `claude-code-router` | Routes Claude Code traffic | Single CLI focus |
| Manual switching | Change shell env vars by hand | Slow, inconsistent, no ledger |
| `ccg-router` | Local routing layer for Claude Code and Codex CLI | One config, shared routing, local usage ledger |

## Status

`v0.1` is a public preview for non-streaming requests. Streaming passthrough is planned for `v0.2`.

## Quickstart

Install from source with Go:

```bash
go install github.com/XZXY-AI/ccg-router/cmd/ccg-router@latest
```

```bash
ccg-router init
export ANTHROPIC_BASE_URL=http://127.0.0.1:17180
export OPENAI_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

Open `http://127.0.0.1:17180/ui/`.

Release binaries, the shell installer, and Homebrew formula are prepared for the first tagged release.

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

Star the repo: https://github.com/XZXY-AI/ccg-router
Discussions:  https://github.com/XZXY-AI/ccg-router/discussions
Hub:          https://github.com/XZXY-AI/awesome-ai-coding-cli

## Contributing

Run `make test` before opening a PR. Keep public docs focused on local routing, official direct upstream examples, privacy, and reproducible behavior.

## License

Apache-2.0
