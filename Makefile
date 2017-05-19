# Variables
GOCMD=go
PROTOC=protoc/protoc-3.1.0-linux-x86_64/bin/protoc

build:	parse prepare post-pb
	$(GOCMD) build cmd/api-gateway.go

parse:
	$(GOCMD) build api-gateway/plugin/protoc-gen-parse

prepare:
	$(GOCMD) run prepare.go

post-pb:
#	echo 'import _ "api-gateway/service"'>> loader/initial.go;
#	@rm parse.json protoc-gen-parse;

devel:
	build
	
test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
