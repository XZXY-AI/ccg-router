# Show HN Draft

**Submission rules check before posting:**
- Title must start with `Show HN:` and be ≤ 80 chars.
- Body must be plain text. No links to social media, no marketing copy.
- Post Tue–Thu 08:00–10:00 PT for highest dwell.
- Be ready to answer questions for the first 4 hours; HN front-page dies quickly without OP engagement.

---

## Title (80 chars max)

```
Show HN: ccg-router – one local router for Claude Code and Codex CLI
```

(69 chars including "Show HN: ".)

## URL

```
https://github.com/XZXY-AI/ccg-router
```

## Body

```
I got tired of editing ANTHROPIC_BASE_URL and OPENAI_BASE_URL by hand every time I switched between Claude Code and Codex CLI. So I wrote a small local daemon that does the routing instead.

ccg-router runs on 127.0.0.1, speaks both Anthropic-compatible /v1/messages and OpenAI-compatible /v1/chat/completions, and routes each request to whichever upstream I configured. Provider keys live in a local config file. A SQLite ledger records one row per request so I can see how much each project actually cost me at the end of the week.

Three routing strategies today: prefer-cheaper, prefer-capable, round-robin. There's also a read-only local UI at /ui/.

v0.1 is the non-streaming public preview. Streaming passthrough is the next major version. The repo is 19 days old and I'm still figuring out which features matter, so feedback from anyone actually using Claude Code + Codex CLI side-by-side would be useful.

Apache-2.0, Go single binary, brew / curl / go install.
```

Word count: ~165 words. HN front-page submissions usually run 100–250 words.

## First comment (post immediately after submission)

```
A few things I left out of the post for length:

- It's not a fork or wrapper of claude-code-router. Independent implementation. The audience overlap is partial but real: ccg-router is aimed at people who use both Claude Code and Codex CLI on the same machine. If you only use one, you do not need this.

- The "ledger" is just SQLite, not a service. Schema is documented in docs/usage-ledger.md. You can query it directly.

- Threat model: no hosted control plane, no telemetry. Keys never leave the box. Local SQLite ledger does not store keys, only metadata.

- I tried to keep the v0.1 surface small. The known gaps (streaming, ledger viz, Windows) are tracked as roadmap-labeled issues.

Happy to answer specifics.
```

## Replies to canned objections

For each objection below, copy the matching answer from `docs/launch/comments-faq.md`:

| Objection on HN | FAQ answer to use |
|---|---|
| "Why not claude-code-router?" | "Why not just use claude-code-router?" |
| "This is just two env vars" | "Isn't this just ANTHROPIC_BASE_URL=…?" |
| "Why should I trust your keys handling?" | "What's the threat model for the keys?" |
| "No streaming = useless" | "No streaming? Dealbreaker." |

Do NOT auto-reply. Wait for the comment to appear, then answer with the FAQ content, lightly edited to fit the specific phrasing.
