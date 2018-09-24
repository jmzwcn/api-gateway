package types

import (
	"encoding/json"
	"net/url"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"google.golang.org/grpc/encoding"
)

type MethodWrapper struct {
	Package             string
	Service             string
	Method              *descriptor.MethodDescriptorProto
	InputTypeDescriptor *descriptor.DescriptorProto
	Pattern             Pattern
	Options             map[string]interface{}
}

type Pattern struct {
	Verb string
	Path string
	Body string
}

type MatchedMethod struct {
	MethodWrapper
	Precision  int
	PathValues url.Values
}

func init() {
	encoding.RegisterCodec(MuxCodec{})
}

type MuxCodec struct{}

func (MuxCodec) Marshal(v interface{}) ([]byte, error) {
	// if v, ok := v.(proto.Message); ok {
	// 	vs, _ := (&jsonpb.Marshaler{}).MarshalToString(v)
	// 	return []byte(vs), nil
	// }
	return json.Marshal(v)
}

func (MuxCodec) Unmarshal(data []byte, v interface{}) error {
	// if v, ok := v.(proto.Message); ok {
	// 	return jsonpb.Unmarshal(bytes.NewBuffer(data), v)
	// }
	return json.Unmarshal(data, v)
}

func (MuxCodec) Name() string {
	return "mux"
}
