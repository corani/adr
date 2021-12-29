package internal

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func ProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	exists := func(path string) bool {
		_, err := os.Stat(path)
		return !errors.Is(err, os.ErrNotExist)
	}

	path := cwd
	for {
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

func AdrRoot() (string, error) {
	root, err := ProjectRoot()
	if err != nil {
		return "", err
	}

	adrRoot := ""

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// TODO(daniel): more robust detection
		if !d.IsDir() && d.Name() == "template.md" {
			adrRoot = filepath.Dir(path)

			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return adrRoot, nil
}
