# Awesome-* List PR Blurbs

This file tracks (a) every awesome-* list we PR'd to, (b) the exact entry text we used, and (c) the PR URL + status.

Target lists (sorted by likely acceptance):

1. **hesreallyhim/awesome-claude-code** — most active claude-code list as of 2026-05.
2. **awesome-codex-cli** — search for the most-starred fork before submitting; the namespace shifts.
3. **steven2358/awesome-generative-ai** (Devtools section) — broader audience, slower review.
4. **sindresorhus/awesome** (parent) — out of scope for a 19-day-old repo; revisit at 100+ stars.

For each list we open one PR. Do not batch — list maintainers reject "shotgun" PRs.

---

## Entry text (use the variant that matches each list's existing format)

### One-line variant

```
- [ccg-router](https://github.com/XZXY-AI/ccg-router) — Local routing daemon that serves both Claude Code (Anthropic-compatible) and Codex CLI (OpenAI-compatible) on one port, with a local SQLite usage ledger and three routing strategies. Apache-2.0, Go.
```

### Multi-line variant (when list uses sub-bullets)

```
- [ccg-router](https://github.com/XZXY-AI/ccg-router)
  - Local Go daemon on 127.0.0.1.
  - Speaks Anthropic-compatible `/v1/messages` and OpenAI-compatible `/v1/chat/completions` on one port.
  - Three routing strategies: `prefer-cheaper`, `prefer-capable`, `round-robin`.
  - Local SQLite usage ledger.
  - Apache-2.0.
```

### Badged variant (when list uses shields.io)

```
- [ccg-router](https://github.com/XZXY-AI/ccg-router) ![GitHub stars](https://img.shields.io/github/stars/XZXY-AI/ccg-router?style=social) — Local routing daemon for Claude Code + Codex CLI, with a SQLite usage ledger. Apache-2.0, Go.
```

---

## PR submission tracker

| List | PR URL | Variant | Submitted | Merged | Notes |
|---|---|---|---|---|---|
| hesreallyhim/awesome-claude-code | _pending_ | one-line | _pending_ | — | First target. |
| (second list TBD) | _pending_ | — | _pending_ | — | Decide after #1 lands. |

When a PR is submitted, edit this table with the URL and date. When merged, edit again. This is the source of truth for "did we actually get listed."

---

## Submission process

1. Fork the target list to `ccg-labs`.
2. Edit the appropriate section. Match the existing alphabetical / category placement; do not re-order other entries.
3. Open PR. Title: `Add ccg-router (local Claude Code + Codex CLI router)`. Body: one paragraph why it fits the list, link to repo, link to license.
4. Do NOT @-mention the maintainer in the PR body.
5. Update the table above.
