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

func TestStr(t *testing.T) {
	tmpYaml(t)
	if testStr, errg := GetStr(keyStr); errg != nil {
		t.Error(errGet(keyStr, errg))
	} else if errc := check(testStr, wantStr); errc != nil {
		t.Error(errc)
	}
}

func BenchmarkYaml(b *testing.B) {
	tmpYamlInit(b)
	f := []string{filepath.Base(tmpYamlCreate(b, tcYaml)), filepath.Base(tmpYamlCreate(b, tcYaml))}
	// Reset benchmark timer
	b.ResetTimer()
	// Run benchmark with all testcases in each iteration
	for i := 0; i < b.N; i++ {
		fn := f[i&0x1]
		if err := ReadInConfig(fn); err != nil {
			b.Fatal(errRd(fn, err))
		}
		if _, errg := GetStr(keyStr); errg != nil {
			b.Error(errGet(keyStr, errg))
		}
		if _, errg := GetUint(keyUint); errg != nil {
			b.Error(errGet(keyUint, errg))
		}
		if _, errg := GetInt(keyInt); errg != nil {
			b.Error(errGet(keyInt, errg))
		}
		keynested := keyN.root + "." + keyN.leaf
		if _, errg := GetStr(keynested); errg != nil {
			b.Error(errGet(keynested, errg))
		}
	}
}

func TestInvalidKeyStr(t *testing.T) {
	tmpYaml(t)
	keyRev := revStr(keyStr)
	if _, errg := GetStr(keyRev); errg == nil {
		t.Error(errExp(keyRev))
	}
}

func TestInvalidKeyInt(t *testing.T) {
	tmpYaml(t)
	keyRev := revStr(keyInt)
	if _, errg := GetInt(keyRev); errg == nil {
		t.Error(errExp(keyRev))
	}
}

func TestInvalidKeyUint(t *testing.T) {
	tmpYaml(t)
	keyRev := revStr(keyUint)
	if _, errg := GetUint(keyRev); errg == nil {
		t.Error(errExp(keyRev))
	}
}

func TestEmptyKey(t *testing.T) {
	tmpYaml(t)
	if _, errg := GetStr(""); errg == nil {
		t.Error(errExp("empty"))
	}
}

func TestUint(t *testing.T) {
	tmpYaml(t)
	if testUint, errg := GetUint(keyUint); errg != nil {
		t.Error(errGet(keyUint, errg))
	} else if errc := check(testUint, wantUint); errc != nil {
		t.Error(errc)
	}
}

func TestInt(t *testing.T) {
	tmpYaml(t)
	if testInt, errg := GetInt(keyInt); errg != nil {
		t.Error(errGet(keyInt, errg))
	} else if errc := check(testInt, wantInt); errc != nil {
		t.Error(errc)
	}
}

func TestNested(t *testing.T) {
	tmpYaml(t)
	keynested := keyN.root + "." + keyN.leaf
	if testLeaf, errg := GetStr(keynested); errg != nil {
		t.Error(errGet(keynested, errg))
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

func TestInvalidYaml(t *testing.T) {
	tmpYamlInit(t)
	f := tmpYamlCreate(t, tcYaml)
	if err := os.Remove(f); err != nil {
		t.Fatalf("removing %v failed: %v", f, err)
	}
	if err := ReadInConfig(filepath.Base(f)); err == nil {
		t.Errorf("Expected error, but no error received for config file %v", f)
	}
}
