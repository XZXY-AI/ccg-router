package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "ccg-router",
		Short: "Unified local router for Claude Code and Codex CLI",
	}
	root.AddCommand(newInitCmd(), newStartCmd(), newDoctorCmd())
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
