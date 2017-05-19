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
	Port     string
	ProtoSet []Proto `json:"proto.set"`
}

type Proto struct {
	Service string
	Path    string
}

func NewConfiguration() *Configuration {
	conf := Configuration{}
	conf.parse()
	return &conf
}
func (config *Configuration) parse() {
	configFile, err := os.Open(DEFAULT_CONFIG_FILE_NAME)
	if err != nil {
		config.Port = "8080"
		log.Error(err, ", will use the default.")
	} else {
		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(config); err != nil {
			log.Error(err)
		}
	}
	log.Debug("Current config:", *config)
}
