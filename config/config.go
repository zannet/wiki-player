package config

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/goinggo/tracelog"
)

type (
	singleton struct {
		once  sync.Once
		value config
	}

	config struct {
		ConfigMap map[string]string // The map of strap key value pairs
	}

	XMLStrap struct {
		XMLName xml.Name `xml:"strap"`
		Key     string   `xml:"key,attr"`
		Value   string   `xml:"value,attr"`
	}

	XMLStraps struct {
		XMLName xml.Name   `xml:"straps"`
		Straps  []XMLStrap `xml:"strap"`
	}
)

var ci singleton

func MustLoad() {
	ci.once.Do(func() {
		// Find the location of the config.xml file
		configFilePath, err := filepath.Abs("config/config.xml")

		// Open the config.xml file
		file, err := os.Open(configFilePath)
		if err != nil {
			tracelog.CompletedError(err, "MustLoad", "os.Open")
			panic(err.Error())
		}

		defer file.Close()

		// Read the config file
		entries, err := readStraps(file)
		if err != nil {
			panic(err.Error())
		}

		// Create a config object
		ci.value = config{
			ConfigMap: make(map[string]string),
		}

		//Store the key/value pairs for each strap
		for _, entry := range entries {
			ci.value.ConfigMap[entry.Key] = entry.Value
		}
	})
}

func readStraps(reader io.Reader) ([]XMLStrap, error) {
	var xmlStraps XMLStraps
	if err := xml.NewDecoder(reader).Decode(&xmlStraps); err != nil {
		tracelog.CompletedError(err, "readStraps", "xml.NewDecoder")
		return nil, err
	}

	return xmlStraps.Straps, nil
}

func Entry(key string) string {
	MustLoad()
	return ci.value.ConfigMap[key]
}

