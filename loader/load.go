package loader

import (
	"encoding/json"

	"github.com/api-gateway/types"
	"github.com/api-gateway/types/log"
)

var RuleStore = make(types.RuleStore)

func init() {
	//log.Debug(PROTO_JSON)
	var methods []types.MethodWrapper
	err := json.Unmarshal([]byte(string(PROTO_JSON)), &methods)
	if err != nil {
		log.Fatal(err)
	}

	for _, md := range methods {
		key := md.Pattern.Verb + ":" + md.Pattern.Path
		//key := md.Pattern.Verb + ":/" + md.Package + md.Pattern.Path
		log.Println(key, "-->", md)
		RuleStore[key] = md
	}
}
