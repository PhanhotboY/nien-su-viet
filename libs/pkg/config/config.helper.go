package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
	typeMapper "github.com/phanhotboy/nien-su-viet/libs/pkg/reflection/typemapper"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"github.com/mcuadros/go-defaults"
)

func BindConfig[T any](environments ...environment.Environment) (T, error) {
	return BindConfigKey[T]("", environments...)
}

func BindConfigKey[T any](
	configKey string,
	environments ...environment.Environment,
) (T, error) {

	cfg := typeMapper.GenericInstanceByT[T]()

	// Set defaults before loading environment variables
	// https://github.com/mcuadros/go-defaults
	defaults.SetDefaults(cfg)

	// Load .env files recursively from current working directory up the tree
	if err := loadEnvFilesRecursive(); err != nil {
		log.Printf("Warning: Could not load .env files: %v", err)
		// Continue without .env file, environment variables may still be set
	}

	// Parse configuration from environment variables
	// https://github.com/caarlos0/env
	opts := env.Options{
		RequiredIfNoDef: false,
	}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		fmt.Printf("Warning parsing environment variables: %+v\n", err)
	}

	return cfg, nil
}

// loadEnvFilesRecursive loads .env files starting from the current working directory
// and moving up the directory tree until a .env file is found.
func loadEnvFilesRecursive() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Keep traversing up the directory hierarchy until you find an ".env" file
	for {
		envFilePath := filepath.Join(dir, ".env")
		err := godotenv.Load(envFilePath)
		if err == nil {
			// .env file found and loaded successfully
			log.Printf("Loaded environment variables from: %s", envFilePath)
			return nil
		}

		// Move to parent directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root directory
			break
		}
		dir = parent
	}

	log.Printf("No .env file found in directory hierarchy")
	return nil
}
