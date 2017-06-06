# Variables
GOCMD=go
PROTOC =protoc

ifeq (${shell uname -s}, Darwin)
	SED=gsed
else
	SED=sed
endif

build: prepare copy	parse initial post-pb
	$(GOCMD) build cmd/api-gateway.go

prepare: SHELL:=bash
prepare:
	@echo "Downloading dependency..."
	@$(GOCMD) get google.golang.org/grpc
#	@$(GOCMD) get github.com/golang/protobuf/protoc-gen-go
	@$(GOCMD) get -u github.com/gogo/protobuf/protoc-gen-go{fast,gofast,gofaster,goslick}

copy:
	$(GOCMD) run plugin/proto.go -copy

parse:
	$(GOCMD) build github.com/api-gateway/plugin/protoc-gen-parse
	$(PROTOC) -I. -Ithird_party -I../../github.com/gogo/protobuf/ service/*/*.proto --plugin=protoc-gen-parse --parse_out=.

initial:
	$(GOCMD) run plugin/proto.go -initial

post-pb:
#	@$(SED) -i '/google\/api/d' service/*/*.pb.go
	@rm parse.json protoc-gen-parse;

devel:
	build
	
test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
