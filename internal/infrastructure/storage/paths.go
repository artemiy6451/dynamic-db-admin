// Package storage resolves per-user application storage locations.
package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const AppDirectoryName = "Dynamic DB Admin"

type Paths struct {
	ConfigDir string
	DataDir   string
	CacheDir  string
}

// ResolvePaths returns platform-appropriate per-user directories. It does not
// create them; callers should call Ensure before writing application data.
func ResolvePaths() (Paths, error) {
	configRoot, err := os.UserConfigDir()
	if err != nil {
		return Paths{}, fmt.Errorf("resolve user config directory: %w", err)
	}
	cacheRoot, err := os.UserCacheDir()
	if err != nil {
		return Paths{}, fmt.Errorf("resolve user cache directory: %w", err)
	}

	dataRoot, err := userDataRoot()
	if err != nil {
		return Paths{}, err
	}

	return Paths{
		ConfigDir: filepath.Join(configRoot, AppDirectoryName),
		DataDir:   filepath.Join(dataRoot, AppDirectoryName),
		CacheDir:  filepath.Join(cacheRoot, AppDirectoryName),
	}, nil
}

// Ensure creates the application directories with user-only permissions.
func (p Paths) Ensure() error {
	for _, dir := range []string{p.ConfigDir, p.DataDir, p.CacheDir} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return fmt.Errorf("create application directory %q: %w", dir, err)
		}
	}
	return nil
}

func userDataRoot() (string, error) {
	switch runtime.GOOS {
	case "windows":
		if root := os.Getenv("LOCALAPPDATA"); root != "" {
			return root, nil
		}
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve user home directory: %w", err)
		}
		return filepath.Join(home, "AppData", "Local"), nil
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve user home directory: %w", err)
		}
		return filepath.Join(home, "Library", "Application Support"), nil
	case "linux":
		if root := os.Getenv("XDG_DATA_HOME"); root != "" {
			return root, nil
		}
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve user home directory: %w", err)
		}
		return filepath.Join(home, ".local", "share"), nil
	default:
		return os.UserConfigDir()
	}
}
