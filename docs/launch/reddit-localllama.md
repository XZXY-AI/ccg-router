# r/LocalLLaMA Post Draft

**Sub rules check:**
- r/LocalLLaMA is technical-leaning and skeptical of hosted tools. Lead with the local-only architecture, not the use-case.
- Mention Go, single static binary, no hosted control plane, in the first paragraph.
- Avoid words like "powerful", "seamless", "next-generation". Engineers downvote those on sight.

---

## Title

```
ccg-router: local Go daemon that proxies Claude Code + Codex CLI through one port, with a SQLite usage ledger
```

(116 chars.)

## Body

```
Posting because the r/LocalLLaMA philosophy (keep the metadata local, run the routing locally, don't hand your provider keys to a SaaS) is exactly what I built this around. It is not a local-LLM tool itself — it's a local routing layer for the API CLIs that most of us use alongside whatever we run locally.

What it does:
- Listens on 127.0.0.1, one port (default 17180).
- Speaks Anthropic-compatible /v1/messages AND OpenAI-compatible /v1/chat/completions on the same daemon.
- Forwards each request to an upstream you configured. Three strategies: prefer-cheaper, prefer-capable, round-robin.
- Writes one row per request to a local SQLite ledger. Schema is documented; you can query it directly with sqlite3.
- No hosted control plane. Keys stay in ~/.config/ccg-router/config.toml or env.

What it does not do:
- Streaming passthrough (v0.2).
- Encrypted ledger at rest (v0.4).
- Run a model. This is purely a routing/proxy/ledger layer.

Why I'm posting here specifically: the API CLIs are useful but the "two CLIs, two sets of env vars, no shared usage view" setup is annoying enough that I want to see if anyone else is solving it differently. If you've got a smaller hack that does the same thing, I want to read it. If you only use one of those CLIs you don't need ccg-router — claude-code-router is more mature for the single-CLI case.

Repo: https://github.com/XZXY-AI/ccg-router
Apache-2.0, Go single binary, brew/curl/go install.
```

## Expected hostile comments and FAQ mapping

Same map as `comments-faq.md`. r/LocalLLaMA specifically will press on:
- "Why not Caddy / nginx / litellm in front?" — answer: those don't have request-by-request ledger semantics, and they don't know about both shapes (Anthropic + OpenAI) in one config. Use the `litellm` comparison answer below.

### Extra answer: "Why not litellm?"

```
litellm is great if you want a Python proxy library inside your app. ccg-router is a standalone daemon, single static binary, and the unit of routing is "a request from Claude Code or Codex CLI" rather than "a request from your Python code". The audience overlap is partial. If you're already happy with litellm as your routing layer, ccg-router won't change your life.
```
