package server

/*
import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"nevis.io/fx/common/log"
	"nevis.io/fx/common/rpc"
)

//TODO
func handleForward1(ctx context.Context, w http.ResponseWriter, req *http.Request, opts ...grpc.CallOption) (proto.Message, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Debug("raw body:", string(body))
		return nil, err
	}

	sm, err := searchMethod(req.Method, req.URL.Path)
	if err != nil {
		return nil, err
	}

	json := mergeToBody(string(body), sm.MergeValues, req)

	pbReq := protoMessage(sm.Method.GetInputType())
	pbRes := protoMessage(sm.Method.GetOutputType())
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(strings.NewReader(json), pbReq); err != nil {
		log.Error(err)
	}

	conn, err := websocket.Upgrade(w, req, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	//sm.package represent for service name by default
	rpcConn := rpc.Dial(sm.Package)
	fullMethod := "/" + sm.Package + "." + sm.Service + "/" + *sm.Method.Name

	streamDesc := &grpc.StreamDesc{
		StreamName:    *sm.Method.Name,
		ClientStreams: sm.Method.GetClientStreaming(),
		ServerStreams: sm.Method.GetServerStreaming(),
	}
	stream, err := grpc.NewClientStream(ctx, streamDesc, rpcConn, fullMethod, opts...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	for {
		if streamDesc.ClientStreams {
			_, payload, err := conn.ReadMessage()
			if err != nil {
				log.Error(err)
				break
			}
			log.Debug("[read] read payload:", string(payload))
			if err = unmarshaler.Unmarshal(bytes.NewBuffer(payload), pbReq); err != nil {
				log.Error(err)
				break
			}
			stream.SendMsg(pbReq)
		}

		if streamDesc.ServerStreams {
			stream.RecvMsg(pbRes)
			marshaler := jsonpb.Marshaler{EmitDefaults: true}
			json, err := marshaler.MarshalToString(pbRes)
			if err != nil {
				log.Error(err)
				break
			}
			conn.WriteMessage(websocket.TextMessage, []byte(json))
		}
	}

	return pbRes, nil
}
*/
