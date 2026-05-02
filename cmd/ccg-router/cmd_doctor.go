package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/ccg-labs/ccg-router/internal/config"
	"github.com/ccg-labs/ccg-router/internal/ledger"
	"github.com/ccg-labs/ccg-router/internal/registry"
	"github.com/spf13/cobra"
)

func newDoctorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Print health of config, ledger, and registry",
		RunE: func(cmd *cobra.Command, args []string) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			cfgPath := defaultConfigPath(home)
			cfg, err := config.Load(cfgPath)
			report(cmd, "config", err)

			conn, err := net.DialTimeout("tcp", cfg.Listen, 300*time.Millisecond)
			if err == nil {
				fmt.Fprintf(cmd.OutOrStdout(), "[warn] something already listening on %s\n", cfg.Listen)
				_ = conn.Close()
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "[ok]   %s is free\n", cfg.Listen)
			}

			ledgerDir := filepath.Join(home, ".ccg")
			_ = os.MkdirAll(ledgerDir, 0o700)
			l, err := ledger.Open(filepath.Join(ledgerDir, "ledger.db"))
			report(cmd, "ledger", err)
			if l != nil {
				_ = l.Close()
			}

			if cfg.Registry.Enabled {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()
				_, err := registry.Load(ctx, registry.Options{
					URL: cfg.Registry.URL, PublicKey: cfg.Registry.PublicKey,
				})
				report(cmd, "registry", err)
			} else {
				fmt.Fprintln(cmd.OutOrStdout(), "[skip] registry disabled")
			}
			return nil
		},
	}
}

func report(cmd *cobra.Command, name string, err error) {
	if err != nil {
		fmt.Fprintf(cmd.OutOrStdout(), "[fail] %s: %v\n", name, err)
		return
	}
	fmt.Fprintf(cmd.OutOrStdout(), "[ok]   %s\n", name)
}
