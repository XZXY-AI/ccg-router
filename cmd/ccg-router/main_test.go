package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRootVersionFlag(t *testing.T) {
	oldVersion := version
	version = "v-test"
	t.Cleanup(func() { version = oldVersion })

	var out bytes.Buffer
	cmd := newRootCmd()
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"--version"})

	require.NoError(t, cmd.Execute())
	require.Contains(t, out.String(), "v-test")
}

func TestFormatVersionFallsBackToModuleVersion(t *testing.T) {
	require.Equal(t, "0.1.2", formatVersion("dev", "none", "v0.1.2"))
	require.Equal(t, "v-test (abc123)", formatVersion("v-test", "abc123", "v0.1.2"))
}
