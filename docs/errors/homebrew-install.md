# Homebrew Install Problem

Install ccg-router with the public Homebrew tap:

```bash
brew install XZXY-AI/tap/ccg-router
ccg-router --version
```

## Checklist

- Run `brew update` and try again.
- Confirm the tap is public with `brew tap XZXY-AI/tap`.
- If `ccg-router` installs but the command is not found, check that Homebrew's bin directory is in `PATH`.
- Use the shell installer as a fallback on Linux or macOS.

```bash
curl -fsSL https://raw.githubusercontent.com/XZXY-AI/ccg-router/main/scripts/install.sh | bash
```

## Related Pages

- Public page: <https://xzxy-ai.github.io/ccg-router/errors/homebrew-install/>
- Install page: <https://xzxy-ai.github.io/ccg-router/install/>
- Releases: <https://github.com/XZXY-AI/ccg-router/releases/latest>
