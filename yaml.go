// Package tsyaml provides a simple API to read from YAML files.
//
// The tsyaml package is based on github.com/spf13/viper. It is a
// wrapper to provide a simple API to read from YAML files. The config
// path containing the yaml file(s) is set by the env variable
// TS_YAMLPATH during the initial startup of the app. The package
// contains a function to read a config file and three get functions
// to receive values of types int, uint and string.
//
// Copyright (c) 2022 thorstenrie
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsyaml

// Import standard library packages, spf13/viper and spf13/cast.
import (
	"fmt" // fmt
	"os"  // os

	"github.com/spf13/cast"  // cast
	"github.com/spf13/viper" // viper
)

// The config path is set by the env variable TS_YAMLPATH during the initial
// startup of the app.
const (
	envn string = "TS_YAMLPATH"
)

// init initializes the config path.
func init() {
	initialize()
}

// initialize sets the config path containing the YAML file(s) with the env
// variable TS_YAMLPATH. If TS_YAMLPATH is not set, then it sets the config
// path to ".".
func initialize() {
	// Look up the env variable
	path, isset := os.LookupEnv(envn)
	// Check if the env variable is set
	if !isset {
		// If not set, then initialize config path with "."
		path = "."
	}
	// Add config path
	viper.AddConfigPath(path)
	// Set config type to yaml
	viper.SetConfigType("yaml")
}

// ReadInConfig reads yaml file cn. It returns an error
// if reading fails.
func ReadInConfig(cn string) error {
	// Set config name to cn
	viper.SetConfigName(cn)
	// Read in config
	err := viper.ReadInConfig()
	// Return error, if not nil
	if err != nil {
		return fmt.Errorf("read in %v failed with %v", cn, err)
	}
	// Return nil
	return err
}

// GetStr returns the value as a string, associated with the key.
// In case of an error, the function returns "" and an error.
func GetStr(key string) (string, error) {
	v, err := get(key)
	if err != nil {
		return "", err
	}
	return cast.ToStringE(v)
}

// GetInt returns the value as an integer, associated with the key.
// In case of an error, the function returns 0 and an error.
func GetInt(key string) (int, error) {
	v, err := get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToIntE(v)
}

// GetUint returns the value as an unsigned integer, associated with the key.
// If it fails, the function returns 0 and an error.
func GetUint(key string) (uint, error) {
	v, err := get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToUintE(v)
}

// get returns the value associated to the key.
// If it fails, it returns nil and an error.
func get(key string) (interface{}, error) {
	// Initialize err with nil
	var err error = nil
	// Check for empty key ""
	if key == "" {
		err = fmt.Errorf("key name cannot be empty")
		return nil, err
	}
	// Get value for key
	v := viper.Get(key)
	// Check if value is nil
	if v == nil {
		err = fmt.Errorf("did not find key %v", key)
	}
	// Return value and err
	return v, err
}
