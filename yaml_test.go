package tsyaml

// Import standard library packages
import (
	"fmt"           // fmt
	"os"            // os
	"path/filepath" // path/filepath
	"testing"       // testing
)

// The testcase comprises keys for strings, uint and int.
const (
	keyStr    string = "testStr"    // string
	keyStrAlt string = "testStrAlt" // alternate string
	keyUint   string = "testUint"   // uint
	keyInt    string = "testInt"    // int
)

// The testcases contain values for strings, uint and int.
const (
	wantStr    string = "Hello World!" // string
	wantStrAlt string = "Alternative"  // alternate string
	wantUint   uint   = 12345          // uint
	wantInt    int    = -67890         // int
)

// nested for testing a simple struct with root and one leaf.
type nested struct {
	root, leaf string
}

// Nested testcase
var (
	keyN  = nested{root: "location", leaf: "Country/City"} // keys
	wantN = nested{root: "", leaf: "Germany/Heidelberg"}   // values
)

// tcYaml string contains the entire testcase which is written to a yaml file
var (
	tcYaml string = keyStr + ": " + wantStr + "\n" +
		keyStrAlt + ": " + wantStrAlt + "\n" +
		keyUint + ": " + fmt.Sprintf("%v", wantUint) + "\n" +
		keyInt + ": " + fmt.Sprintf("%v", wantInt) + "\n" +
		keyN.root + ":\n" +
		"    " + keyN.leaf + ": " + wantN.leaf + "\n"
)

// TestStr writes a testcase yaml file, reads it and retrieves a string value.
// Expected result is that the retrieved string matches the wanted string.
func TestStr(t *testing.T) {
	// Create yaml file with testcase in temporary directory and read it in
	tmpYaml(t)
	// Get string for corresponding key in testcase
	if testStr, errg := GetStr(keyStr); errg != nil {
		t.Error(errGet(keyStr, errg))
	} else if errc := check(testStr, wantStr); errc != nil { // Compare retrieved value with wanted value
		t.Error(errc)
	}
}

// TestUint writes a testcase yaml file, reads it and retrieves an uint value.
// Expected result is that the retrieved value matches the wanted value.
func TestUint(t *testing.T) {
	// Create yaml file with testcase in temporary directory and read it in
	tmpYaml(t)
	// Get uint for corresponding key in testcase
	if testUint, errg := GetUint(keyUint); errg != nil {
		t.Error(errGet(keyUint, errg))
	} else if errc := check(testUint, wantUint); errc != nil { // Compare retrieved value with wanted value
		t.Error(errc)
	}
}

// TestInt writes a testcase yaml file, reads it and retrieves an int value.
// Expected result is that the retrieved value matches the wanted value.
func TestInt(t *testing.T) {
	// Create yaml file with testcase in temporary directory and read it in
	tmpYaml(t)
	// Get int for corresponding key in testcase
	if testInt, errg := GetInt(keyInt); errg != nil {
		t.Error(errGet(keyInt, errg))
	} else if errc := check(testInt, wantInt); errc != nil { // Compare retrieved value with wanted value
		t.Error(errc)
	}
}

// TestNested writes a testcase yaml file, reads it and retrieves the nested value.
// Expected result is that the retrieved value matches the wanted value.
func TestNested(t *testing.T) {
	// Create yaml file with testcase in temporary directory and read it in
	tmpYaml(t)
	// Nested key
	keynested := keyN.root + "." + keyN.leaf
	// Get string for corresponding nested key in testcase
	if testLeaf, errg := GetStr(keynested); errg != nil {
		t.Error(errGet(keynested, errg))
	} else if errc := check(testLeaf, wantN.leaf); errc != nil { // Compare retrieved value with wanted value
		t.Error(errc)
	}
}

// TestInvalidKeyStr writes a testcase yaml file, reads it and tries to retrieve
// a non-existing string key. Expected result is an error.
func TestInvalidKeyStr(t *testing.T) {
	// Create yaml file with testcase in temporary directory and read it in
	tmpYaml(t)
	// Reverse key
	keyRev := revStr(keyStr)
	// Get string for reversed key
	if _, errg := GetStr(keyRev); errg == nil {
		t.Error(errExp(keyRev))
	}
}

// TestInvalidKeyInt writes a testcase yaml file, reads it and tries to retrieve
// a non-existing int key. Expected result is an error.
func TestInvalidKeyInt(t *testing.T) {
	// Create yaml file with testcase in temporary directory and read it in
	tmpYaml(t)
	// Reverse key
	keyRev := revStr(keyInt)
	// Get int for reversed key
	if _, errg := GetInt(keyRev); errg == nil {
		t.Error(errExp(keyRev))
	}
}

// TestInvalidKeyUint writes a testcase yaml file, reads it and tries to retrieve
// a non-existing uint key. Expected result is an error.
func TestInvalidKeyUint(t *testing.T) {
	// Create yaml file with testcase in temporary directory and read it in
	tmpYaml(t)
	// Reverse key
	keyRev := revStr(keyUint)
	// Get uint for reversed key
	if _, errg := GetUint(keyRev); errg == nil {
		t.Error(errExp(keyRev))
	}
}

// TestEmptyKey writes a testcase yaml file, reads it and tries to retrieve
// an empty key. Expected result is an error.
func TestEmptyKey(t *testing.T) {
	// Create yaml file with testcase in temporary directory and read it in
	tmpYaml(t)
	// Get string for empty key
	if _, errg := GetStr(""); errg == nil {
		t.Error(errExp("empty"))
	}
}

// TestInvalidYaml writes a testcase yaml file, removes it, and tries to read it.
// Expected result is an error.
func TestInvalidYaml(t *testing.T) {
	// Set config path to the temporary directory
	tmpYamlInit(t)
	// Create a yaml file in the temporary directory containing the testcase tcYaml
	f := tmpYamlCreate(t, tcYaml)
	// Remove the yaml file
	if err := os.Remove(f); err != nil {
		t.Fatalf("removing %v failed: %v", f, err)
	}
	// Read in yaml file
	if err := ReadInConfig(filepath.Base(f)); err == nil {
		t.Errorf("Expected error, but no error received for config file %v", f)
	}
}

// BenchmarkYaml performs a benchmark reading the testcase from a yaml file and
// retrieving four values of type string, uint, int, nested
func BenchmarkYaml(b *testing.B) {
	// Set config path to the temporary directory
	tmpYamlInit(b)
	// Create two yaml files in the temporary directory containing the testcase tcYaml
	f := []string{filepath.Base(tmpYamlCreate(b, tcYaml)), filepath.Base(tmpYamlCreate(b, tcYaml))}
	// Reset benchmark timer
	b.ResetTimer()
	// Run benchmark with all testcases in each iteration
	for i := 0; i < b.N; i++ {
		// Alternate the yaml file in each iteration
		fn := f[i&0x1]
		// Read yaml file
		if err := ReadInConfig(fn); err != nil {
			b.Fatal(errRd(fn, err))
		}
		// Retrieve string
		if _, errg := GetStr(keyStr); errg != nil {
			b.Error(errGet(keyStr, errg))
		}
		// Retrieve uint
		if _, errg := GetUint(keyUint); errg != nil {
			b.Error(errGet(keyUint, errg))
		}
		// Retrieve int
		if _, errg := GetInt(keyInt); errg != nil {
			b.Error(errGet(keyInt, errg))
		}
		// Retrieve nested
		keynested := keyN.root + "." + keyN.leaf
		if _, errg := GetStr(keynested); errg != nil {
			b.Error(errGet(keynested, errg))
		}
	}
}

// revStr reverses string s and returns the reversed string as string
func revStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
