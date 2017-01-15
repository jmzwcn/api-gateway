# Variables
GOCMD=go
PROTOC=protoc-3.1.0-linux-x86_64/bin/protoc

build:	parse generate init-pb
	$(GOCMD) build cmd/api-gateway.go

parse:
	$(GOCMD) build api-gateway/plugin/protoc-gen-parse
	$(PROTOC) -I. -Ithird_party --plugin=protoc-gen-parse --parse_out=. service/*.proto

generate:
	$(PROTOC) -Ithird_party -I.  --go_out=plugins=grpc:.  service/*.proto
	sed -i '/google\/api/d' service/*.pb.go

LIST = `ls service/*.proto`
init-pb:
	echo "package loader"> loader/initial.go;\
	echo 'import profile "api-gateway/service"'>> loader/initial.go;\
	echo "func initPB() {}">> loader/initial.go;\
	for filename in $(LIST); do \
	sn=`grep -n service $$filename | cut -d " " -f2`;\
#	sed -i  '/package loader/a \import '$$sn' "api-gateway/service"' loader/initial.go;\
	echo  'var _ = profile.New'$$sn'Client' >> loader/initial.go;\
	done;\
	json=`cat parse.json`;\
	echo  "const PROTO_JSON = "$$json >> loader/initial.go;
devel:
	build
	
test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
