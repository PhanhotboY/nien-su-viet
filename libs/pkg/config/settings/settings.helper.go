package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"emperror.dev/errors"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
)

func searchForConfigFileDir(
	rootDir string,
	env environment.Environment,
) (string, error) {
	var result string

	// Walk the directory tree starting from rootDir
	err := filepath.Walk(
		rootDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Check if the file is named "config.%s.json" (replace %s with the env)
			if !info.IsDir() &&
				strings.EqualFold(
					info.Name(),
					fmt.Sprintf("config.%s.json", env),
				) ||
				strings.EqualFold(
					info.Name(),
					fmt.Sprintf("config.%s.yaml", env),
				) ||
				strings.EqualFold(
					info.Name(),
					fmt.Sprintf("config.%s.yml", env),
				) {
				// Get the directory name containing the config file
				dir := filepath.Dir(path)
				result = dir

				return filepath.SkipDir // Skip further traversal
			}

			return nil
		},
	)

	if result != "" {
		return result, nil
	}

	return "", errors.WrapIf(err, "No directory with config file found")
}

func searchRootDirectory(
	dir string,
) (string, error) {
	// List files and directories in the current directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", errors.WrapIf(err, "Error reading directory")
	}

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			if strings.EqualFold(
				fileName,
				"go.mod",
			) {
				return dir, nil
			}
		}
	}

	// If no config file found in this directory, recursively search its parent
	parentDir := filepath.Dir(dir)
	if parentDir == dir {
		// We've reached the root directory, and no go.mod file was found
		return "", errors.WrapIf(err, "No go.mod file found")
	}

	return searchRootDirectory(parentDir)
}
