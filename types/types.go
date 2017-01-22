package types

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type MethodWrapper struct {
	Package string
	Service string
	Method  *descriptor.MethodDescriptorProto
	Pattern *Pattern
}

type Pattern struct {
	Verb string
	Path string
	Body string
}

type RuleStore struct {
	Store map[string]MethodWrapper
}

func CreateRuleStore() *RuleStore {
	store := make(map[string]MethodWrapper)
	return &RuleStore{Store: store}
}

func (rs *RuleStore) Compile(key string) MethodWrapper {
	return rs.Store[key] //TODO enhance with regular express
}
