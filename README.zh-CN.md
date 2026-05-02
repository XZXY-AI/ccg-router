# ccg-router

[English](README.md) · **中文**

> Claude Code 和 Codex CLI 的统一本地路由器。
> 一个本地 OpenAI-compatible 和 Anthropic-compatible proxy，支持共享 upstream、路由策略和本地 SQLite 用量 ledger。

![CI](https://github.com/XZXY-AI/ccg-router/actions/workflows/ci.yml/badge.svg)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)
![GitHub stars](https://img.shields.io/github/stars/XZXY-AI/ccg-router?style=social)

![demo](docs/demo.gif)

`ccg-router` 运行在 `127.0.0.1`，接收 Anthropic-compatible 和 OpenAI-compatible 请求，把它们路由到你配置的 upstream，并把用量元数据写入本地 SQLite ledger。你的 provider key 保留在本地配置或环境变量里。它适合同时使用 Claude Code、Codex CLI、OpenAI-compatible API 和 Anthropic-compatible API 的开发者。

如果它能帮你少切几次 CLI 环境变量，给 repo 一个 star 可以帮助更多 AI coding 用户发现它。

## Why ccg-router?

| 工具 | 做什么 | 差异 |
|---|---|---|
| `claude-code-router` | 路由 Claude Code 流量 | 只聚焦单个 CLI |
| 手动切换 | 手动改 shell 环境变量 | 慢、不一致、没有 ledger |
| `ccg-router` | Claude Code 和 Codex CLI 的本地路由层 | 一套配置、共享路由、本地用量 ledger |

## Status

`v0.1` 是 public preview，支持非 streaming 请求。Streaming passthrough 计划放到 `v0.2`。

## What Works Today

- Claude Code router endpoint：Anthropic-compatible `/v1/messages`
- Codex CLI router endpoint：OpenAI-compatible `/v1/chat/completions`
- 单端口本地 OpenAI-compatible 和 Anthropic-compatible proxy
- 路由策略：`prefer-cheaper`、`prefer-capable`、`round-robin`
- 本地 SQLite usage ledger
- 签名 preset registry loader
- 只读本地 dashboard：`/ui/`

## Not Yet

- Streaming passthrough
- 托管 preset registry
- 加密本地 ledger
- 更深入的 dashboard analytics

## Quickstart

用 Homebrew、release installer 或 Go 安装：

```bash
brew install XZXY-AI/tap/ccg-router
curl -fsSL https://raw.githubusercontent.com/XZXY-AI/ccg-router/main/scripts/install.sh | bash
go install github.com/XZXY-AI/ccg-router/cmd/ccg-router@latest
```

```bash
ccg-router init
export ANTHROPIC_BASE_URL=http://127.0.0.1:17180
export OPENAI_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

打开 `http://127.0.0.1:17180/ui/`。

## How it works

Claude Code 把 Anthropic-compatible 请求发到 `127.0.0.1:17180`。Codex CLI 把 OpenAI-compatible 请求发到同一个 daemon。`ccg-router` 会归一化请求、选择 upstream、转发原始 body，并写入本地 ledger。

## Features

- Anthropic-compatible `/v1/messages`
- OpenAI-compatible `/v1/chat/completions`
- 三种路由策略
- 本地 SQLite usage ledger
- 签名 preset registry 加载器
- 只读本地 UI

## Compare

| 需求 | 更适合 |
|---|---|
| 只路由 Claude Code 流量 | `claude-code-router` |
| Claude Code 和 Codex CLI 共用一个本地路由层 | `ccg-router` |
| 手动切换 `ANTHROPIC_BASE_URL` 和 `OPENAI_BASE_URL` | Shell 环境变量 |
| 多个 compatible API 共用，同时让 provider key 留在本地 | `ccg-router` |

## Search Terms

用户通常会通过这些词找到本项目：Claude Code router、Codex CLI router、OpenAI-compatible proxy、Anthropic-compatible proxy、local LLM router、AI coding CLI router、本地 Claude Code 用量统计。

## Configuration

见 `docs/configuration.md`。

## Routing strategies

见 `docs/routing-strategies.md`。

## Preset registry

见 `docs/preset-registry.md`。

## Roadmap

- v0.1: 本地 daemon、路由、ledger、registry 校验、UI
- v0.2: streaming passthrough、更完整的本地 dashboard、更多 CLI adapter
- Later: 加密 ledger、plugin hooks、更深入的用量分析

## FAQ

见 `docs/faq.md`。

## Community

Star the repo: https://github.com/XZXY-AI/ccg-router
Discussions:  https://github.com/XZXY-AI/ccg-router/discussions
Hub:          https://github.com/XZXY-AI/awesome-ai-coding-cli

## Contributing

提交 PR 前请运行 `make test`。公开文档聚焦本地路由、官方直连 upstream 示例、隐私和可复现行为。

## License

Apache-2.0
