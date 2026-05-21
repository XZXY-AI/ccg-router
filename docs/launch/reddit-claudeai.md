# r/ClaudeAI Post Draft

**Sub rules check:**
- r/ClaudeAI allows self-promo if it's actually useful to Claude Code users. Don't lead with "I made a thing."
- Lead with the problem you had. Tool comes second.
- Flair: "Tool / Resource" if available, otherwise no flair.
- Post weekday morning US time.

---

## Title

```
I kept editing ANTHROPIC_BASE_URL by hand every time I switched between Claude Code and Codex CLI – made a small local router that handles it
```

(143 chars. Reddit title cap is 300.)

## Body

```
For the past month I've been bouncing between Claude Code and Codex CLI on the same machine, and the friction was the env vars. Every time I wanted Claude Code to talk to a different upstream, I had to remember which shell had which `ANTHROPIC_BASE_URL` set, which terminal had Codex CLI's `OPENAI_BASE_URL`, and whether I'd accidentally clobbered them with a `source ~/.zshrc`.

So I wrote ccg-router. It's a Go daemon you run on 127.0.0.1 that:

- speaks both Anthropic-compatible (`/v1/messages`) and OpenAI-compatible (`/v1/chat/completions`) on one port,
- routes each request to upstreams you configure (so Claude Code and Codex CLI can share the same upstream pool or have different ones),
- writes a row to a local SQLite ledger per request, so I can finally answer "how much did this side project actually cost me?"

Three routing strategies: prefer-cheaper, prefer-capable, round-robin. Read-only local UI at /ui/.

v0.1 is non-streaming. Streaming is the next major version.

Repo: https://github.com/XZXY-AI/ccg-router

Mostly looking for feedback from anyone who actually has this dual-CLI setup. If you only use one CLI you don't need it.
```

## Comment to leave on your own post 5–10 minutes after submission

```
One thing I'll add: this is not a fork of claude-code-router. If you only use Claude Code, claude-code-router is a more mature single-purpose tool. The reason I built ccg-router separately is that I wanted both CLIs talking to one daemon with shared upstreams and one ledger, and forking would have been weird given how different the request-shape handling is.
```

## When to escalate to LocalLLaMA

If this post hits 30+ upvotes in 12h and gets at least 1 substantive comment, queue the LocalLLaMA post for the next morning. If it stalls under 5 upvotes, do NOT push LocalLLaMA — diagnose framing first.
