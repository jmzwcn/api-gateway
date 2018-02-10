package loader

import (
	"encoding/json"
	"log"

	"github.com/api-gateway/types"
)

var RuleStore = make(types.RuleStore)

func Services() {
	load()
}

func load() {
	//log.Debug(PROTO_JSON)
	var methods []types.MethodWrapper
	err := json.Unmarshal([]byte(string(PROTO_JSON)), &methods)
	if err != nil {
		log.Panicln(err)
	}

	for _, md := range methods {
		key := md.Pattern.Verb + ":" + md.Pattern.Path
		//key := md.Pattern.Verb + ":/" + md.Package + md.Pattern.Path
		log.Println(key, "->", md)
		RuleStore[key] = md
	}
}
