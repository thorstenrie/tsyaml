package tsyaml

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// A testingtype interface implements Errorf and Fatalf for T, B and F.
// The interface is used generic functions for testing types T, B and F.
type testingtype interface {
	*testing.T | *testing.B | *testing.F
	Errorf(format string, a ...any)
	Fatalf(format string, a ...any)
}

// An icheck interface is constrained to the value types associated with a key.
type icheck interface {
	string | uint | int
}

// check evaluates, if want and test are equal. If not, the function returns an error.
func check[T icheck](want T, test T) error {
	if test != want {
		return fmt.Errorf("expected %v but received %v", want, test)
	}
	return nil
}

// errGet has the purpose to return an error for cases when get for a key fails with err.
func errGet(key string, err error) error {
	return fmt.Errorf("get %T for key %v failed: %w", key, key, err)
}

// errExp has the purpose to return an error for cases when an error is expected
// but not received.
func errExp(key string) error {
	return fmt.Errorf("expected error, but no error received for key %v", key)
}

// errRd has the purpose to return an error for cases when reading a yaml file fails.
func errRd(f string, err error) error {
	return fmt.Errorf("read in config of %v failed: %v", f, err)
}

// tmpYaml creates a temp yaml file in the temporary directory. The yaml file
// contains the defined testcase tcYaml. The config path is set to the temporary
// directory and the temp yaml file is read in.
func tmpYaml[T testingtype](tt T) {
	// Set config path to the temporary directory
	tmpYamlInit(tt)
	// Create a temporary yaml file containing the testcase tcYaml
	f := tmpYamlCreate(tt, tcYaml)
	// Read the temporary yaml file in or log an error followed by FailNow
	if err := tmpYamlRead(tt, f); err != nil {
		tt.Fatalf("read in config of %v failed: %v", f, err)
	}
}

func tmpYamlRead[T testingtype](tt T, f string) error {
	fn := filepath.Base(f)
	return ReadInConfig(fn)
}

func tmpYamlCreate[T testingtype](tt T, tc string) string {
	// Create temp file tsyaml_test_* in the temp directory
	f, err := os.CreateTemp(os.TempDir(), "tsyaml_test_*.yaml")
	if err != nil {
		f.Close()
		tt.Fatalf("creating %v failed: %v", f.Name(), err)
	}
	if _, err := f.WriteString(tc); err != nil {
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
