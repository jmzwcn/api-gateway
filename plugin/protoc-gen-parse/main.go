package main

import (
	options "api-gateway/service/third_party/protobuf/google/api"
	"api-gateway/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("error reading input")
	}

	request := new(plugin.CodeGeneratorRequest)
	if err := proto.Unmarshal(data, request); err != nil {
		fmt.Println(err)
	}

	//l := len(request.GetProtoFile())
	var epf *descriptor.FileDescriptorProto
	name := request.FileToGenerate[0]
	for _, pf := range request.GetProtoFile() {
		if pf.GetName() == name {
			epf = pf
		}
	}
	var methods []types.MethodWrapper
	var exts []interface{}
	for _, md := range epf.Service[0].GetMethod() {
		ext, err := proto.GetExtension(md.Options, options.E_Http)
		if err == nil {
			exts = append(exts, ext)
			method := types.MethodWrapper{}
			method.Package = epf.Package
			method.Service = epf.Service[0].Name
			method.Method = md.Name
			pattern := types.Pattern{}
			rule := ext.(*options.HttpRule)
			pattern.Verb = getVerb(rule)
			pattern.Path = rule.GetGet() + rule.GetPost() + rule.GetPut() + rule.GetDelete()
			pattern.Body = rule.Body
			method.Pattern = &pattern
			methods = append(methods, method)
		}
	}

	jsonOut, err := json.Marshal(methods)
	if err != nil {
		log.Println("json.Marshal eror", err)
	}
	//fmt.Print(jsonOut)
	os.Stdout.Write(jsonOut)
}

func getVerb(opts *options.HttpRule) string {
	var httpMethod string
	switch {
	case opts.GetGet() != "":
		httpMethod = "GET"
		//pathTemplate = opts.GetPut()
	case opts.GetPost() != "":
		httpMethod = "POST"
		//pathTemplate = opts.GetPut()
	case opts.GetPut() != "":
		httpMethod = "PUT"
	//pathTemplate = opts.GetPut()
	case opts.GetDelete() != "":
		httpMethod = "DELETE"
	//pathTemplate = opts.GetPut()
	case opts.GetPatch() != "":
		httpMethod = "PATCH"
		//pathTemplate = opts.GetPut()
	}
	return httpMethod
}
