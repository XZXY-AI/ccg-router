package registry

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad_ValidSignatureAccepted(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)

	payload := SignedRegistry{
		Version: 1,
		Presets: []Preset{{
			ID: "trusted-provider-x", Protocol: "openai",
			BaseURL: "https://example.com", Recommended: true,
			AuthHeader: "Authorization: Bearer ${X_API_KEY}",
		}},
	}
	body, err := json.Marshal(payload)
	require.NoError(t, err)
	sig := ed25519.Sign(priv, body)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/registry.json" {
			_, _ = w.Write(body)
			return
		}
		if r.URL.Path == "/registry.sig" {
			_, _ = w.Write([]byte(base64.StdEncoding.EncodeToString(sig)))
			return
		}
		http.NotFound(w, r)
	}))
	defer srv.Close()

	pubB64 := base64.StdEncoding.EncodeToString(pub)
	reg, err := Load(context.Background(), Options{
		URL: srv.URL + "/registry.json", PublicKey: pubB64,
	})
	require.NoError(t, err)
	require.Len(t, reg.Presets, 1)
	require.Equal(t, "trusted-provider-x", reg.Presets[0].ID)
}

func TestLoad_TamperedBodyRejected(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	good := []byte(`{"version":1,"presets":[]}`)
	sig := ed25519.Sign(priv, good)
	tampered := []byte(`{"version":1,"presets":[{"id":"bad"}]}`)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/registry.json" {
			_, _ = w.Write(tampered)
			return
		}
		_, _ = w.Write([]byte(base64.StdEncoding.EncodeToString(sig)))
	}))
	defer srv.Close()
	_, err = Load(context.Background(), Options{
		URL:       srv.URL + "/registry.json",
		PublicKey: base64.StdEncoding.EncodeToString(pub),
	})
	require.Error(t, err)
}

func TestParseEd25519PublicKey_AcceptsRawAndSPKI(t *testing.T) {
	pub, _, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)

	rawB64 := base64.StdEncoding.EncodeToString(pub)
	gotRaw, err := parseEd25519PublicKey(rawB64)
	require.NoError(t, err)
	require.Equal(t, ed25519.PublicKey(pub), gotRaw)

	spki, err := x509.MarshalPKIXPublicKey(pub)
	require.NoError(t, err)
	spkiB64 := base64.StdEncoding.EncodeToString(spki)
	gotSPKI, err := parseEd25519PublicKey(spkiB64)
	require.NoError(t, err)
	require.Equal(t, ed25519.PublicKey(pub), gotSPKI)
}

func TestParseEd25519PublicKey_RejectsGarbage(t *testing.T) {
	_, err := parseEd25519PublicKey(base64.StdEncoding.EncodeToString([]byte("not-a-key")))
	require.Error(t, err)
}
