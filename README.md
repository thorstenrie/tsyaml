# tsyaml

[![Go Report Card](https://goreportcard.com/badge/github.com/thorstenrie/tsyaml)](https://goreportcard.com/report/github.com/thorstenrie/tsyaml)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorstenrie/tsyaml)](https://pkg.go.dev/mod/github.com/thorstenrie/tsyaml)

[Go](https://go.dev/) package to read from YAML files with a simple API ([KISS principle](https://en.wikipedia.org/wiki/KISS_principle)).

- **Configuration**: Configured with one environment variable, initialized at startup of your app
- **Usage**: Simple API to read in yaml files and retrieve values
- **Tested**: Unit tests with high code coverage
- **Dependencies**: Based on [github.com/spf13/viper](https://github.com/spf13/viper)

## Usage

Before app execution, set the environment variable `TS_YAMLPATH` to the yaml file directory (see [Configuration](#Configuration)).

E.g., in a linux terminal, run

```
export TS_YAMLPATH=./config
```

E.g., in VS Code, add to the `configuration` block:

```
"env": {
    "TS_YAMLPATH": "./config"
}
```

In the Go app, the package is imported with

```
import "github.com/thorstenrie/tsyaml"
```

Yaml files are read in by using [ReadInConfig](https://pkg.go.dev/github.com/thorstenrie/tsyaml#ReadInConfig):

```
tsyaml.ReadInConfig("example") // Read "example.yaml" in directory defined by TS_YAMLPATH
```

Values of associated keys are retrieved by using the get functions [GetStr](https://pkg.go.dev/github.com/thorstenrie/tsyaml#GetStr), [GetUint](https://pkg.go.dev/github.com/thorstenrie/tsyaml#GetUint), [GetInt](https://pkg.go.dev/github.com/thorstenrie/tsyaml#GetInt):

```
out1, err1 := tsyaml.GetStr("test")   // string
out2, err2 := tsyaml.GetInt("test2")  // int
out3, err3 := tsyaml.GetUint("test3") // uint
```

## Configuration

The tsyaml package reads in yaml files from the directory defined in the environment variable `TS_YAMLPATH`. If `TS_YAMLPATH` is not set, then tsyaml will use the default directory ".".

## Example

```
import (
	"fmt"

	"github.com/thorstenrie/tsyaml"
)

func main() {
	tsyaml.ReadInConfig("example")
	out, _ := tsyaml.GetStr("test")
	out2, _ := tsyaml.GetInt("test2")
	out3, _ := tsyaml.GetUint("test3")
	fmt.Println(out)
	fmt.Println(out2)
	fmt.Println(out3)
}
```

with `example.yaml`:

```
test: 1234asdf
test2: -1234
test3: 1234
```

## Links

[Godoc](https://pkg.go.dev/github.com/thorstenrie/tsyaml)

[Go Report Card](https://goreportcard.com/report/github.com/thorstenrie/tsyaml)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorstenrie%2Ftsyaml)
