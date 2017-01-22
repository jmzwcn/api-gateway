package server

import (
	"api-gateway/loader"
	"api-gateway/types"
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

func handleForward(ctx context.Context, req *http.Request, opts ...grpc.CallOption) (string, error) {
	method := getPBMethod(req.Method, req.RequestURI)
	if method == nil || method.Method == nil {
		errStr := "URL:" + req.RequestURI + " not found."
		return errStr, errors.New(errStr)
	}
	inputType := method.Method.GetInputType()
	typeName := strings.TrimLeft(inputType, ".")
	//log.Println(proto.MessageType(typeName))
	outputType := method.Method.GetOutputType()
	outTtypeName := strings.TrimLeft(outputType, ".")
	protoRes := reflect.New(proto.MessageType(outTtypeName).Elem()).Interface().(proto.Message)
	//out := new(any.Any)
	rpcConn, err := grpc.Dial("127.0.0.1:9090", []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		return "", err
	}

	protoReq := reflect.New(proto.MessageType(typeName).Elem()).Interface().(proto.Message)

	body, _ := ioutil.ReadAll(req.Body)
	log.Println(string(body))
	err = jsonpb.UnmarshalString(string(body), protoReq)
	if err != nil {
		log.Println(err)
	}
	//log.Println(protoReq)
	rpcURL := "/" + method.Package + "." + method.Service + "/" + *method.Method.Name
	log.Println(rpcURL)
	err = grpc.Invoke(ctx, rpcURL, protoReq, protoRes, rpcConn, opts...)
	if err != nil {
		return "", err
	}
	return new(jsonpb.Marshaler).MarshalToString(protoRes)
}

func getPBMethod(method, path string) *types.MethodWrapper {
	key := method + ":" + path
	log.Println("key", key)
	methodWrapper := loader.RuleStore.Compile(key)
	if &methodWrapper != nil {
		return &methodWrapper
	}
	log.Println(key + " not been found.")
	return nil
}
