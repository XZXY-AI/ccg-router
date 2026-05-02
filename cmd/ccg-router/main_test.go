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
