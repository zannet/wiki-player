package utils

import (
	"encoding/json"
	"fmt"
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
		Host        string
		Database    string
		Driver      string
		Username    string
		Password    string
		LogDir      string
		Salt        string
		SecretKey   string
		SessionName string
		StaticDir   string
	}
)

var ci configSingleton

// MustLoadConfig loads the configs and assign it to configMap
func MustLoadConfig() {
	ci.once.Do(func() {
		// Get app mode -- test/development/production
		mode := EnvConfigEntry("Mode")
		fmt.Println(mode)
		// Find the location of the config.json file
		configFilePath, err := filepath.Abs("app/configs/" + mode + ".json")

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
			tracelog.CompletedError(err, "MustLoadConfig", "decoder.Decode")
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

// ConfigEntry returns a config value
func ConfigEntry(key string) string {
	MustLoadConfig()
	return ci.value.ConfigMap[key]
}
