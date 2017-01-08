package main

import (
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func main() {
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
