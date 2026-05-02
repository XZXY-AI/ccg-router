# Usage Ledger

`ccg-router` records one local SQLite row per upstream request in `~/.ccg/ledger.db`.

Privacy stance: the ledger stores no API keys and no prompts. It stores routing and usage metadata only.

## Schema

Table: `entries`

| Column | Meaning |
|---|---|
| `ts` | UTC timestamp in RFC3339Nano format. |
| `source_cli` | `claude-code`, `codex`, or `unknown`. |
| `upstream_id` | Selected upstream id. |
| `model` | Requested model name. |
| `input_tokens` | Input token count when known. |
| `output_tokens` | Output token count when known. |
| `usd_cost` | Estimated USD cost when known. |
| `fallback` | `1` if the request used fallback, otherwise `0`. |

## 5h Spend Query

```sql
SELECT
  COUNT(*) AS requests,
  SUM(input_tokens) AS input_tokens,
  SUM(output_tokens) AS output_tokens,
  SUM(usd_cost) AS usd_cost
FROM entries
WHERE ts >= datetime('now', '-5 hours');
```

The local UI uses the same ledger as its data source.
