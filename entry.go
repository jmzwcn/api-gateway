package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func main() {
	if _, err := exec.Command("go", "build", "api-gateway/plugin/protoc-gen-parse").CombinedOutput(); err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("protoc/protoc-3.1.0-linux-x86_64/bin/protoc", "--plugin=protoc-gen-parse", "--parse_out=.", "service/helloworld.proto")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	var pfs descriptor.ServiceDescriptorProto
	err = json.Unmarshal(out, &pfs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pfs.Name)
}

func main1() {
	log.Println("begin")

	data, err := ioutil.ReadFile("proto/helloworld.proto")
	if err != nil {
		log.Fatal("read file error: ", err)
	}
	//log.Println(string(data))

	var pb descriptor.FileDescriptorProto
	if err := proto.Unmarshal(data, &pb); err != nil {
		log.Fatal("unmarshaling error: ", err)
	} else {
		log.Println(pb)
	}
}
