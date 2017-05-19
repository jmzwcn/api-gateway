package main

import (
	"api-gateway/common"
	"api-gateway/config"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	cmdOut []byte
	err    error
)

const (
	protoc = "protoc/protoc-3.1.0-linux-x86_64/bin"
)

func main() {
	os.Setenv("PATH", os.Getenv("PATH")+":.:"+protoc)
	configuration := config.NewConfiguration()
	protoFiles := make([]string, 0)
	initialContent := "package loader\n"

	for _, proto := range configuration.ProtoSet {
		log.Debug("Current proto:", proto.Service, proto.Path)
		os.Remove("service")
		//copy source to destinationDir
		destinationDir := "service/" + proto.Service
		os.MkdirAll(destinationDir, os.ModePerm)

		directory, _ := os.Open(proto.Path)
		objects, _ := directory.Readdir(-1)
		for _, obj := range objects {
			if !obj.IsDir() && strings.HasSuffix(obj.Name(), "proto") {
				protoFiles = append(protoFiles, destinationDir+"/"+obj.Name())
				CopyFile(directory.Name()+"/"+obj.Name(), destinationDir+"/"+obj.Name())

				args := append([]string{"-I.", "-Ithird_party", "--go_out=."}, destinationDir+"/"+obj.Name())
				if cmdOut, err = exec.Command("protoc", args...).CombinedOutput(); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				log.Debug(string(cmdOut))
			}
		}
		initialContent = initialContent + "import _ \"api-gateway/service/" + proto.Service + "\"\n"

	}

	args := append([]string{"-I.", "-Ithird_party", "--plugin=protoc-gen-parse", "--parse_out=."}, protoFiles...)
	if cmdOut, err = exec.Command("protoc", args...).CombinedOutput(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	sha := string(cmdOut)
	log.Debug(sha)

	bytess, err := ioutil.ReadFile("parse.json")
	if err != nil {
		log.Error(err)
	}

	initialContent = initialContent + "const PROTO_JSON = " + string(bytess)
	err = ioutil.WriteFile("loader/initial.go", []byte(initialContent), 0644)
	if err != nil {
		log.Error(err)
	}
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}
