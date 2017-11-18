package server

import (
	"context"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/api-gateway/common"
	"github.com/api-gateway/loader"
	"github.com/api-gateway/types"

	goproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	rpcHost = "localhost" // e.g:grpc_service_name
	rpcPort = "50051"
)

func handleForward(ctx context.Context, req *http.Request, opts ...grpc.CallOption) (string, error) {
	log.Debug("Header", req.Header)
	body, _ := ioutil.ReadAll(req.Body)
	log.Debug("Body", string(body))

	sm, err := searchMethod(req.Method, req.URL.Path)
	if err != nil {
		return "", err
	}

	req.ParseForm()
	jsonContent := mergeToBody(string(body), sm.MergeValue, req)
	log.Debug("jsonContent:", jsonContent)

	protoReq := protoMessage(sm.Method.GetInputType())
	protoRes := protoMessage(sm.Method.GetOutputType())
	if err = jsonpb.UnmarshalString(jsonContent, protoReq); err != nil {
		log.Error(err)
	}
	//sm.package represents for module name by default, meaning service name
	rpcServer := sm.Package + ":" + rpcPort
	rpcConn, err := grpc.Dial(rpcServer, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	requestURL := "/" + sm.Package + "." + sm.Service + "/" + *sm.Method.Name
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
	return nil, status.Error(codes.NotFound, key+" not been found.")
}

func mergeToBody(bodyValue, pathValue string, req *http.Request) string {
	if bodyValue == "" {
		return ""
	}
	queryValue := ""
	for k, v := range req.Form {
		queryValue = queryValue + ",\"" + k + "\":\"" + v[0] + "\"" + ",\"" + k + "\":" + v[0]
	}
	return strings.TrimSuffix(bodyValue, "}") + pathValue + queryValue + "}"

}

func protoMessage(messageTypeName string) proto.Message {
	typeName := strings.TrimLeft(messageTypeName, ".")
	messageType := goproto.MessageType(typeName)
	if messageType == nil {
		messageType = proto.MessageType(typeName)
	}
	return reflect.New(messageType.Elem()).Interface().(proto.Message)
}
