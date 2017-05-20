package server

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/api-gateway/common"
	"github.com/api-gateway/loader"
	"github.com/api-gateway/types"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

const (
	RPC_SERVER = "localhost:50051" // e.g:grpc_service_name:50051
)

func handleForward(ctx context.Context, req *http.Request, opts ...grpc.CallOption) (string, error) {
	log.Debug("Header", req.Header)
	body, _ := ioutil.ReadAll(req.Body)
	log.Debug("Body", string(body))

	method, err := searchMethod(req.Method, req.URL.Path)
	if err != nil {
		return "", err
	}

	req.ParseForm()
	jsonContent := mergeToBody(string(body), method.MergeValue, req)
	log.Debug("jsonContent:" + jsonContent)

	protoReq := protoMessage(method.Method.GetInputType())
	protoRes := protoMessage(method.Method.GetOutputType())
	if err = jsonpb.UnmarshalString(jsonContent, protoReq); err != nil {
		log.Error(err)
	}

	rpcConn, err := grpc.Dial(RPC_SERVER, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	requestURL := "/" + method.Package + "." + method.Service + "/" + *method.Method.Name
	log.Debug(requestURL)

	if err = grpc.Invoke(ctx, requestURL, protoReq, protoRes, rpcConn, opts...); err != nil {
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

func protoMessage(messageTypeName string) proto.Message {
	messageType := proto.MessageType(strings.TrimLeft(messageTypeName, ".")).Elem()
	return reflect.New(messageType).Interface().(proto.Message)
}
