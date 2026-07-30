package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/corani/adr/config"
	"gopkg.in/yaml.v3"
)

type Config = config.Config

func ProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("%w: get root: %w", ErrConfig, err)
	}

	path := cwd

	for {
		if exists(filepath.Join(path, ".adr.yaml")) {
			return path, nil
		}

		// TODO(daniel): more robust detection
		if exists(filepath.Join(path, ".git")) {
			return path, nil
		}

		path = filepath.Dir(path)
		if len(path) <= 1 {
			break
		}
	}

	return cwd, nil
}

func ReadConfig() (*config.Config, error) {
	root, err := ProjectRoot()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(root, ".adr.yaml")

	if exists(path) {
		out, err := os.Open(path) // #nosec G304
		if err != nil {
			return nil, fmt.Errorf("%w: read: %w", ErrConfig, err)
		}
		defer out.Close() //nolint:errcheck

		var cfg config.Config

		if err := yaml.NewDecoder(out).Decode(&cfg); err != nil {
			return nil, fmt.Errorf("%w: read: %w", ErrConfig, err)
		}

		cfg.Project = root

		return &cfg, nil
	}

	return nil, os.ErrNotExist
}

func WriteConfig(root string, cfg *config.Config) error {
	out, err := os.Create(filepath.Join(root, ".adr.yaml")) // #nosec G304
	if err != nil {
		return fmt.Errorf("%w: write: %w", ErrConfig, err)
	}
	defer out.Close() //nolint:errcheck

	if err := yaml.NewEncoder(out).Encode(cfg); err != nil {
		return fmt.Errorf("%w: write: %w", ErrConfig, err)
	}

	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}
