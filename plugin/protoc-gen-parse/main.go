package main

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/api-gateway/third_party/google/api"
	"github.com/api-gateway/third_party/runtime"
	"github.com/api-gateway/types"
	"github.com/api-gateway/types/log"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
)

var (
	oneProtoFile *plugin_go.CodeGeneratorResponse_File
	messagesMap  = make(map[string]*descriptor.DescriptorProto)
)

func main() {
	//gogofast with extension
	req := command.Read()
	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)

	vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
	vanity.ForEachFile(files, vanity.TurnOnSizerAll)
	vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)

	resp := command.Generate(req)

	var methods []types.MethodWrapper
	for _, protoFile := range files {
		for _, generateFileName := range req.FileToGenerate {
			if *protoFile.Name == generateFileName {
				for _, v := range protoFile.MessageType {
					key := "." + *protoFile.Package + "." + *v.Name
					messagesMap[key] = v
				}

				for _, service := range protoFile.Service {
					for _, md := range service.Method {
						method := types.MethodWrapper{}
						options := make(map[string]interface{})

						method.Package = *protoFile.Package
						method.Service = *service.Name
						method.Method = md
						method.InputTypeDescriptor = messagesMap[*md.InputType]

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
		log.Fatalln("json.Marshal eror", err)
	}
	//inject api in one *.pb.go
	oneProtoFile = resp.File[len(resp.File)-1]
	lines := strings.Split(*oneProtoFile.Content, "\n")
	for i, line := range lines {
		if strings.Contains(line, `import fmt "fmt"`) {
			lines[i] = lines[i] + "\n" +
				`import http "net/http"` + "\n" +
				`import strings "strings"`
		}
	}
	output := strings.Join(lines, "\n")
	append := "\nconst PROTO_JSON =" + strconv.Quote(string(jsonOut)) + "\n" + `
func init() {
	if _, err := (&http.Client{}).Post("http://api-gateway:8080/rules", "", strings.NewReader(PROTO_JSON)); err != nil {
		panic(err)
	}
}`
	newContent := output + append
	oneProtoFile.Content = &newContent
	command.Write(resp)
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
