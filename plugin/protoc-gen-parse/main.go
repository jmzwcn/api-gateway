package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gogo/protobuf/proto"
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
	pfs := request.GetProtoFile()
	service := pfs[0].Service[0]

	jsonOut, err := json.Marshal(service)
	if err != nil {
		log.Println("json.Marshal eror", err)
	}
	//fmt.Print(jsonOut)
	os.Stdout.Write(jsonOut)
	//fmt.Print(service.GetMethod())
	//fmt.Println(serice.Options)
}
