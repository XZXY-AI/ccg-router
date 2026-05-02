# Routing Strategies

## `prefer-cheaper`

Selection rule: pick the first enabled healthy upstream in config order.

Example: if your config lists `openai-direct` before `anthropic-direct`, the router picks `openai-direct` while it is healthy.

Use it when cost control matters more than model capability.

## `prefer-capable`

Selection rule: pick the last enabled healthy upstream in config order.

Example: list a low-cost endpoint first and a stronger endpoint last. The router picks the stronger endpoint while it is healthy.

Use it when answer quality matters more than spend.

## `round-robin`

Selection rule: rotate across enabled healthy upstreams in config order.

Example: with upstreams `a`, `b`, and `c`, six requests select `a`, `b`, `c`, `a`, `b`, `c`.

Use it when you want simple load distribution across equivalent upstreams.
