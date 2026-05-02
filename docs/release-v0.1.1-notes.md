# ccg-router v0.1.1

Patch release focused on install and launch readiness.

## Changes

- Add `ccg-router --version` so users can verify installed release binaries.
- Update README Quickstart to show Homebrew, shell installer, and Go install now that release artifacts exist.
- Keep `v0.1` status clear: non-streaming public preview; streaming passthrough is planned for `v0.2`.

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
