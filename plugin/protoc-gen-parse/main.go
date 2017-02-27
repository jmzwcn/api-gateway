package main

import (
	"api-gateway/common"
	options "api-gateway/third_party/google/api"
	"api-gateway/types"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Error("error reading input")
	}

	request := new(plugin.CodeGeneratorRequest)
	if err := proto.Unmarshal(data, request); err != nil {
		log.Error(err)
	}

	var methods []types.MethodWrapper
	for _, allProtoBuff := range request.GetProtoFile() {
		for _, generateProtoBuff := range request.FileToGenerate {
			if *allProtoBuff.Name == generateProtoBuff {
				for _, service := range allProtoBuff.Service {
					for _, md := range service.Method {
						ext, err := proto.GetExtension(md.Options, options.E_Http)
						if err == nil {
							//exts = append(exts, ext)
							method := types.MethodWrapper{}
							method.Package = *allProtoBuff.Package
							method.Service = *service.Name
							method.Method = md

							pattern := types.Pattern{}
							rule := ext.(*options.HttpRule)
							pattern.Verb = getVerb(rule)
							pattern.Path = rule.GetGet() + rule.GetPost() + rule.GetPut() + rule.GetDelete()
							pattern.Body = rule.Body
							method.Pattern = &pattern
							methods = append(methods, method)
						}
					}
				}
			}
		}
	}

	jsonOut, err := json.Marshal(methods)
	if err != nil {
		log.Error("json.Marshal eror", err)
	}
	f, _ := os.Create("parse.json")
	str := strconv.Quote(string(jsonOut))
	f.WriteString(str)
}

func getVerb(opts *options.HttpRule) string {
	var httpMethod string
	switch {
	case opts.GetGet() != "":
		httpMethod = "GET"
	case opts.GetPost() != "":
		httpMethod = "POST"
	case opts.GetPut() != "":
		httpMethod = "PUT"
	case opts.GetDelete() != "":
		httpMethod = "DELETE"
	case opts.GetPatch() != "":
		httpMethod = "PATCH"
	}
	return httpMethod
}
