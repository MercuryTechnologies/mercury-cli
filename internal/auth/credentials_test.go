package auth

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

// setupIsolatedStorage installs an in-memory keyring mock and redirects the
// plaintext fallback path to a per-test temp dir. Returns the temp dir so
// tests can inspect the written file.
func setupIsolatedStorage(t *testing.T) string {
	t.Helper()
	keyring.MockInit()
	tempDir := t.TempDir()
	prev := credentialsPathFunc
	credentialsPathFunc = func() (string, error) {
		return filepath.Join(tempDir, "credentials.json"), nil
	}
	t.Cleanup(func() { credentialsPathFunc = prev })
	return tempDir
}

func sampleTokens() *TokenSet {
	return &TokenSet{
		AccessToken:  "access-abc",
		RefreshToken: "refresh-xyz",
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(time.Hour).UTC().Round(time.Second),
	}
}

func seedFallbackFile(t *testing.T, dir string, file credentialsFile) {
	t.Helper()
	data, err := json.MarshalIndent(file, "", "  ")
	require.NoError(t, err)
	require.NoError(t, os.MkdirAll(dir, 0700))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "credentials.json"), data, 0600))
}

func TestSaveAndLoadToken_UsesKeyring(t *testing.T) {
	tempDir := setupIsolatedStorage(t)
	tokens := sampleTokens()

	insecure, err := SaveToken("production", tokens)
	require.NoError(t, err)
	require.False(t, insecure, "expected keyring to be used")

	// No plaintext file should exist when the keyring accepted the write.
	_, statErr := os.Stat(filepath.Join(tempDir, "credentials.json"))
	require.True(t, os.IsNotExist(statErr), "expected no fallback file, got: %v", statErr)

	loaded, err := LoadToken("production")
	require.NoError(t, err)
	require.NotNil(t, loaded)
	require.Equal(t, tokens.AccessToken, loaded.AccessToken)
	require.Equal(t, tokens.RefreshToken, loaded.RefreshToken)
	require.Equal(t, tokens.TokenType, loaded.TokenType)
	require.WithinDuration(t, tokens.Expiry, loaded.Expiry, time.Second)
}

func TestSaveToken_KeyringErrorFallsBackToFile(t *testing.T) {
	tempDir := setupIsolatedStorage(t)
	keyring.MockInitWithError(errors.New("keyring unavailable"))

	tokens := sampleTokens()
	insecure, err := SaveToken("production", tokens)
	require.NoError(t, err)
	require.True(t, insecure, "expected plaintext fallback to be used")

	path := filepath.Join(tempDir, "credentials.json")
	info, err := os.Stat(path)
	require.NoError(t, err)
	if runtime.GOOS != "windows" {
		require.Equal(t, os.FileMode(0600), info.Mode().Perm(), "fallback file must be 0600")
	}

	loaded, err := LoadToken("production")
	require.NoError(t, err)
	require.NotNil(t, loaded)
	require.Equal(t, tokens.AccessToken, loaded.AccessToken)
}

func TestSaveToken_SuccessfulKeyringWriteCleansUpStaleFileEntry(t *testing.T) {
	tempDir := setupIsolatedStorage(t)

	// Simulate an upgrade scenario: a prior run wrote to the plaintext file
	// (keyring was unavailable then), and now keyring is working again.
	stale := sampleTokens()
	stale.AccessToken = "stale"
	seedFallbackFile(t, tempDir, credentialsFile{"production": stale})

	fresh := sampleTokens()
	fresh.AccessToken = "fresh"
	insecure, err := SaveToken("production", fresh)
	require.NoError(t, err)
	require.False(t, insecure)

	// The file entry for "production" should be gone — keyring is now the
	// only source of truth for that environment.
	file, err := loadCredentialsFile()
	require.NoError(t, err)
	require.Nil(t, file["production"], "expected stale plaintext entry to be removed")

	loaded, err := LoadToken("production")
	require.NoError(t, err)
	require.Equal(t, "fresh", loaded.AccessToken)
}

func TestSaveToken_CleanupPreservesOtherEnvironments(t *testing.T) {
	tempDir := setupIsolatedStorage(t)

	// Sandbox tokens live in the file (e.g., keyring was unavailable at the time
	// of that login). A successful keyring write for "production" must not
	// touch the sandbox entry.
	sandboxTokens := sampleTokens()
	sandboxTokens.AccessToken = "sandbox-token"
	seedFallbackFile(t, tempDir, credentialsFile{"sandbox": sandboxTokens})

	_, err := SaveToken("production", sampleTokens())
	require.NoError(t, err)

	file, err := loadCredentialsFile()
	require.NoError(t, err)
	require.NotNil(t, file["sandbox"], "sandbox entry should be preserved")
	require.Equal(t, "sandbox-token", file["sandbox"].AccessToken)
}

func TestLoadToken_ReadsFromFileWhenKeyringEmpty(t *testing.T) {
	tempDir := setupIsolatedStorage(t)

	legacy := sampleTokens()
	legacy.AccessToken = "legacy-token"
	seedFallbackFile(t, tempDir, credentialsFile{"production": legacy})

	loaded, err := LoadToken("production")
	require.NoError(t, err)
	require.NotNil(t, loaded)
	require.Equal(t, "legacy-token", loaded.AccessToken)
}

func TestLoadToken_ReadsFromFileWhenKeyringErrors(t *testing.T) {
	tempDir := setupIsolatedStorage(t)
	seedFallbackFile(t, tempDir, credentialsFile{"production": sampleTokens()})

	// Install an errorful mock *after* seeding; this mimics a box where the
	// keyring daemon is broken but the file is still readable.
	keyring.MockInitWithError(errors.New("keyring unavailable"))

	loaded, err := LoadToken("production")
	require.NoError(t, err)
	require.NotNil(t, loaded)
	require.Equal(t, "access-abc", loaded.AccessToken)
}

func TestLoadToken_MissingReturnsNil(t *testing.T) {
	setupIsolatedStorage(t)

	loaded, err := LoadToken("production")
	require.NoError(t, err)
	require.Nil(t, loaded)
}

func TestClearToken_RemovesFromBothBackends(t *testing.T) {
	tempDir := setupIsolatedStorage(t)

	_, err := SaveToken("production", sampleTokens())
	require.NoError(t, err)

	// Ensure there's something in the file too so we prove both paths are cleared.
	seedFallbackFile(t, tempDir, credentialsFile{
		"production": sampleTokens(),
		"sandbox":    sampleTokens(),
	})

	require.NoError(t, ClearToken("production"))

	loaded, err := LoadToken("production")
	require.NoError(t, err)
	require.Nil(t, loaded, "production should be gone from both keyring and file")

	// Sandbox entry in the file must remain.
	file, err := loadCredentialsFile()
	require.NoError(t, err)
	require.NotNil(t, file["sandbox"], "ClearToken(production) should not affect sandbox")
}

func TestClearToken_NoopWhenMissing(t *testing.T) {
	setupIsolatedStorage(t)
	require.NoError(t, ClearToken("production"))
}
