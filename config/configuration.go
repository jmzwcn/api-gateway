package config

import (
	"api-gateway/common"
	"encoding/json"
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
		log.Error(err)
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(config); err != nil {
		log.Error(err)
	}
}
