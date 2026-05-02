package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitCommandWritesDefaultConfig(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	cmd := newInitCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)

	require.NoError(t, cmd.Execute())
	path := filepath.Join(home, ".ccg", "ccg.toml")
	b, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Contains(t, string(b), `listen    = "127.0.0.1:17180"`)
	require.Contains(t, out.String(), "Wrote ")
}
