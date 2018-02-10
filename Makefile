# Variables
SERVICE=api-gateway
IMG_HUB?=registry.test.com/test
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
prepare:
	@-docker swarm init
	@-docker network create --driver=overlay devel

parse:
	go install github.com/api-gateway/plugin/...
	protoc -I${GOPATH}/src -I${GOPATH}/src/github.com/api-gateway/third_party -I${GOPATH}/src/github.com/gogo/protobuf ${GOPATH}/src/github.com/api-gateway/example/*/service/*.proto --parse_out=.

initial:
	@echo "package loader"> loader/initial.go;
	@for dir in $(shell cd ../../ && ls -d github.com/api-gateway/example/*/service); do \
	echo 'import _ "'$$dir'"'>> loader/initial.go;\
	done;\
	json=`cat parse.json`;\
	echo "const PROTO_JSON = "$$json >> loader/initial.go;
	@echo Initial successfully.

build:parse initial
	go build -ldflags="$(LD_FLAGS)" -o bundles/$(SERVICE) cmd/main.go
	@rm parse.json

image:build
	docker build -t $(IMG_HUB)/$(SERVICE):$(TAG) .

run:prepare image
	cd example/echo && make run
	cd example/helloworld && make run
	@-docker service rm $(SERVICE) > /dev/null 2>&1  || true	
	@docker service create --name $(SERVICE) --network devel -p 8080:8080 $(IMG_HUB)/$(SERVICE):$(TAG)

test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
