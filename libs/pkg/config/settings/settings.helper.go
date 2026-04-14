package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"emperror.dev/errors"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
	"github.com/spf13/viper"
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

func BindEnvs(v *viper.Viper, iface interface{}, prefix string) {
	t := reflect.TypeOf(iface)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		name, squash := parseTag(field.Tag.Get("mapstructure"), field.Name)

		var key string
		if prefix != "" && !squash {
			key = prefix + "." + name
		} else if squash {
			key = prefix
		} else {
			key = name
		}

		// If nested struct → recurse
		if field.Type.Kind() == reflect.Struct {
			BindEnvs(v, reflect.New(field.Type).Elem().Interface(), key)
			continue
		}

		v.SetDefault(key, nil) // registers key
		v.BindEnv(key)
	}
}
func parseTag(tag string, fieldName string) (name string, squash bool) {
	if tag == "" {
		return strings.ToLower(fieldName), false
	}

	parts := strings.Split(tag, ",")

	name = parts[0]
	if name == "" {
		name = strings.ToLower(fieldName)
	}

	for _, opt := range parts[1:] {
		if opt == "squash" {
			squash = true
		}
	}

	return
}
