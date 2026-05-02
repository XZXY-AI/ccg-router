# ccg-router v0.1.2

Patch release focused on version reporting across all install paths.

## Changes

- `ccg-router --version` now reports the module version for `go install ...@version` builds.
- GoReleaser-built binaries still include the release commit.

## Install

```bash
brew install XZXY-AI/tap/ccg-router
curl -fsSL https://raw.githubusercontent.com/XZXY-AI/ccg-router/main/scripts/install.sh | bash
go install github.com/XZXY-AI/ccg-router/cmd/ccg-router@latest
```

## Verify

```bash
ccg-router --version
ccg-router doctor
```
