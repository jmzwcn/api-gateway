# Variables
GOCMD=go

ifeq (${shell uname -s}, Darwin)
	SED=gsed
else
	SED=sed
endif

build: init	parse generate post-pb
	$(GOCMD) build cmd/api-gateway.go

init: SHELL:=bash
init:
	@$(GOCMD) get google.golang.org/grpc
	@$(GOCMD) get github.com/golang/protobuf/protoc-gen-go
#	@$(GOCMD) get -u github.com/gogo/protobuf/protoc-gen-go{fast,gofast,gofaster,goslick}

parse:
	$(GOCMD) build api-gateway/plugin/protoc-gen-parse

generate:
	$(GOCMD) run plugin/generate.go

post-pb:
	@$(SED) -i '/google\/api/d' service/*/*.pb.go
	@rm parse.json protoc-gen-parse;

devel:
	build
	
test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
