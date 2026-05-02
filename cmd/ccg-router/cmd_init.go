package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const defaultConfig = `# ccg-router configuration
listen    = "127.0.0.1:17180"
strategy  = "prefer-cheaper"   # or "prefer-capable" | "round-robin"

# Official-direct upstreams. Replace env var names with yours, or leave
# empty and set the env var before ccg-router start.

[[upstream]]
id          = "anthropic-direct"
protocol    = "anthropic"
base_url    = "https://api.anthropic.com"
auth_header = "x-api-key: ${ANTHROPIC_API_KEY}"
enabled     = true

[[upstream]]
id          = "openai-direct"
protocol    = "openai"
base_url    = "https://api.openai.com"
auth_header = "Authorization: Bearer ${OPENAI_API_KEY}"
enabled     = true

# Optional: subscribe to an external preset registry (signed).
# See docs/preset-registry.md for how to self-host.
# [registry]
# enabled    = false
# url        = "https://example.com/ccg-router/registry.json"
# public_key = "PASTE_BASE64_ED25519_PUBKEY_HERE"
`

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Write a default config to ~/.ccg/ccg.toml",
		RunE: func(cmd *cobra.Command, args []string) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			dir := filepath.Join(home, ".ccg")
			if err := os.MkdirAll(dir, 0o700); err != nil {
				return err
			}
			path := filepath.Join(dir, "ccg.toml")
			if _, err := os.Stat(path); err == nil {
				return fmt.Errorf("%s already exists", path)
			}
			if err := os.WriteFile(path, []byte(defaultConfig), 0o600); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Wrote %s\n\n", path)
			fmt.Fprintln(cmd.OutOrStdout(), "To point Claude Code at ccg-router, add to your shell:")
			fmt.Fprintln(cmd.OutOrStdout(), "  export ANTHROPIC_BASE_URL=http://127.0.0.1:17180")
			fmt.Fprintln(cmd.OutOrStdout(), "To point Codex CLI at ccg-router, add:")
			fmt.Fprintln(cmd.OutOrStdout(), "  export OPENAI_BASE_URL=http://127.0.0.1:17180")
			return nil
		},
	}
}
