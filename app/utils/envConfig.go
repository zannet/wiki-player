package utils

import (
    "reflect"
    "sync"

    "github.com/goinggo/tracelog"
    "github.com/kelseyhightower/envconfig"
)

type (
    envConfigSingleton struct {
        once  sync.Once
        value envConfigMap
    }

    envConfigMap struct {
        EnvConfigMap map[string]string
    }

    envConfig struct {
        Mode string
        Branch string
    }
)

var ec envConfigSingleton

// MustLoadEnvConfig loads the configs and assign it to envConfigMap
func MustLoadEnvConfig() {
    ec.once.Do(func() {
        var e envConfig
        err := envconfig.Process("app", &e)

        if err != nil {
            tracelog.CompletedError(err, "MustLoadConfig", "envconfig.Process")
            panic(err.Error())
        }

        // Create a envConfigMap object
        ec.value = envConfigMap{EnvConfigMap: make(map[string]string)}

        // Assign config field:value pairs to EnvConfigMap
        v := reflect.ValueOf(&e).Elem()
        vType := v.Type()
        for i := 0; i < v.NumField(); i++ {
            f := v.Field(i)
            ec.value.EnvConfigMap[vType.Field(i).Name] = f.Interface().(string)
        }
    })
}

// EnvConfigEntry returns a config value
func EnvConfigEntry(key string) string {
    MustLoadEnvConfig()
    return ec.value.EnvConfigMap[key]
}
