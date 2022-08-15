package tsyaml

import (
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

func ReadInConfig(cn string) {

	viper.SetConfigName(cn)
	err := viper.ReadInConfig()
	if err != nil {
		tslog.E.Printf("config name %v: %v", cn, err)
	}
}

func GetStr(key string) string {
	v := get(key).(string)
	return v
}

func GetUInt(key string) uint {
	v := get(key).(uint)
	return v
}

func get(key string) interface{} {
	if key == "" {
		tslog.E.Println("attribute name cannot be empty")
		return nil
	}
	v := viper.Get(key)
	if v == nil {
		tslog.E.Printf("did not find key %v", key)
	}
	return v
}
