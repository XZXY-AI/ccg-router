# ccg-router

**中文** · [English](README.md)

> 一个本地守护进程。Claude Code 和 Codex CLI 都接它。
> 同一个端口，同时讲 Anthropic 协议 **和** OpenAI 协议 —— 顺带一个记录每一次请求的本地 SQLite ledger。

![CI](https://github.com/XZXY-AI/ccg-router/actions/workflows/ci.yml/badge.svg)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)
![GitHub stars](https://img.shields.io/github/stars/XZXY-AI/ccg-router?style=social)

![ccg-router 用一个本地端口同时服务 Claude Code 和 Codex CLI](docs/demo.png)

## 30 秒就能跑起来

```bash
# 1. 安装
brew install XZXY-AI/tap/ccg-router

# 2. init + 把两个 CLI 都指向本地 router
ccg-router init
export ANTHROPIC_BASE_URL=http://127.0.0.1:17180
export OPENAI_BASE_URL=http://127.0.0.1:17180

# 3. 启动
ccg-router start
```

之后 Claude Code 和 Codex CLI 都经一个守护进程走。Provider key 留在 `~/.config/ccg-router/config.toml`。每一次请求写一行到本地 SQLite ledger，你可以直接用 `sqlite3` 查。

## 和 claude-code-router 有什么不同？

| | `claude-code-router` | **`ccg-router`** |
|---|---|---|
| Claude Code 路由 | ✓ | ✓ |
| Codex CLI 路由 | — | ✓ |
| 一个守护进程同时服务两个 CLI | — | ✓ |
| 本地 SQLite 用量 ledger | — | ✓ |
| 多种路由策略 | — | `prefer-cheaper` / `prefer-capable` / `round-robin` |
| 托管控制面 | — | — （也没有） |

如果你只用 Claude Code，[`claude-code-router`](https://github.com/musistudio/claude-code-router) 在单 CLI 场景下更成熟。`ccg-router` 是为双 CLI 工作流而存在的。

## 为什么今天就值得试

- 你在 Claude Code 和 Codex CLI 之间切换，手动改环境变量已经烦了。
- 你想要一个本地 per-request ledger，不用托管 dashboard 就能回答"这个副业项目实际花了多少钱"。
- 你想用一套配置切换路由策略，而不是去改 shell rc。

如果以上都不沾边，那暂时不需要 —— star 一下，等 `v0.2` 加上 streaming 再来。

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

## Guides

- [Claude Code router](docs/guides/claude-code-router.md)
- [Codex CLI router](docs/guides/codex-cli-router.md)
- [OpenAI-compatible proxy](docs/guides/openai-compatible-proxy.md)
- [Anthropic-compatible proxy](docs/guides/anthropic-compatible-proxy.md)
- [Local usage ledger](docs/guides/local-usage-ledger.md)
- [ccg-router vs claude-code-router](docs/guides/compare-claude-code-router.md)

## SEO Landing Pages

- [Claude Code router](https://xzxy-ai.github.io/ccg-router/claude-code-router/)
- [Codex CLI router](https://xzxy-ai.github.io/ccg-router/codex-cli-router/)
- [Route Claude Code to OpenAI-compatible APIs](https://xzxy-ai.github.io/ccg-router/claude-code-openai-compatible/)
- [Codex CLI OpenAI-compatible router](https://xzxy-ai.github.io/ccg-router/codex-cli-openai-compatible/)
- [ANTHROPIC_BASE_URL local router](https://xzxy-ai.github.io/ccg-router/anthropic-base-url/)
- [OPENAI_BASE_URL local router](https://xzxy-ai.github.io/ccg-router/openai-base-url/)
- [Claude Code usage tracking](https://xzxy-ai.github.io/ccg-router/claude-code-usage-tracking/)
- [Local LLM router for AI coding CLI tools](https://xzxy-ai.github.io/ccg-router/local-llm-router/)

## Tutorials

- [Use one local router for Claude Code and Codex CLI](https://xzxy-ai.github.io/ccg-router/tutorials/one-local-router/)
- [Run an OpenAI-compatible and Anthropic-compatible local proxy](https://xzxy-ai.github.io/ccg-router/tutorials/openai-anthropic-compatible-proxy/)
- [Track Claude Code and Codex CLI usage locally](https://xzxy-ai.github.io/ccg-router/tutorials/local-usage-ledger/)

## Troubleshooting

- [ANTHROPIC_BASE_URL not working with Claude Code](docs/errors/anthropic-base-url-not-working.md)
- [OPENAI_BASE_URL not working with Codex CLI](docs/errors/openai-base-url-not-working.md)
- [Homebrew install problem](docs/errors/homebrew-install.md)
- [go install ccg-router@latest problem](docs/errors/go-install-latest.md)
- [ccg-router doctor failed](docs/errors/doctor-failed.md)

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
