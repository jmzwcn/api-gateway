# Variables
GOCMD=go
PROTOC=protoc/protoc-3.1.0-linux-x86_64/bin/protoc

build:	source parse generate init-pb
	$(GOCMD) build cmd/api-gateway.go

source:
	echo "copy source...";\
	protoset=`cat config.json|awk -F"[proto.set]" '/path|service/{print$0}'`;\
	aa=`echo $$protoset | sed s/[[:space:]]//g`;\
	bb=$${aa//\"\"/ };\
	for i in $${aa//\"\"/ }; do \
	j=$${i//\"/};\
	k=$${j//service:/};\
	l=$${k//path:/};\
	m=$${l//,/ };\
	n=($$m);\
	rm -rf service/$${n[0]};\
	mkdir service/$${n[0]};\
	cp $${n[1]}/*.proto service/$${n[0]};\
	done;\

parse:
	$(GOCMD) build api-gateway/plugin/protoc-gen-parse
	protoset=`cat config.json|awk -F"[proto.set]" '/path|service/{print$0}'`;\
	aa=`echo $$protoset | sed s/[[:space:]]//g`;\
	bb=$${aa//\"\"/ };\
	proto_dir="";\
	for i in $${aa//\"\"/ }; do \
	j=$${i//\"/};\
	k=$${j//service:/};\
	l=$${k//path:/};\
	m=$${l//,/ };\
	n=($$m);\
	proto_dir=$$proto_dir" service/$${n[0]}/*.proto";\
	done;\
	$(PROTOC) -I. -Ithird_party --plugin=protoc-gen-parse --parse_out=. $$proto_dir

generate:
	protoset=`cat config.json|awk -F"[proto.set]" '/path|service/{print$0}'`;\
	aa=`echo $$protoset | sed s/[[:space:]]//g`;\
	bb=$${aa//\"\"/ };\
	proto_dir="";\
	for i in $${aa//\"\"/ }; do \
	j=$${i//\"/};\
	k=$${j//service:/};\
	l=$${k//path:/};\
	m=$${l//,/ };\
	n=($$m);\
	proto_dir=$$proto_dir" service/$${n[0]}/*.proto";\
	$(PROTOC) -Ithird_party -I.  --go_out=.  service/$${n[0]}/*.proto;\
	sed -i '/google\/api/d' service/$${n[0]}/*.pb.go;\
	done;\

FILES = `ls service/*.proto`
init-pb:
	echo "package loader"> loader/initial.go;\
#	echo 'import _ "api-gateway/service"'>> loader/initial.go;
	protoset=`cat config.json|awk -F"[proto.set]" '/path|service/{print$0}'`;\
	aa=`echo $$protoset | sed s/[[:space:]]//g`;\
	bb=$${aa//\"\"/ };\
	for i in $${aa//\"\"/ }; do \
	j=$${i//\"/};\
	k=$${j//service:/};\
	l=$${k//path:/};\
	m=$${l//,/ };\
	n=($$m);\
	echo 'import _ "api-gateway/service/'$${n[0]}'"'>> loader/initial.go;\
	done;\
	json=`cat parse.json`;\
	echo  "const PROTO_JSON = "$$json >> loader/initial.go;
	@rm parse.json protoc-gen-parse;
devel:
	build
	
test:
	$(GOCMD) test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
