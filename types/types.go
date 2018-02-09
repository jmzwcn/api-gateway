package types

import (
	"net/url"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

type MethodWrapper struct {
	Package string
	Service string
	Method  *descriptor.MethodDescriptorProto
	Pattern Pattern
	Options map[string]interface{}
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
