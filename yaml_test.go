package tsyaml

import (
	"fmt"
	"os"
	"testing"
)

const (
	testStr   string = "Hello World!"
	testUint  uint   = 12345
	testInt   int    = -67890
	testCntry string = "Germany"
	testCty   string = "Heidelberg"
)

var (
	tcYaml string = "testStr: " + testStr + "\n" +
		"testUint: " + fmt.Sprintf("%v", testUint) + "\n" +
		"testInt: " + fmt.Sprintf("%v", testInt) + "\n" +
		"location:\n" +
		"    country: " + testCntry + "\n" +
		"    city: " + testCty + "\n"
)

type location struct {
	country, city string
}

// A testingtype interface implements Errorf for T, B and F.
// The interface enables generic functions for all test types T, B and F.
type testingtype interface {
	*testing.T | *testing.B | *testing.F
	Errorf(format string, a ...any)
	Fatalf(format string, a ...any)
}

func TestTmpYaml(t *testing.T) {
	// Todo
}

func tmpYaml[T testingtype](tt T) string {
	// Create temp log file tsyaml_test_* in the temp directory
	f, err := os.CreateTemp(os.TempDir(), "tsyaml_test_*")
	if err != nil {
		f.Close()
		tt.Fatalf("creating %v failed: %v", f.Name(), err)
	}
	// Set TS_YAMLPATH to temp directory and re-initialize yaml path
	if err := os.Setenv(envn, os.TempDir()); err != nil {
		f.Close()
		tt.Fatalf("setting env variable %v to %v failed: %v", envn, os.TempDir(), err)
	}
	if _, err := f.WriteString(tcYaml); err != nil {
		f.Close()
		tt.Fatalf("writing test yaml file %v failed: %v", f.Name(), err)
	}
	if err := f.Close(); err != nil {
		tt.Fatalf("closing test yaml file %v failed: %v", f.Name(), err)
	}
	initialize()
	return f.Name()
}
