package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	root := newRootCmd()
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "ccg-router",
		Short:   "Unified local router for Claude Code and Codex CLI",
		Version: version + " (" + commit + ")",
	}
	root.AddCommand(newInitCmd(), newStartCmd(), newDoctorCmd())
	return root
}
