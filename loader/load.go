package loader

import (
	"encoding/json"

	"github.com/api-gateway/types"
	"google.golang.org/grpc/grpclog"
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
		grpclog.Error(err)
	}

	for _, md := range methods {
		key := md.Pattern.Verb + ":" + md.Pattern.Path
		//key := md.Pattern.Verb + ":/" + md.Package + md.Pattern.Path
		grpclog.Println(key, "->", md)
		RuleStore[key] = md
	}
}
