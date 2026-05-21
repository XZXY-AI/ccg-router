# Launch Checklist — ccg-router

Run this from top to bottom. Each day's actions should take under 30 minutes. Skipping a day is fine; do not skip the abort gate on D7.

**Channels you'll need logged in:**
- GitHub (ccg-labs identity, already used for the repo).
- Hacker News account, ≥ 30 days old.
- Reddit account, ≥ 30 days old, with at least some posting history outside r/ClaudeAI.
- X / Twitter account, ≥ 30 days old.

---

## D0 (today) — Repo polish merged

- [ ] README hero v2 merged into `main` (Task 9 of plan).
- [ ] README.zh-CN.md mirrored (Task 10).
- [ ] `docs/index.html` hero updated (Task 11).
- [ ] `docs/demo.png` regenerated as 1280×640 OG card (Task 12).
- [ ] `docs/launch/*.md` all committed.
- [ ] CI green.

**Manual step (only you can do this):**

Open https://github.com/XZXY-AI/ccg-router/settings, scroll to "Social preview", upload the new `docs/demo.png`. GitHub does not expose this via API.

---

## D1 — awesome-claude-code PR

- [ ] Find the current most-active awesome-claude-code list. As of plan-writing: `hesreallyhim/awesome-claude-code`. Confirm by checking star count + most recent commit.
- [ ] Fork it under `ccg-labs`.
- [ ] Add the entry from `docs/launch/awesome-pr-blurbs.md` (one-line variant unless the list uses sub-bullets).
- [ ] Open PR. Update the tracker table in `awesome-pr-blurbs.md`.

**Metric to watch:** PR opened (binary). Acceptance can take days.

---

## D2 — Second awesome-* PR

- [ ] Pick the second list (see candidates in `awesome-pr-blurbs.md`). If none seem active, skip.
- [ ] Same fork → edit → PR flow.

---

## D3 — Record `docs/demo.gif` v2

- [ ] On your laptop, install `vhs`: `brew install vhs`.
- [ ] Run: `vhs docs/launch/demo-gif.tape`.
- [ ] Check output: `du -h docs/demo.gif` — must be ≤ 5MB.
- [ ] If > 5MB, edit the tape file's `Set FontSize` down by 2 or reduce the recorded steps.
- [ ] Commit the regenerated `docs/demo.gif`.

(This step is yours because `vhs` records your terminal. The CI cannot record itself.)

---

## D4 — r/ClaudeAI post

- [ ] Read `docs/launch/reddit-claudeai.md`.
- [ ] Post title + body verbatim. Do not change the title — it's hand-tuned.
- [ ] 5–10 minutes after submission, post the self-comment from the same file.
- [ ] Reply to comments for the next 4 hours. Use answers from `docs/launch/comments-faq.md`.

**Metric to watch (12h):**
- ≥ 30 upvotes AND ≥ 1 substantive comment → queue D5.
- 5–30 upvotes → still queue D5 but expect a slower week.
- < 5 upvotes → do NOT post LocalLLaMA. Diagnose framing. Re-read top comments for what landed flat.

---

## D5 — X / Twitter thread

- [ ] Read `docs/launch/x-thread.md`.
- [ ] Post tweet 1 with `docs/demo.png` attached.
- [ ] Post tweets 2–6 as replies in the thread.
- [ ] Pin the thread to your profile.
- [ ] Quote-tweet any substantive replies from your own account.

---

## D6 — Decision: r/LocalLLaMA or hold

Based on D4 result:
- If green light, post r/LocalLLaMA from `docs/launch/reddit-localllama.md`.
- If yellow, skip and review framing instead.
- If red (under 5 upvotes on D4), STOP — go to D7 review.

---

## D7 — Mid-cycle review (ABORT GATE)

Count:
- Net new stars since D0: ____
- External issues opened: ____
- Awesome PRs merged: ____
- Reddit comments with substance: ____

**Go criteria for Phase 2 (HN):**
- Net new stars since D0 ≥ 5 AND
- At least one of: (a) awesome PR merged, (b) substantive Reddit comment with engagement, (c) > 1 external issue.

**If go:** schedule Show HN for D8 morning, 08:00 PT.

**If no-go:** DO NOT post HN. Open a follow-up spec for product / positioning diagnosis. The kit's job in this case is to have generated signal, not stars — write up what you learned in `docs/launch/postmortem-phase1.md`.

---

## D8 — Show HN (only if D7 said go)

- [ ] Verify title still ≤ 80 chars: `docs/launch/hn.md`.
- [ ] Submit at 08:00 PT exactly.
- [ ] Post the first comment within 60 seconds of submission.
- [ ] Stay near the keyboard for 4 hours. HN front-page rank is decided fast.
- [ ] Reply to top-level comments. Use `comments-faq.md` answers; never paste verbatim, always slightly edit to match the commenter's wording.

**Do NOT:**
- Ask friends to upvote. HN detects vote rings.
- Edit the title after submission.
- Argue with hostile comments. Answer the substance, move on.

---

## D9–D14 — Tail + overflow

- [ ] Triage every issue and external PR within 24h. Even "I won't merge this" gets a polite reply.
- [ ] If HN front-paged: schedule the r/codex / r/OpenAI post (D10).
- [ ] If HN died: do the postmortem (`docs/launch/postmortem-phase2.md`). Note what the front-page comments said.

---

## D14 — Final metric capture

Write to `docs/launch/results-2026-05-21.md`:
- Star delta (start → end).
- Channel attribution (which post brought which spike — use GitHub Insights → Traffic).
- Clones count.
- Top 3 issues / PRs received.
- One lesson worth remembering for the next launch.

This file is the ledger we use to plan the next cycle.
