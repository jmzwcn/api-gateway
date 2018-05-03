package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/api-gateway/types"
	"github.com/api-gateway/types/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	rpcPort = "50051"
)

func handleForward(ctx context.Context, req *http.Request, opts ...grpc.CallOption) ([]byte, error) {
	sm, err := searchMethod(req.Method, req.URL.Path)
	if err != nil {
		return nil, err
	}

	mergedJSON, err := mergeBody(req, sm)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}
	log.Infoln("mergedJSON:", mergedJSON)

	input := json.RawMessage([]byte(mergedJSON))
	output := json.RawMessage{}

	//sm.package represents for service name by default
	service := sm.Package + ":" + rpcPort
	conn, err := grpc.Dial(service, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer conn.Close()

	fullMethod := "/" + sm.Package + "." + sm.Service + "/" + *sm.Method.Name
	if err = grpc.Invoke(ctx, fullMethod, input, &output, conn, grpc.CallContentSubtype(types.MuxCodec{}.Name())); err != nil {
		log.Error(err)
		return nil, err
	}
	return output, nil
}

func searchMethod(method, path string) (*types.MatchedMethod, error) {
	key := method + ":" + path
	matchedMethod := ruleStore.Match(key)
	if matchedMethod != nil {
		return matchedMethod, nil
	}
	return nil, status.Error(codes.NotFound, key+" not been found.")
}
