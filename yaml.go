package tsyaml

import (
	"fmt"
	"os"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/thorstenrie/tslog"
)

const (
	envn string = "TS_YAMLPATH"
)

func init() {
	initialize()
}

func initialize() {
	path, isset := os.LookupEnv(envn)

	if !isset {
		path = "."
	}

	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
}

func ReadInConfig(cn string) error {
	viper.SetConfigName(cn)
	err := viper.ReadInConfig()
	if err != nil {
		tslog.E.Printf("Read in %v failed: %v", cn, err)
	}
	return err
}

func GetStr(key string) (string, error) {
	v, err := get(key)
	if err != nil {
		return "", err
	}
	return cast.ToStringE(v)
}

func GetInt(key string) (int, error) {
	v, err := get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToIntE(v)
}

func GetUInt(key string) (uint, error) {
	v, err := get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToUintE(v)
}

func get(key string) (interface{}, error) {
	var err error = nil
	if key == "" {
		err = fmt.Errorf("attribute name cannot be empty")
		tslog.E.Println(err)
		return nil, err
	}
	v := viper.Get(key)
	if v == nil {
		err = fmt.Errorf("did not find key %v", key)
		tslog.I.Println(err)
	}
	return v, err
}
