package server

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
)

func WSHandler(req *http.Request, ws *websocket.Conn) {
	sm, err := searchMethod(req.Method, req.URL.Path)
	if err != nil {
		log.Panicln(err)
	}
	in := protoMessage(sm.Method.GetInputType())
	out := protoMessage(sm.Method.GetOutputType())

	service := sm.Package + ":" + rpcPort
	conn, err := grpc.Dial(service, grpc.WithInsecure(), grpc.WithBlock())
	fullMethod := "/" + sm.Package + "." + sm.Service + "/" + *sm.Method.Name

	streamDesc := &grpc.StreamDesc{
		StreamName:    *sm.Method.Name,
		ClientStreams: sm.Method.GetClientStreaming(),
		ServerStreams: sm.Method.GetServerStreaming(),
	}
	stream, err := grpc.NewClientStream(context.Background(), streamDesc, conn, fullMethod)
	if err != nil {
		log.Panicln(err)
	}
	if streamDesc.ClientStreams {
		go func() {
			for {
				var jsonStr string
				err := websocket.Message.Receive(ws, &jsonStr)
				if err == io.EOF {
					ws.Close()
					break
				}
				if err = jsonpb.UnmarshalString(jsonStr, in); err != nil {
					log.Panicln(err)
					//return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
				}
				stream.SendMsg(&in)
			}
		}()
	}

	if streamDesc.ServerStreams {
		go func() {
			for {
				err := stream.RecvMsg(&out)
				if err == io.EOF {
					ws.Close()
					break
				}
				json, err := (&jsonpb.Marshaler{}).MarshalToString(out)
				if err != nil {
					log.Panicln(err)
				}
				websocket.Message.Send(ws, &json)
			}
		}()
	}
}
