# ccg-router Troubleshooting

Use these pages when ccg-router installs correctly but Claude Code, Codex CLI, Homebrew, Go install, or `ccg-router doctor` does not behave as expected.

## Common Fixes

- [ANTHROPIC_BASE_URL not working with Claude Code](anthropic-base-url-not-working.md)
- [OPENAI_BASE_URL not working with Codex CLI](openai-base-url-not-working.md)
- [Homebrew install problem](homebrew-install.md)
- [go install ccg-router@latest problem](go-install-latest.md)
- [ccg-router doctor failed](doctor-failed.md)

## Quick Checks

```bash
ccg-router --version
ccg-router doctor
export ANTHROPIC_BASE_URL=http://127.0.0.1:17180
export OPENAI_BASE_URL=http://127.0.0.1:17180
ccg-router start
```

For the public troubleshooting pages, see <https://xzxy-ai.github.io/ccg-router/errors/>.
