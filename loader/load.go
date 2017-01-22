package loader

import (
	"api-gateway/types"
	"encoding/json"
	"log"
)

var RuleStore = types.CreateRuleStore()

func ParseAndLoad() {
	load()
}

func load() {
	log.Println(PROTO_JSON)
	var methods []types.MethodWrapper
	err := json.Unmarshal([]byte(string(PROTO_JSON)), &methods)
	if err != nil {
		log.Println(err)
	}

	for _, md := range methods {
		key := md.Pattern.Verb + ":" + md.Pattern.Path
		log.Println(key, md)
		RuleStore.Store[key] = md
	}
}
