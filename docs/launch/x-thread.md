# X / Twitter Launch Thread

**Posting notes:**
- Thread of 6 tweets. First tweet is the hook. Last tweet asks for the star.
- Each tweet ≤ 270 chars to leave room for handles / quote-tweets.
- Attach `docs/demo.png` to tweet 1. Attach `docs/demo.gif` to tweet 2 if file ≤ 5MB.
- Post weekday morning US time.

---

## Tweet 1 (hook + image)

```
I got tired of editing ANTHROPIC_BASE_URL and OPENAI_BASE_URL by hand every time I switched between Claude Code and Codex CLI.

So I wrote a small Go daemon. One config, one port, two CLIs.

It's called ccg-router. 🧵
```

(Attach: `docs/demo.png`)

## Tweet 2 (the picture in motion)

```
What it does:

ccg-router runs on 127.0.0.1. Claude Code points ANTHROPIC_BASE_URL at it. Codex CLI points OPENAI_BASE_URL at it. Both go through the same daemon, the same upstream pool, the same routing strategies.
```

(Attach: `docs/demo.gif`)

## Tweet 3 (the differentiator that matters)

```
Per-request local SQLite ledger.

One row per call: timestamp, upstream, model, tokens.

I can finally answer "how much did this side project cost me?" without scraping a hosted dashboard.

No hosted control plane. Keys stay in ~/.config/ccg-router/config.toml.
```

## Tweet 4 (honest about scope)

```
v0.1 is non-streaming. Streaming passthrough is v0.2.

Three routing strategies: prefer-cheaper, prefer-capable, round-robin.

If you only use one CLI, you don't need this. claude-code-router is more mature for single-CLI.
```

## Tweet 5 (install)

```
Install:

brew install XZXY-AI/tap/ccg-router

or

go install github.com/XZXY-AI/ccg-router/cmd/ccg-router@latest

Apache-2.0. Single static binary. macOS + Linux.
```

## Tweet 6 (the ask)

```
Repo with full docs + the SQLite schema + the routing config reference:

https://github.com/XZXY-AI/ccg-router

If the dual-CLI friction is something you've also hit, a star helps other people in the same situation find it.
```

## After posting

- Pin the thread to your profile for the launch week.
- Quote-tweet any substantive replies from your own account.
- Do NOT mass-tag people. Cold tags get muted and the algorithm penalizes them.
