package config

import (
	"encoding/json"
	"os"

	"github.com/api-gateway/common"
)

const (
	DEFAULT_CONFIG_FILE_NAME = "config.json"
)

type Configuration struct {
	comment  string
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
		log.Error(err)
	} else {
		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(config); err != nil {
			log.Error(err)
		}
	}
}
