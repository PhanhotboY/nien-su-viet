package chelper

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"emperror.dev/errors"
	"github.com/spf13/viper"
)

func SearchRootDirectory(
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

	return SearchRootDirectory(parentDir)
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

func GetEnvPrefix() string {
	envPrefix := os.Getenv("ENV_PREFIX")
	if envPrefix == "" {
		panic("missing ENV_PREFIX environment variable")
	}
	return strings.ToUpper(envPrefix)
}
