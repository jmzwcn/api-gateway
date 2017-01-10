package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"api-gateway/types"
)

var PatternStore = make(map[string]types.MethodWrapper)

func ParseAndLoad() {
	load()
	//PatternStore["abc"] = urlMapping{Method: "PUT", URLPattern: "abc"}
}

func load() {
	if _, err := exec.Command("go", "build", "api-gateway/plugin/protoc-gen-parse").CombinedOutput(); err != nil {
		log.Fatalln(err)
	}

	//cmd := exec.Command("protoc/protoc-3.1.0-linux-x86_64/bin/protoc", "--plugin=protoc-gen-parse", "--parse_out=.", "service/helloworld.proto")
	cmd := exec.Command("protoc/protoc-3.1.0-linux-x86_64/bin/protoc",
		"-I.", "-I/home/jmzwcn/work/src", "-Iservice/third_party/protobuf",
		"--plugin=protoc-gen-parse", "--parse_out=.", "service/message.proto")
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
		log.Println(md.Pattern)
	}

	//md := pf.GetService()[0].GetMethod()[0]
	//ext, _ := proto.GetExtension(md.Options, options.E_Http)
	//log.Println(ext)
}
