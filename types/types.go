package types

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type MethodWrapper struct {
	Package string
	Service string
	Method  *descriptor.MethodDescriptorProto
	Pattern Pattern
}

type Pattern struct {
	Verb string
	Path string
	Body string
}
