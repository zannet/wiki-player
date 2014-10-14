package utils

import (
	"encoding/json"
	"os"
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
		// Open the config.json file
		// Note: filepath.Abs() doesn't work for some reason when the app is ran
		// from tests so I'm just hardcoding the abs path of config.json
		file, err := os.Open("/home/adred/Golang/src/github.com/adred/wiki-player/config/config.json")
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

// ConfigEntry returns a config value
func ConfigEntry(key string) string {
	MustLoadConfig()
	return ci.value.ConfigMap[key]
}
