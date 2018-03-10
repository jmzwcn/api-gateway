##################################################
# All macro-services's parent directory
SERVICES_PARENT_DIR=github.com/api-gateway/example
# PROTO_DIR includes *.proto
PROTO_DIR=service
##################################################
# Variables
SERVICE=api-gateway
IMG_HUB?=registry.test.io/test
TAG?=latest
# Version information
VERSION=1.0.0
REVISION=${shell git rev-parse --short HEAD}
RELEASE=production
BUILD_HASH=${shell git rev-parse HEAD}
BUILD_TIME=${shell date "+%Y-%m-%d@%H:%M:%SZ%z"}
LD_FLAGS:=-X main.Version=$(VERSION) -X main.Revision=$(REVISION) -X main.Release=$(RELEASE) -X main.BuildHash=$(BUILD_HASH) -X main.BuildTime=$(BUILD_TIME)

ifeq (${shell uname -s}, Darwin)
	SED=gsed
else
	SED=sed
endif

prepare: SHELL:=bash
prepare:download
	@-docker swarm init
	@-docker network create --driver=overlay devel
	@go install github.com/api-gateway/plugin/...

download:
	@echo "Download dependencies..."
	@go get google.golang.org/grpc
	@go get github.com/gogo/protobuf/protoc-gen-gogofast

parse:	
	@protoc -I${GOPATH}/src \
	-I${GOPATH}/src/github.com/api-gateway/third_party \
	-I${GOPATH}/src/github.com/gogo/protobuf ${GOPATH}/src/$(SERVICES_PARENT_DIR)/*/$(PROTO_DIR)/*.proto --parse_out=.
	@echo Generate successfully.

initial:
	@echo "package router\n"> router/initial.go;
	@for dir in $(shell cd ../../ && ls -d $(SERVICES_PARENT_DIR)/*/$(PROTO_DIR)); do \
	echo 'import _ "'$$dir'"'>> router/initial.go; done;\
	json=`cat rules.json`;\
	echo "\nconst PROTO_JSON = "$$json >> router/initial.go;
	@echo Initial successfully.

build:parse initial
	@go build -ldflags="$(LD_FLAGS)" -o bundles/$(SERVICE) cmd/main.go	
	docker build -t $(IMG_HUB)/$(SERVICE):$(TAG) .
	@rm rules.json

run:prepare
	cd example/echo && make run
	cd example/helloworld && make run
	@make build
	@-docker service rm $(SERVICE)	
	@docker service create --name $(SERVICE) --network devel -p 8080:8080 $(IMG_HUB)/$(SERVICE):$(TAG)

test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
