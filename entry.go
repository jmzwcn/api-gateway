package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func main() {
	if _, err := exec.Command("go", "build", "api-gateway/plugin/protoc-gen-parse").CombinedOutput(); err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("protoc/protoc-3.1.0-linux-x86_64/bin/protoc", "--plugin=protoc-gen-parse", "--parse_out=.", "service/helloworld.proto")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error:", err)
	}

	log.Println(string(out))
	nv := strings.Split(string(out), "unparseable:")[1]
	log.Println(nv)
	nnv := strings.Replace(nv, "\\", "", -1)
	log.Println(nnv)

	var pfs descriptor.ServiceDescriptorProto
	err = json.Unmarshal([]byte(nnv), &pfs)
	if err != nil {
		log.Println(err)
	}
	log.Println(*pfs.GetMethod()[0].InputType)
}
