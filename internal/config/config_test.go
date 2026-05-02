package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad_DefaultsWhenFileMissing(t *testing.T) {
	dir := t.TempDir()
	c, err := Load(filepath.Join(dir, "ccg.toml"))
	require.NoError(t, err)
	require.Equal(t, "127.0.0.1:17180", c.Listen)
	require.Equal(t, "prefer-cheaper", c.Strategy)
	require.Empty(t, c.Upstreams)
}

func TestLoad_ParsesUpstreams(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "ccg.toml")
	require.NoError(t, os.WriteFile(path, []byte(`
listen = "127.0.0.1:17181"
strategy = "prefer-capable"

[[upstream]]
id = "anthropic-direct"
protocol = "anthropic"
base_url = "https://api.anthropic.com"
auth_header = "x-api-key: ${ANTHROPIC_API_KEY}"

[[upstream]]
id = "openai-direct"
protocol = "openai"
base_url = "https://api.openai.com"
auth_header = "Authorization: Bearer ${OPENAI_API_KEY}"
`), 0o600))

	c, err := Load(path)
	require.NoError(t, err)
	require.Equal(t, "127.0.0.1:17181", c.Listen)
	require.Equal(t, "prefer-capable", c.Strategy)
	require.Len(t, c.Upstreams, 2)
	require.Equal(t, "anthropic-direct", c.Upstreams[0].ID)
	require.Equal(t, "anthropic", c.Upstreams[0].Protocol)
}

func TestLoad_RejectsUnknownStrategy(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "ccg.toml")
	require.NoError(t, os.WriteFile(path, []byte(`strategy = "magical"`), 0o600))
	_, err := Load(path)
	require.Error(t, err)
}

func TestLoad_RequiresAuthTokenForNonLoopbackListen(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "ccg.toml")
	require.NoError(t, os.WriteFile(path, []byte(`listen = "0.0.0.0:17180"`), 0o600))
	_, err := Load(path)
	require.Error(t, err)
}

func TestLoad_AllowsAuthTokenForNonLoopbackListen(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "ccg.toml")
	require.NoError(t, os.WriteFile(path, []byte(`
listen = "0.0.0.0:17180"
auth_token = "secret"
`), 0o600))
	cfg, err := Load(path)
	require.NoError(t, err)
	require.Equal(t, "secret", cfg.AuthToken)
}
