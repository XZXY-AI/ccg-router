// Package ledger persists per-request usage into a local SQLite DB.
// Never stores API keys or user prompts.
package ledger

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Entry struct {
	Timestamp    time.Time
	SourceCLI    string
	UpstreamID   string
	Model        string
	InputTokens  int
	OutputTokens int
	USDCost      float64
	Fallback     bool
}

type Summary struct {
	Requests     int
	InputTokens  int
	OutputTokens int
	USDCost      float64
}

type Ledger struct {
	db *sql.DB
}

func Open(path string) (*Ledger, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS entries(
			ts TEXT NOT NULL,
			source_cli TEXT NOT NULL,
			upstream_id TEXT NOT NULL,
			model TEXT NOT NULL,
			input_tokens INTEGER NOT NULL,
			output_tokens INTEGER NOT NULL,
			usd_cost REAL NOT NULL,
			fallback INTEGER NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_entries_ts ON entries(ts);
	`); err != nil {
		return nil, fmt.Errorf("create schema: %w", err)
	}
	return &Ledger{db: db}, nil
}

func (l *Ledger) Close() error { return l.db.Close() }

func (l *Ledger) Record(ctx context.Context, e Entry) error {
	_, err := l.db.ExecContext(ctx,
		`INSERT INTO entries(ts, source_cli, upstream_id, model,
			input_tokens, output_tokens, usd_cost, fallback)
		 VALUES(?, ?, ?, ?, ?, ?, ?, ?)`,
		e.Timestamp.UTC().Format(time.RFC3339Nano),
		e.SourceCLI, e.UpstreamID, e.Model,
		e.InputTokens, e.OutputTokens, e.USDCost, boolToInt(e.Fallback))
	return err
}

func (l *Ledger) Window(ctx context.Context, from, to time.Time) ([]Entry, error) {
	rows, err := l.db.QueryContext(ctx,
		`SELECT ts, source_cli, upstream_id, model,
		        input_tokens, output_tokens, usd_cost, fallback
		 FROM entries
		 WHERE ts >= ? AND ts <= ?
		 ORDER BY ts ASC`,
		from.UTC().Format(time.RFC3339Nano),
		to.UTC().Format(time.RFC3339Nano),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Entry
	for rows.Next() {
		var e Entry
		var ts string
		var fb int
		if err := rows.Scan(&ts, &e.SourceCLI, &e.UpstreamID, &e.Model,
			&e.InputTokens, &e.OutputTokens, &e.USDCost, &fb); err != nil {
			return nil, err
		}
		t, err := time.Parse(time.RFC3339Nano, ts)
		if err != nil {
			return nil, err
		}
		e.Timestamp = t
		e.Fallback = fb == 1
		out = append(out, e)
	}
	return out, rows.Err()
}

func (l *Ledger) Summary(ctx context.Context, from, to time.Time) (Summary, error) {
	row := l.db.QueryRowContext(ctx,
		`SELECT COUNT(*), COALESCE(SUM(input_tokens),0),
		        COALESCE(SUM(output_tokens),0), COALESCE(SUM(usd_cost),0)
		 FROM entries WHERE ts >= ? AND ts <= ?`,
		from.UTC().Format(time.RFC3339Nano),
		to.UTC().Format(time.RFC3339Nano),
	)
	var s Summary
	if err := row.Scan(&s.Requests, &s.InputTokens, &s.OutputTokens, &s.USDCost); err != nil {
		return Summary{}, err
	}
	return s, nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
