package tsyaml

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// A testingtype interface implements Errorf for T, B and F.
// The interface enables generic functions for all test types T, B and F.
type testingtype interface {
	*testing.T | *testing.B | *testing.F
	Errorf(format string, a ...any)
	Fatalf(format string, a ...any)
}

type icheck interface {
	string | uint | int
}

func check[T icheck](want T, test T) error {
	if test != want {
		return fmt.Errorf("expected %v but received %v", want, test)
	}
	return nil
}

func errGet(key string, err error) error {
	return fmt.Errorf("get %T for key %v failed: %w", key, key, err)
}

func errExp(key string) error {
	return fmt.Errorf("expected error, but no error received for key %v", key)
}

func tmpYaml[T testingtype](tt T) {
	tmpYamlInit(tt)
	f := tmpYamlCreate(tt)
	if err := tmpYamlRead(tt, f); err != nil {
		tt.Fatalf("read in config of %v failed: %v", f, err)
	}
}

func tmpYamlRead[T testingtype](tt T, f string) error {
	fn := filepath.Base(f)
	return ReadInConfig(fn)
}

func tmpYamlCreate[T testingtype](tt T) string {
	// Create temp log file tsyaml_test_* in the temp directory
	f, err := os.CreateTemp(os.TempDir(), "tsyaml_test_*.yaml")
	if err != nil {
		f.Close()
		tt.Fatalf("creating %v failed: %v", f.Name(), err)
	}
	if _, err := f.WriteString(tcYaml); err != nil {
		f.Close()
		tt.Fatalf("writing test yaml file %v failed: %v", f.Name(), err)
	}
	if err := f.Close(); err != nil {
		tt.Fatalf("closing test yaml file %v failed: %v", f.Name(), err)
	}
	return f.Name()
}

func tmpYamlInit[T testingtype](tt T) {
	// Set TS_YAMLPATH to temp directory and re-initialize yaml path
	if err := os.Setenv(envn, os.TempDir()); err != nil {
		tt.Fatalf("setting env variable %v to %v failed: %v", envn, os.TempDir(), err)
	}
	initialize()
}
