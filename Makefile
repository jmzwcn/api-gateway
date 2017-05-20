# Variables
GOCMD=go
PROTOC = "protoc/protoc-3.1.0-linux-x86_64/bin"

ifeq (${shell uname -s}, Darwin)
	SED=gsed
else
	SED=sed
endif

build: prepare	parse generate post-pb
	$(GOCMD) build cmd/api-gateway.go

prepare: SHELL:=bash
prepare:
	@echo "Downloading dependency..."
	@$(GOCMD) get google.golang.org/grpc
	@$(GOCMD) get github.com/golang/protobuf/protoc-gen-go
#	@$(GOCMD) get -u github.com/gogo/protobuf/protoc-gen-go{fast,gofast,gofaster,goslick}

parse:
	$(GOCMD) build github.com/api-gateway/plugin/protoc-gen-parse

generate:
#	@echo ${PATH}
	$(GOCMD) run plugin/generate.go --protoc=$(PROTOC)

post-pb:
	@$(SED) -i '/google\/api/d' service/*/*.pb.go
	@rm parse.json protoc-gen-parse;

devel:
	build
	
test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
