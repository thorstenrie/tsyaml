package tsyaml

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	keyStr    string = "testStr"
	keyStrAlt string = "testStrAlt"
	keyUint   string = "testUint"
	keyInt    string = "testInt"
)

const (
	wantStr    string = "Hello World!"
	wantStrAlt string = "Alternative"
	wantUint   uint   = 12345
	wantInt    int    = -67890
)

type nested struct {
	root, leaf string
}

var (
	keyN  = nested{root: "location", leaf: "Country/City"}
	wantN = nested{root: "", leaf: "Germany/Heidelberg"}
)

var (
	tcYaml string = keyStr + ": " + wantStr + "\n" +
		keyStrAlt + ": " + wantStrAlt + "\n" +
		keyUint + ": " + fmt.Sprintf("%v", wantUint) + "\n" +
		keyInt + ": " + fmt.Sprintf("%v", wantInt) + "\n" +
		keyN.root + ":\n" +
		"    " + keyN.leaf + ": " + wantN.leaf + "\n"
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

func TestStr(t *testing.T) {
	tmpYaml(t)
	if testStr, errg := GetStr(keyStr); errg != nil {
		t.Errorf("Get string for key %v failed: %v", keyStr, errg)
	} else if errc := check(testStr, wantStr); errc != nil {
		t.Error(errc)
	}
}

func TestInvalidKey(t *testing.T) { // TODO
	tmpYaml(t)
	keyRev := revStr(keyStr)
	if testStr, errg := GetStr(keyRev); errg != nil {
		t.Errorf("Get string for key %v failed: %v", keyRev, errg)
	} else if errc := check(testStr, wantStr); errc != nil {
		t.Error(errc)
	}
}

func TestInvalidValue(t *testing.T) { // TODO
	tmpYaml(t)
	if testStr, errg := GetStr(keyStrAlt); errg != nil {
		t.Errorf("Get string for key %v failed: %v", keyStrAlt, errg)
	} else if errc := check(testStr, wantStr); errc != nil {
		t.Error(errc)
	}
}

func TestUint(t *testing.T) {
	tmpYaml(t)
	if testUint, errg := GetUInt(keyUint); errg != nil {
		t.Errorf("Get uint for key %v failed: %v", keyUint, errg)
	} else if errc := check(testUint, wantUint); errc != nil {
		t.Error(errc)
	}
}

func TestInt(t *testing.T) {
	tmpYaml(t)
	if testInt, errg := GetInt(keyInt); errg != nil {
		t.Errorf("Get int for key %v failed: %v", keyInt, errg)
	} else if errc := check(testInt, wantInt); errc != nil {
		t.Error(errc)
	}
}

func TestNested(t *testing.T) {
	tmpYaml(t)
	keynested := keyN.root + "." + keyN.leaf
	if testLeaf, errg := GetStr(keynested); errg != nil {
		t.Errorf("Get int for key %v failed: %v", keynested, errg)
	} else if errc := check(testLeaf, wantN.leaf); errc != nil {
		t.Error(errc)
	}
}

func revStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func tmpYaml[T testingtype](tt T) {
	// Create temp log file tsyaml_test_* in the temp directory
	f, err := os.CreateTemp(os.TempDir(), "tsyaml_test_*.yaml")
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
	fn := filepath.Base(f.Name())
	if err := ReadInConfig(fn); err != nil {
		tt.Fatalf("Read in config of %v failed: %v", fn, err)
	}
}
