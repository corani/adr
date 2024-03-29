package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Project       string `yaml:"-"`
	Root          string `yaml:"root"`
	AdrTemplate   string `yaml:"adrTemplate"`
	IndexTemplate string `yaml:"indexTemplate"`
}

func ProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("%w: get root: %v", ErrConfig, err)
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

func ReadConfig() (*Config, error) {
	root, err := ProjectRoot()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(root, ".adr.yaml")

	if exists(path) {
		out, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("%w: read: %v", ErrConfig, err)
		}
		defer out.Close()

		var config Config

		if err := yaml.NewDecoder(out).Decode(&config); err != nil {
			return nil, fmt.Errorf("%w: read: %v", ErrConfig, err)
		}

		config.Project = root

		return &config, nil
	}

	return nil, os.ErrNotExist
}

func WriteConfig(root string, config *Config) error {
	out, err := os.Create(filepath.Join(root, ".adr.yaml"))
	if err != nil {
		return fmt.Errorf("%w: write: %v", ErrConfig, err)
	}
	defer out.Close()

	if err := yaml.NewEncoder(out).Encode(config); err != nil {
		return fmt.Errorf("%w: write: %v", ErrConfig, err)
	}

	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}
