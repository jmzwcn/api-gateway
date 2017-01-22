package config

import (
	"encoding/json"
	"log"
	"os"
)

const (
	DEFAULT_CONFIG_FILE_NAME = "config.json"
)

type Configuration struct {
	Port string
}

func NewConfiguration() *Configuration {
	conf := Configuration{}
	conf.parse()
	return &conf
}
func (config *Configuration) parse() {
	configFile, err := os.Open(DEFAULT_CONFIG_FILE_NAME)
	if err != nil {
		log.Println(err)
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(config); err != nil {
		log.Println(err)
	}
}
