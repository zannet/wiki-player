package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sync"

	"github.com/goinggo/tracelog"
)

type (
	configSingleton struct {
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
		Salt     string
	}
)

var ci configSingleton

func MustLoadConfig() {
	ci.once.Do(func() {
		// Find the location of the config.json file
		configFilePath, err := filepath.Abs("config/config.json")

		// Open the config.json file
		file, err := os.Open(configFilePath)
		if err != nil {
			tracelog.CompletedError(err, "MustLoadConfig", "os.Open")
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

func ConfigEntry(key string) string {
	MustLoadConfig()
	return ci.value.ConfigMap[key]
}
