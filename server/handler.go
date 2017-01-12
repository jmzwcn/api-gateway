package server

import (
	"context"
	"io/ioutil"
	"log"
	"lovev/api"
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
)

func handleForward(ctx context.Context, req *http.Request, opts ...grpc.CallOption) (*string, error) {
	out := new(lovev.ProfileModel)
	//out := new(proto.Message)
	rpcConn, err := grpc.Dial("127.0.0.1:9090", []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		return nil, err
	}

	var protoReq lovev.ProfileRequest

	body, _ := ioutil.ReadAll(req.Body)
	log.Println(string(body))
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	err = unmarshaler.Unmarshal(strings.NewReader(string(body)), &protoReq)
	if err != nil {
		log.Println(err)
	}
	log.Println(protoReq)

	err = grpc.Invoke(ctx, "/lovev.Profile/Get", &protoReq, out, rpcConn, opts...)
	if err != nil {
		return nil, err
	}
	str := (*out).String()
	return &str, nil
}
