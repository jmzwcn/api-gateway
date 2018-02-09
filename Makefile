# Variables
GOCMD=go
PROTOC =protoc

ifeq (${shell uname -s}, Darwin)
	SED=gsed
else
	SED=sed
endif

build: prepare parse initial post-pb
	$(GOCMD) build cmd/api-gateway.go

prepare: SHELL:=bash
prepare:
	@echo "Downloading dependency..."
	@$(GOCMD) get google.golang.org/grpc
#	@$(GOCMD) get github.com/golang/protobuf/protoc-gen-go
	@$(GOCMD) get -u github.com/gogo/protobuf/protoc-gen-go{fast,gofast,gofaster,goslick}

parse:
	$(GOCMD) install github.com/api-gateway/plugins/...
	$(PROTOC) -I${GOPATH}/src -I${GOPATH}/src/github.com/api-gateway/third_party -I${GOPATH}/src/github.com/gogo/protobuf ${GOPATH}/src/github.com/*/service/*.proto --parse_out=.

initial:
	@echo "package loader"> loader/initial.go;
	@for dir in $(shell cd ../../ && ls -d github.com/*/service); do \
	echo 'import _ "'$$dir'"'>> loader/initial.go;\
	done;\
	json=`cat parse.json`;\
	echo "const PROTO_JSON = "$$json >> loader/initial.go;
	@echo Initial successfully.

post-pb:
#	@$(SED) -i '/google\/api/d' service/*/*.pb.go
	@rm parse.json

devel:
	build
	
test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
