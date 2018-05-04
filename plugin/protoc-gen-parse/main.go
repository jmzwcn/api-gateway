package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/api-gateway/third_party/google/api"
	"github.com/api-gateway/third_party/runtime"
	"github.com/api-gateway/types"
	"github.com/api-gateway/types/log"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

var (
	oneProtoFile = ""
	messagesMap  = make(map[string]*descriptor.DescriptorProto)
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("error reading input")
	}

	request := new(plugin_go.CodeGeneratorRequest)
	if err := proto.Unmarshal(data, request); err != nil {
		log.Fatalln(err)
	}

	var methods []types.MethodWrapper
	for _, protoFile := range request.GetProtoFile() {
		for _, generateFileName := range request.FileToGenerate {
			oneProtoFile = generateFileName
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

	oneProtoFile = os.Getenv("GOPATH") + "/src/" + strings.TrimSuffix(oneProtoFile, ".proto") + ".pb.go"
	input, err := ioutil.ReadFile(oneProtoFile)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
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
	err = ioutil.WriteFile(oneProtoFile, []byte(output+append), 0644)
	if err != nil {
		log.Fatalln(err)
	}
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
