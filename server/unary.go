package server

import (
	"context"
	"net/http"
	"reflect"
	"strings"

	"github.com/api-gateway/loader"
	"github.com/api-gateway/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var (
	rpcPort = "50051"
)

func handleForward(ctx context.Context, req *http.Request, opts ...grpc.CallOption) (proto.Message, error) {
	sm, err := searchMethod(req.Method, req.URL.Path)
	if err != nil {
		return nil, err
	}

	in := protoMessage(sm.Method.GetInputType())
	out := protoMessage(sm.Method.GetOutputType())

	jsonContent, err := mergeBody(req, sm.PathValues, in)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}
	grpclog.Println("jsonContent:", jsonContent)

	if err = jsonpb.UnmarshalString(jsonContent, in); err != nil {
		grpclog.Error(err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	//sm.package represents for service name by default
	service := sm.Package + ":" + rpcPort
	conn, err := grpc.Dial(service, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}
	defer conn.Close()

	fullMethod := "/" + sm.Package + "." + sm.Service + "/" + *sm.Method.Name
	if err = grpc.Invoke(ctx, fullMethod, in, out, conn, opts...); err != nil {
		grpclog.Error(err)
		return nil, err
	}
	return out, nil
}

func searchMethod(method, path string) (*types.MatchedMethod, error) {
	key := method + ":" + path
	matchedMethod := loader.RuleStore.Match(key)
	if matchedMethod != nil {
		return matchedMethod, nil
	}
	return nil, status.Error(codes.NotFound, key+" not been found.")
}

func protoMessage(messageTypeName string) proto.Message {
	typeName := strings.TrimLeft(messageTypeName, ".")
	messageType := proto.MessageType(typeName)
	return reflect.New(messageType.Elem()).Interface().(proto.Message)
}
