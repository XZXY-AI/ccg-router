package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/XZXY-AI/ccg-router/internal/config"
	"github.com/XZXY-AI/ccg-router/internal/ledger"
	"github.com/XZXY-AI/ccg-router/internal/router"
	"github.com/XZXY-AI/ccg-router/internal/server"
	"github.com/XZXY-AI/ccg-router/internal/upstream"
	"github.com/spf13/cobra"
)

func newStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the ccg-router daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			cfgPath := defaultConfigPath(home)
			cfg, err := config.Load(cfgPath)
			if err != nil {
				return fmt.Errorf("load config: %w", err)
			}
			pool, err := upstream.NewPool(cfg, nil)
			if err != nil {
				return fmt.Errorf("upstream pool: %w", err)
			}
			eng, err := router.New(cfg.Strategy)
			if err != nil {
				return fmt.Errorf("router: %w", err)
			}
			ledgerDir := filepath.Join(home, ".ccg")
			if err := os.MkdirAll(ledgerDir, 0o700); err != nil {
				return err
			}
			l, err := ledger.Open(filepath.Join(ledgerDir, "ledger.db"))
			if err != nil {
				return fmt.Errorf("ledger: %w", err)
			}
			defer l.Close()

			srv := &http.Server{
				Addr:              cfg.Listen,
				Handler:           server.New(server.Deps{Pool: pool, Engine: eng, Ledger: l, AuthToken: cfg.AuthToken}).Handler(),
				ReadHeaderTimeout: 10 * time.Second,
			}
			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer stop()
			go func() {
				<-ctx.Done()
				sd, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()
				_ = srv.Shutdown(sd)
			}()
			fmt.Fprintf(cmd.OutOrStdout(), "ccg-router listening on %s\n", cfg.Listen)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}
			return nil
		},
	}
}

func defaultConfigPath(home string) string {
	if v := os.Getenv("CCG_CONFIG"); v != "" {
		return v
	}
	return filepath.Join(home, ".ccg", "ccg.toml")
}
