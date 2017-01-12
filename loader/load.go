package loader

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"api-gateway/types"
)

var PatternStore = make(map[string]types.MethodWrapper)

func ParseAndLoad() map[string]types.MethodWrapper {
	load()
	return PatternStore
}

func load() {
	if _, err := exec.Command("go", "build", "api-gateway/plugin/protoc-gen-parse").CombinedOutput(); err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("protoc/protoc-3.1.0-linux-x86_64/bin/protoc",
		"-I.", "-Iservice/third_party", "--plugin=protoc-gen-parse", "--parse_out=.",
		"service/message.proto", "service/echo.proto", "service/helloworld.proto", "service/user.proto")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error:", err)
	}

	//log.Println(string(out))
	nv := strings.Split(string(out), "unparseable:")[1]
	//log.Println(nv)
	nnv := strings.Replace(nv, "\\", "", -1)
	log.Println(nnv)

	var methods []types.MethodWrapper
	err = json.Unmarshal([]byte(nnv), &methods)
	if err != nil {
		log.Println(err)
	}

	for _, md := range methods {
		if md.Package != nil {
			log.Println(*md.Package, *md.Pattern)
		} else {
			log.Println(*md.Pattern)
		}
		key := *md.Method + ":" + md.Pattern.Path
		PatternStore[key] = md
	}
}
