package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

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
		Version: buildVersion(),
	}
	root.AddCommand(newInitCmd(), newStartCmd(), newDoctorCmd())
	return root
}

func buildVersion() string {
	moduleVersion := ""
	if info, ok := debug.ReadBuildInfo(); ok {
		moduleVersion = info.Main.Version
	}
	return formatVersion(version, commit, moduleVersion)
}

func formatVersion(ldVersion, ldCommit, moduleVersion string) string {
	if ldVersion == "" || ldVersion == "dev" {
		if moduleVersion != "" && moduleVersion != "(devel)" {
			return strings.TrimPrefix(moduleVersion, "v")
		}
	}
	if ldCommit == "" || ldCommit == "none" {
		return ldVersion
	}
	return ldVersion + " (" + ldCommit + ")"
}
