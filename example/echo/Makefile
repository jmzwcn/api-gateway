# Variables
SERVICE=echo
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

prepare:
	go install github.com/api-gateway/plugin/...

generate-go:
	@protoc -I${GOPATH}/src \
	-I${GOPATH}/src/github.com/api-gateway/third_party \
	-I${GOPATH}/src/github.com/gogo/protobuf \
	--gogofast_out=\
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
	plugins=grpc:../../../../ \
	 $(shell pwd)/service/*.proto
	@$(SED) -i '/google\/api/d' service/*.pb.go
	@echo Generate successfully.

parse:prepare generate-go	
	@protoc -I${GOPATH}/src -I${GOPATH}/src/github.com/api-gateway/third_party \
	$(shell pwd)/service/*.proto --parse_out=.
	@echo Parse successfully.

build:parse
	go build -ldflags="$(LD_FLAGS)" -o bundles/$(SERVICE) internal/cmd/main.go

image:build
	docker build -t $(IMG_HUB)/$(SERVICE):$(TAG) .

run:image
	@-docker service rm $(SERVICE) > /dev/null 2>&1  || true	
	@docker service create --name $(SERVICE) --network devel $(IMG_HUB)/$(SERVICE):$(TAG)

test:
	go test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
