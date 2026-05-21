# Launch Comments FAQ

Pre-written answers to questions we expect on HN / Reddit / X. Tone: short, plain, no marketing voice. Engineers smell hype instantly.

If a comment is hostile, do not match the hostility. If a comment is asking in good faith, answer the substance, not the framing.

---

## "Why not just use claude-code-router?"

`claude-code-router` is great for routing Claude Code traffic. `ccg-router` solves a different shape:

- Both Claude Code and Codex CLI talk to the same local daemon. One config, one process, one port.
- A local SQLite ledger writes a row per request, so you can answer "how many tokens did I spend in Cursor vs. Claude Code last week?" without a hosted control plane.
- Three routing strategies (`prefer-cheaper` / `prefer-capable` / `round-robin`) over the upstreams you configure.

If you only use Claude Code, `claude-code-router` is fine — you don't need this.

## "Isn't this just `ANTHROPIC_BASE_URL=...` and `OPENAI_BASE_URL=...`?"

That gets you a proxy, yes. It does not get you:

- One process behind both env vars.
- A request-by-request ledger.
- Strategy-based upstream selection per request.
- Keeping your provider keys out of your CLI env every time you open a new shell.

If the env-var approach is enough for your workflow, great — that's a totally valid choice.

## "What's the threat model for the keys?"

Provider keys live in `~/.config/ccg-router/config.toml` or environment variables. The router never sends them to a hosted control plane (there is none). The local SQLite ledger does not store keys — only request metadata (timestamp, upstream, model, token counts).

If your laptop is compromised, the keys are compromised. Same threat model as keeping them in your shell rc.

## "Why Go?"

Single static binary, no runtime install. Works on macOS and Linux today. Windows is on the roadmap once we have a tester.

## "No streaming? Dealbreaker."

Yes, `v0.1` is non-streaming. Streaming passthrough is the headline `v0.2` feature; tracking issue is pinned. If streaming is non-negotiable for your workflow, this isn't ready for you yet — star it and check back in a few weeks.

## "Why should I trust a 19-day-old repo with my keys?"

You shouldn't. Read the code — it's 100% Go, no dependencies on hosted services, and the network call surface is small. The Apache-2.0 license means you can fork it the moment we go silent.

## "What about [hosted alternative]?"

Hosted routers add a network hop, require shipping your keys (or tokens scoped to your account) to a third party, and centralize observability. `ccg-router` is the opposite trade: everything local, no third party, you give up dashboards and convenience features in exchange.

## "Roadmap?"

- `v0.2`: streaming passthrough (the main thing).
- `v0.3`: ledger visualization in the local UI.
- `v0.4`: encrypted-at-rest ledger.

Issues with the `roadmap` label track each. PRs welcome.

## "Will you accept my PR?"

Probably yes for: bug fixes, new upstream adapters, docs, tests.
Probably no for: hosted-service integrations, telemetry, anything that requires us to run infrastructure.

## "Is this a fork of X?"

No. Independent implementation in Go. Inspirations are credited in `docs/credits.md` (TODO if this question actually comes up — do not preemptively write it).
