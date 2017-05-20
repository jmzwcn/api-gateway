package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/api-gateway/common"
	"github.com/api-gateway/config"
)

var (
	cmdOut []byte
	err    error
)

func main() {
	initProtoc()
	configuration := config.NewConfiguration()
	protoFiles := make([]string, 0)
	initialContent := "package loader\n"

	for _, proto := range configuration.ProtoSet {
		os.Remove("service")
		destinationDir := "service/" + proto.Service
		os.MkdirAll(destinationDir, os.ModePerm)

		directory, err := os.Open(proto.Path)
		if err != nil {
			log.Error(err)
		}
		objects, _ := directory.Readdir(-1)
		for _, obj := range objects {
			if !obj.IsDir() && strings.HasSuffix(obj.Name(), "proto") {
				protoFiles = append(protoFiles, destinationDir+"/"+obj.Name())
				copyFile(directory.Name()+"/"+obj.Name(), destinationDir+"/"+obj.Name())

				args := append([]string{"-I.", "-Ithird_party", "--go_out=."}, destinationDir+"/"+obj.Name())
				if cmdOut, err = exec.Command("protoc", args...).CombinedOutput(); err != nil {
					log.Error(err, string(cmdOut))
				}
			}
		}
		initialContent = initialContent + "import _ \"github.com/api-gateway/service/" + proto.Service + "\"\n"
	}

	args := append([]string{"-I.", "-Ithird_party", "--plugin=protoc-gen-parse", "--parse_out=."}, protoFiles...)
	if cmdOut, err = exec.Command("protoc", args...).CombinedOutput(); err != nil {
		log.Error(err, string(cmdOut))
	}

	bytes, err := ioutil.ReadFile("parse.json")
	if err != nil {
		log.Error(err)
	}

	initialContent = initialContent + "const PROTO_JSON = " + string(bytes)
	err = ioutil.WriteFile("loader/initial.go", []byte(initialContent), os.ModePerm)
	if err != nil {
		log.Error(err)
	}
}

func initProtoc() {
	var protoc string
	flag.StringVar(&protoc, "protoc", "", "protoc")
	flag.Parse()
	os.Setenv("PATH", os.Getenv("PATH")+":../../../bin:.:"+protoc)
}

func copyFile(source string, dest string) (err error) {
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
			os.Chmod(dest, sourceinfo.Mode())
		}
	}
	return
}
