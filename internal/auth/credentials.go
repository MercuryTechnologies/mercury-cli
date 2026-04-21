package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/zalando/go-keyring"
)

const (
	keyringService = "mercury-cli"
	keyringTimeout = 3 * time.Second
)

var errKeyringTimeout = errors.New("keyring operation timed out")

// TokenSet holds OAuth tokens for a single environment.
type TokenSet struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type"`
	Expiry       time.Time `json:"expiry"`
}

// IsExpired returns true if the access token is expired or will expire within 30 seconds.
func (t *TokenSet) IsExpired() bool {
	return time.Now().After(t.Expiry.Add(-30 * time.Second))
}

// LoadToken returns the stored token set for the given environment. It
// checks the system keyring first, then the plaintext fallback file.
// Returns (nil, nil) if no tokens are stored in either location.
func LoadToken(environment string) (*TokenSet, error) {
	if secret, err := keyringGet(keyringService, environment); err == nil {
		var tokens TokenSet
		if jerr := json.Unmarshal([]byte(secret), &tokens); jerr != nil {
			return nil, fmt.Errorf("parsing stored credentials: %w", jerr)
		}
		return &tokens, nil
	}

	file, err := loadCredentialsFile()
	if err != nil {
		return nil, fmt.Errorf("reading fallback credentials: %w", err)
	}
	tokens := file[environment]
	if tokens == nil {
		return nil, nil
	}
	return tokens, nil
}

// SaveToken persists tokens for the given environment. It tries the system
// keyring first; if that fails or times out, it writes to a plaintext file
// at ~/.config/mercury/credentials.json with 0600 permissions. The returned
// bool is true when the plaintext fallback was used.
func SaveToken(environment string, tokens *TokenSet) (insecure bool, err error) {
	data, err := json.Marshal(tokens)
	if err != nil {
		return false, err
	}

	if kerr := keyringSet(keyringService, environment, string(data)); kerr == nil {
		// Keyring write succeeded — clean up any stale plaintext entry so
		// keyring is unambiguously the source of truth.
		_ = clearCredentialsFileEntry(environment)
		return false, nil
	}

	// Keyring write failed — best-effort purge of any prior keyring entry so
	// a successful keyringGet later can't shadow the fresh fallback value.
	_ = keyringDelete(keyringService, environment)

	file, err := loadCredentialsFile()
	if err != nil {
		return false, fmt.Errorf("reading fallback credentials: %w", err)
	}
	file[environment] = tokens
	if err := saveCredentialsFile(file); err != nil {
		return false, fmt.Errorf("writing fallback credentials: %w", err)
	}
	return true, nil
}

// ClearToken removes stored tokens for the given environment from both the
// keyring and the plaintext fallback. A missing keyring entry is not an
// error; other keyring failures propagate so a failed logout does not look
// successful while the tokens remain readable on the next LoadToken.
func ClearToken(environment string) error {
	var errs []error
	if err := keyringDelete(keyringService, environment); err != nil && !errors.Is(err, keyring.ErrNotFound) {
		errs = append(errs, fmt.Errorf("clearing keyring credentials: %w", err))
	}
	if err := clearCredentialsFileEntry(environment); err != nil {
		errs = append(errs, fmt.Errorf("clearing fallback credentials: %w", err))
	}
	return errors.Join(errs...)
}

// Timeouts protect against Secret Service / kwalletd hangs on Linux.

func keyringSet(service, user, secret string) error {
	ch := make(chan error, 1)
	go func() { ch <- keyring.Set(service, user, secret) }()
	select {
	case err := <-ch:
		return err
	case <-time.After(keyringTimeout):
		return errKeyringTimeout
	}
}

func keyringGet(service, user string) (string, error) {
	type result struct {
		val string
		err error
	}
	ch := make(chan result, 1)
	go func() {
		v, e := keyring.Get(service, user)
		ch <- result{v, e}
	}()
	select {
	case r := <-ch:
		return r.val, r.err
	case <-time.After(keyringTimeout):
		return "", errKeyringTimeout
	}
}

func keyringDelete(service, user string) error {
	ch := make(chan error, 1)
	go func() { ch <- keyring.Delete(service, user) }()
	select {
	case err := <-ch:
		return err
	case <-time.After(keyringTimeout):
		return errKeyringTimeout
	}
}

type credentialsFile map[string]*TokenSet

// credentialsPathFunc resolves the fallback file path. Tests swap it to
// redirect writes into a t.TempDir() so they don't touch ~/.config/mercury.
var credentialsPathFunc = defaultCredentialsPath

func CredentialsPath() (string, error) {
	return credentialsPathFunc()
}

func defaultCredentialsPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "mercury", "credentials.json"), nil
}

func loadCredentialsFile() (credentialsFile, error) {
	path, err := CredentialsPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return credentialsFile{}, nil
		}
		return nil, err
	}
	var file credentialsFile
	if err := json.Unmarshal(data, &file); err != nil {
		return nil, err
	}
	if file == nil {
		file = credentialsFile{}
	}
	return file, nil
}

func saveCredentialsFile(file credentialsFile) error {
	path, err := CredentialsPath()
	if err != nil {
		return err
	}
	if len(file) == 0 {
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(file, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

func clearCredentialsFileEntry(environment string) error {
	file, err := loadCredentialsFile()
	if err != nil {
		return err
	}
	if _, ok := file[environment]; !ok {
		return nil
	}
	delete(file, environment)
	return saveCredentialsFile(file)
}
