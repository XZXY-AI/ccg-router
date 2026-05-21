# Awesome-* List PR Blurbs

This file tracks (a) every awesome-* list we PR'd to, (b) the exact entry text we used, and (c) the PR URL + status.

Target lists (sorted by likely acceptance):

1. **RoggeOhta/awesome-codex-cli** — 196★, exact topical fit ("Model Providers & Proxies"), explicit "PRs welcome" badge. ✅ PR opened.
2. **hesreallyhim/awesome-claude-code** — 44k★. ⚠️ Does NOT accept PRs. New resources must be submitted via the github.com issue form by a human. Programmatic submission triggers auto-close + spam-deterrent penalties on the submitter's account. Handoff details below.
3. **steven2358/awesome-generative-ai** (Devtools section) — broader audience, slower review. Revisit if we want more reach after the codex-cli PR lands.
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

| List | URL | Variant | Submitted | Merged | Notes |
|---|---|---|---|---|---|
| RoggeOhta/awesome-codex-cli | [PR #49](https://github.com/RoggeOhta/awesome-codex-cli/pull/49) | Model Providers & Proxies, badged | 2026-05-21 | — | Submitted from `linny006`. |
| hesreallyhim/awesome-claude-code | _user UI submission required_ | issue form | _pending_ | — | See "Manual issue-form submission" section below. |

When a PR or issue is submitted, edit this table with the URL and date. When merged, edit again. This is the source of truth for "did we actually get listed."

---

## Manual issue-form submission for `hesreallyhim/awesome-claude-code`

The maintainer explicitly forbids `gh` CLI / API submissions: "submissions will be automatically closed" with "increasingly severe penalties" on the submitter's account. The submission must be a human-driven UI form fill.

**You (the operator) do this:**

1. Open https://github.com/hesreallyhim/awesome-claude-code/issues/new?template=recommend-resource.yml in a browser while logged in.
2. Paste the field values below.
3. Tick all five checkboxes in the Recommendation Checklist.
4. Submit.

| Field | Value |
|---|---|
| Display Name | `ccg-router` |
| Category | `Tooling` |
| Sub-Category | (leave blank, none of the existing options fit cleanly; the maintainer will classify) |
| Primary Link | `https://github.com/XZXY-AI/ccg-router` |
| Author Name | `XZXY-AI` |
| Author Link | `https://github.com/XZXY-AI` |
| License | `Apache-2.0` |
| Description | `Local Go daemon serving Claude Code (Anthropic-compatible /v1/messages) and Codex CLI (OpenAI-compatible /v1/chat/completions) on one port. SQLite usage ledger per request, three routing strategies (prefer-cheaper / prefer-capable / round-robin), no hosted control plane.` |
| Validate Claims | `Run brew install XZXY-AI/tap/ccg-router && ccg-router init && ccg-router start; export ANTHROPIC_BASE_URL=http://127.0.0.1:17180 and start Claude Code — requests are visible at sqlite3 ~/.local/share/ccg-router/ledger.db.` |
| Specific Task(s) | `Configure one upstream with prefer-cheaper strategy, run a Claude Code session, then query the local ledger to see which upstream was selected per request.` |
| Specific Prompt(s) | `In the Claude Code session: "list files in the current directory". Then: sqlite3 ~/.local/share/ccg-router/ledger.db 'select upstream, model, total_tokens from requests order by id desc limit 5;'` |
| Additional Comments | (leave blank) |

After submitting, update the tracker row above with the issue URL.

---

## Submission process for normal PR-friendly lists

1. Fork the target list to your account (or `ccg-labs` if you have admin rights there).
2. Edit the appropriate section. Match the existing alphabetical / category placement; do not re-order other entries.
3. Open PR. Title: `Add ccg-router (local Claude Code + Codex CLI router)`. Body: one paragraph why it fits the list, link to repo, link to license.
4. Do NOT @-mention the maintainer in the PR body.
5. Update the table above.
