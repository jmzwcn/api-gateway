package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/api-gateway/third_party/google/api"
	"github.com/api-gateway/third_party/runtime"
	"github.com/api-gateway/types"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panicln("error reading input")
	}

	request := new(plugin_go.CodeGeneratorRequest)
	if err := proto.Unmarshal(data, request); err != nil {
		log.Panicln(err)
	}

	var methods []types.MethodWrapper
	for _, protoFile := range request.GetProtoFile() {
		for _, generateFileName := range request.FileToGenerate {
			if *protoFile.Name == generateFileName {
				for _, service := range protoFile.Service {
					for _, md := range service.Method {
						method := types.MethodWrapper{}
						options := make(map[string]interface{})

						method.Package = *protoFile.Package
						method.Service = *service.Name
						method.Method = md

						if aut, err := proto.GetExtension(md.Options, runtime.E_Authentication); err == nil {
							au := aut.(*bool)
							options[runtime.E_Authentication.Name] = au
						}
						method.Options = options

						if ext, err := proto.GetExtension(md.Options, google_api.E_Http); err == nil {
							pattern := types.Pattern{}
							rule := ext.(*google_api.HttpRule)
							verb, path := getVerbAndPath(rule)
							pattern.Verb = verb
							pattern.Path = path
							pattern.Body = rule.Body
							method.Pattern = pattern
							//options[google_api.E_Http.Name] = rule
							methods = append(methods, method)
						}
					}
				}
			}
		}
	}

	jsonOut, err := json.Marshal(methods)
	if err != nil {
		log.Panicln("json.Marshal eror", err)
	}
	f, _ := os.Create("parse.json")
	str := strconv.Quote(string(jsonOut))
	f.WriteString(str)
}

func getVerbAndPath(opts *google_api.HttpRule) (string, string) {
	var httpMethod, path string
	switch {
	case opts.GetGet() != "":
		httpMethod = "GET"
		path = opts.GetGet()
	case opts.GetPost() != "":
		httpMethod = "POST"
		path = opts.GetPost()
	case opts.GetPut() != "":
		httpMethod = "PUT"
		path = opts.GetPut()
	case opts.GetDelete() != "":
		httpMethod = "DELETE"
		path = opts.GetDelete()
	case opts.GetPatch() != "":
		httpMethod = "PATCH"
		path = opts.GetPatch()
	case opts.GetCustom() != nil:
		custom := opts.GetCustom()
		httpMethod = custom.Kind
		path = custom.Path
	}
	return httpMethod, path
}
