# Social Posts

Use these posts after a release is verified.

## X / Twitter

```text
I built ccg-router: one local router for Claude Code and Codex CLI.

It accepts Anthropic-compatible and OpenAI-compatible requests, routes them through one config, and keeps a local SQLite usage ledger.

No hosted control plane. Your provider keys stay local.

https://github.com/XZXY-AI/ccg-router
```

## Hacker News

Title:

```text
Show HN: ccg-router – one local router for Claude Code and Codex CLI
```

Body:

```text
I built ccg-router after getting tired of switching local CLI endpoint env vars by hand.

It runs on 127.0.0.1 and exposes both Anthropic-compatible and OpenAI-compatible endpoints, so Claude Code and Codex CLI can share one local routing layer. It supports basic routing strategies, records request metadata in a local SQLite ledger, and keeps provider keys in local config or environment variables.

v0.1 is a non-streaming public preview. Streaming passthrough is the next major feature.

Repo: https://github.com/XZXY-AI/ccg-router
```

## Reddit

```text
I built a small local router for people using Claude Code and Codex CLI side by side.

The problem: both tools can point at compatible API endpoints, but manually switching env vars and tracking usage gets messy.

ccg-router runs locally on 127.0.0.1, exposes Anthropic-compatible and OpenAI-compatible endpoints on one daemon, routes requests through one config, and keeps a local SQLite usage ledger. No hosted control plane and no prompt/key storage in the ledger.

Repo: https://github.com/XZXY-AI/ccg-router
```

## 中文社区

```text
我做了 ccg-router：Claude Code 和 Codex CLI 共用的本地 router。

它运行在 127.0.0.1，同时接 Anthropic-compatible 和 OpenAI-compatible 请求，用一套配置做路由，并在本地 SQLite ledger 里记录用量元数据。

没有托管控制面，provider key 留在本地配置或环境变量里。

Repo: https://github.com/XZXY-AI/ccg-router
```
