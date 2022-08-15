package tsyaml

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/thorstenrie/tslog"
)

func init() {

	dirname, isset := os.LookupEnv("TS_YAMLDIR")

	if !isset {
		dirname = "."
	}

	viper.AddConfigPath(dirname)
	viper.SetConfigType("yaml")
}

func ReadInConfig(cn string) error {
	viper.SetConfigName(cn)
	err := viper.ReadInConfig()
	if err != nil {
		tslog.E.Printf("config name %v: %v", cn, err)
	}
	return err
}

func GetStr(key string) (string, error) {
	v, err := get(key)
	return v.(string), err
}

func GetUInt(key string) (uint, error) {
	v, err := get(key)
	return v.(uint), err
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
