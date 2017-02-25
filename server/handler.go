package server

import (
	"api-gateway/common"
	"api-gateway/loader"
	"api-gateway/types"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

const (
	RPC_SERVER = "localhost:9090" // e.g:service:9090
)

func handleForward(ctx context.Context, req *http.Request, opts ...grpc.CallOption) (string, error) {
	method, err := searchMethod(req.Method, req.URL.Path)
	if err != nil {
		return "", err
	}
	inputType := method.Method.GetInputType()
	typeName := strings.TrimLeft(inputType, ".")
	//log.Debug(proto.MessageType(typeName))
	outputType := method.Method.GetOutputType()
	outTtypeName := strings.TrimLeft(outputType, ".")
	protoRes := reflect.New(proto.MessageType(outTtypeName).Elem()).Interface().(proto.Message)
	//out := new(any.Any)
	rpcConn, err := grpc.Dial(RPC_SERVER, []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		return "", err
	}

	protoReq := reflect.New(proto.MessageType(typeName).Elem()).Interface().(proto.Message)

	body, _ := ioutil.ReadAll(req.Body)
	log.Debug(string(body))
	req.ParseForm()
	jsonContent := mergeToBody(string(body), method.MergeValue, req)
	log.Debug("jsonContent:" + jsonContent)
	err = jsonpb.UnmarshalString(jsonContent, protoReq)
	if err != nil {
		log.Error(err)
	}
	rpcURL := "/" + method.Package + "." + method.Service + "/" + *method.Method.Name
	log.Debug(rpcURL)
	err = grpc.Invoke(ctx, rpcURL, protoReq, protoRes, rpcConn, opts...)
	if err != nil {
		return "", err
	}
	return new(jsonpb.Marshaler).MarshalToString(protoRes)
}

func searchMethod(method, path string) (*types.MatchedMethod, error) {
	key := method + ":" + path
	log.Debug("key", key)
	matchedMethod := loader.RuleStore.Match(key)
	if matchedMethod != nil {
		return matchedMethod, nil
	}
	return nil, errors.New(key + " not been found.")
}

func mergeToBody(bodyValue, pathValue string, req *http.Request) string {
	queryValue := ""
	for k, v := range req.Form {
		queryValue = queryValue + ",\"" + k + "\":\"" + v[0] + "\""
	}
	return strings.TrimSuffix(bodyValue, "}") + pathValue + queryValue + "}"
}
