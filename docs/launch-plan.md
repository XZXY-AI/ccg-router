# 首轮曝光计划

目标：让 `ccg-router` 在 48 小时内获得第一批 star、安装尝试和真实反馈。

## 发布前

- 确认 README 首屏能在 10 秒内讲清楚：
  - 给 Claude Code 和 Codex CLI 共用一个本地 router。
  - 支持 Anthropic-compatible 和 OpenAI-compatible endpoint。
  - 本地 SQLite ledger 记录用量。
  - Provider key 留在本地配置或环境变量。
- 确认 Quickstart 的第一条安装命令当前可用。
- 确认 latest CI 成功。
- 确认 Issues、Discussions、secret scanning、branch protection 已启用。

## 发布日动作

1. 发布 `v0.1.0`。
2. 验证 Go install、shell installer、Homebrew 三条安装路径。
3. Pin `XZXY-AI/ccg-router` 到 `XZXY-AI` 组织首页。
4. 在 `XZXY-AI/awesome-ai-coding-cli` 增加 `ccg-router` 入口。
5. 发第一轮传播。

## 推荐渠道

- GitHub Release notes
- X/Twitter
- Hacker News `Show HN`
- Reddit：`r/ClaudeAI`、`r/OpenAI`、`r/LocalLLaMA`
- Claude Code / Codex CLI 用户社群
- 中文开发者社群和 AI coding 群

## 短文案

English:

```text
I built ccg-router: one local router for Claude Code and Codex CLI.

It accepts Anthropic-compatible and OpenAI-compatible requests, routes them through one config, and keeps a local SQLite usage ledger.

No hosted control plane. Your provider keys stay local.

Repo: https://github.com/XZXY-AI/ccg-router
```

中文：

```text
我做了 ccg-router：Claude Code 和 Codex CLI 共用的本地 router。

它同时接 Anthropic-compatible 和 OpenAI-compatible 请求，用一套配置做路由，并在本地 SQLite ledger 里记录用量。

没有托管控制面，provider key 留在本地。

Repo: https://github.com/XZXY-AI/ccg-router
```

## 48 小时内只看这些指标

- GitHub stars
- clone 数
- release download 数
- issue/discussion 数
- README 访问后的安装问题

如果 star 有增长但 issue 里集中反馈安装问题，优先修 Quickstart 和 release 产物，不急着扩功能。
