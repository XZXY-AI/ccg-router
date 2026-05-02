# Getting Started

## Install

### Homebrew (macOS, Linux)

```bash
brew install XZXY-AI/tap/ccg-router
```

### Go

```bash
go install github.com/XZXY-AI/ccg-router/cmd/ccg-router@latest
```

## Three Steps

1. Initialize a config:

```bash
ccg-router init
```

2. Export the endpoints Claude Code and Codex CLI should use:

```bash
export ANTHROPIC_BASE_URL=http://127.0.0.1:17180
export OPENAI_BASE_URL=http://127.0.0.1:17180
```

3. Start the daemon:

```bash
ccg-router start
```

Open `http://127.0.0.1:17180/ui/` for the local panel.
