# ccg-router v0.1.0

`ccg-router` 是 Claude Code 和 Codex CLI 的统一本地路由器：一套配置、共享 upstream、路由策略、本地用量 ledger。

## Highlights

- 支持 Anthropic-compatible `/v1/messages`
- 支持 OpenAI-compatible `/v1/chat/completions`
- 支持 `prefer-cheaper`、`prefer-capable`、`round-robin` 三种路由策略
- 本地 SQLite usage ledger
- 只读本地 UI：`/ui/`
- 签名 preset registry loader，支持 raw 32-byte 和 DER-SPKI Ed25519 public key
- Linux/macOS amd64/arm64 release archives
- cosign keyless signatures

## Install

```bash
go install github.com/XZXY-AI/ccg-router/cmd/ccg-router@latest
```

Release binary、shell installer 和 Homebrew formula 会在正式发布后可用。

## Known Limits

- `v0.1` 支持非 streaming 请求。
- Streaming passthrough 计划放到 `v0.2`。
- 默认不内置第三方 endpoint；用户需要配置自己的 upstream。

## Verify

```bash
ccg-router doctor
```
