# r/OpenAI (Codex CLI) Post Draft

**Sub rules check:**
- r/codex may not exist as an active sub — check before posting. If empty, fall back to r/OpenAI with a Codex-CLI angle, or skip.
- r/OpenAI moderation removes "tool launch" posts that don't clearly help users. Lead with the friction, not the announcement.

---

## Title

```
If you use Codex CLI alongside Claude Code: I made a local daemon so they share one set of upstreams + a usage ledger
```

(124 chars.)

## Body

```
Codex CLI is the one I keep on for autocompletes, Claude Code is the one I keep on for "do the thing" tasks. The setup that wasted my time was: every time I changed which model I was routing to, I had to fix both `OPENAI_BASE_URL` (for Codex CLI) and `ANTHROPIC_BASE_URL` (for Claude Code) in different terminals, and I never had a unified view of how much each one was actually spending.

ccg-router runs as one Go daemon on 127.0.0.1. Both CLIs talk to it. Same config, same routing strategies, same SQLite ledger you can query with sqlite3.

What's relevant for Codex CLI specifically:
- Speaks OpenAI-compatible /v1/chat/completions natively.
- The model name Codex CLI sends gets mapped to whichever upstream model you configured (so you can point "gpt-5" at a different actual provider if you want).
- Codex CLI's quirks around tool_calls and function-calling are preserved (the router does not rewrite tool-call payloads).

v0.1 = non-streaming. Streaming is the next major version. If you live on streaming completions in Codex CLI, this isn't ready for you.

Repo: https://github.com/XZXY-AI/ccg-router

Apache-2.0, Go single binary. Looking for feedback specifically from people running Codex CLI day-to-day.
```

## If r/codex sub does not exist or is dead

Skip this channel. Do not force-post to r/OpenAI if the post does not have a clear "this helps Codex CLI users" angle. Forced posts get removed and burn the account's posting reputation.
