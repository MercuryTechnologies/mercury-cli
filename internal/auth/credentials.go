package auth

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

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

// Credentials maps environment names to their token sets.
type Credentials map[string]*TokenSet

// credentialsDir returns the directory for storing Mercury CLI config.
func credentialsDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "mercury"), nil
}

// CredentialsPath returns the path to the credentials file.
func CredentialsPath() (string, error) {
	dir, err := credentialsDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "credentials.json"), nil
}

// LoadCredentials reads credentials from disk. Returns empty credentials if the file doesn't exist.
func LoadCredentials() (Credentials, error) {
	path, err := CredentialsPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Credentials{}, nil
		}
		return nil, err
	}

	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, err
	}
	return creds, nil
}

// SaveCredentials writes credentials to disk with secure permissions.
func SaveCredentials(creds Credentials) error {
	dir, err := credentialsDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return err
	}

	path := filepath.Join(dir, "credentials.json")
	return os.WriteFile(path, data, 0600)
}

// ClearCredentials removes the token set for a specific environment.
func ClearCredentials(environment string) error {
	creds, err := LoadCredentials()
	if err != nil {
		return err
	}
	delete(creds, environment)
	if len(creds) == 0 {
		// Remove the file entirely if no credentials remain.
		path, err := CredentialsPath()
		if err != nil {
			return err
		}
		err = os.Remove(path)
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return SaveCredentials(creds)
}
