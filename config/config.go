package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sync"

	"github.com/goinggo/tracelog"
)

type (
	singleton struct {
		once  sync.Once
		value configMap
	}

	configMap struct {
		ConfigMap map[string]string
	}

	config struct {
		Host     string
		Database string
		Driver   string
		Username string
		Password string
		LogDir   string
	}
)

var ci singleton

func MustLoad() {
	ci.once.Do(func() {
		// Find the location of the config.json file
		configFilePath, err := filepath.Abs("config/config.json")

		// Open the config.json file
		file, err := os.Open(configFilePath)
		if err != nil {
			tracelog.CompletedError(err, "MustLoad", "os.Open")
			panic(err.Error())
		}
		defer file.Close()

		// Read the config file
		decoder := json.NewDecoder(file)
		c := &config{}
		err = decoder.Decode(&c)
		if err != nil {
			panic(err.Error())
		}

		// Create a configMap object
		ci.value = configMap{ConfigMap: make(map[string]string)}

		// Assign config field:value pairs to ConfigMap
		v := reflect.ValueOf(c).Elem()
		vType := v.Type()
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			ci.value.ConfigMap[vType.Field(i).Name] = f.Interface().(string)
		}
	})
}

func Entry(key string) string {
	MustLoad()
	return ci.value.ConfigMap[key]
}
