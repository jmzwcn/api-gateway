package loader

import (
	"api-gateway/common"
	"api-gateway/types"
	"encoding/json"
)

var RuleStore = make(types.RuleStore)

func ParseAndLoad() {
	load()
}

func load() {
	log.Debug(PROTO_JSON)
	var methods []types.MethodWrapper
	err := json.Unmarshal([]byte(string(PROTO_JSON)), &methods)
	if err != nil {
		log.Error(err)
	}

	for _, md := range methods {
		key := md.Pattern.Verb + ":" + md.Pattern.Path
		log.Debug(key, md)
		RuleStore[key] = md
	}
}
