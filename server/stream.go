package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/api-gateway/types"
	"github.com/api-gateway/types/log"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
)

func streamHandler(req *http.Request, ws *websocket.Conn) {
	sm, err := searchMethod(req.Method, req.URL.Path)
	if err != nil {
		log.Error(err)
	}
	in := json.RawMessage{}     //protoMessage(sm.Method.GetInputType())
	out := new(json.RawMessage) //protoMessage(sm.Method.GetOutputType())

	service := sm.Package + ":" + rpcPort
	conn, err := grpc.Dial(service, grpc.WithInsecure(), grpc.WithBlock())
	fullMethod := "/" + sm.Package + "." + sm.Service + "/" + *sm.Method.Name

	streamDesc := &grpc.StreamDesc{
		StreamName:    *sm.Method.Name,
		ClientStreams: sm.Method.GetClientStreaming(),
		ServerStreams: sm.Method.GetServerStreaming(),
	}
	stream, err := grpc.NewClientStream(context.Background(), streamDesc, conn, fullMethod, grpc.CallContentSubtype(types.MuxCodec{}.Name()))
	if err != nil {
		log.Error(err)
	}
	//write
	go func() {
		for {
			err := stream.RecvMsg(out)
			if err == io.EOF || err != nil {
				break
			}
			// json, err := (&jsonpb.Marshaler{}).MarshalToString(out)
			// if err != nil {
			// 	log.Error(err)
			// 	continue
			// }
			websocket.Message.Send(ws, out)
		}
	}()
	//read
	for {
		//var jsonStr string
		err := websocket.Message.Receive(ws, &in)
		if err == io.EOF {
			ws.Close()
			break
		}
		// if err = jsonpb.UnmarshalString(jsonStr, in); err != nil {
		// 	log.Error(err)
		// 	continue
		// }
		stream.SendMsg(in)
	}
}
