# Variables
GOCMD=go
PROTOC=../protoc/protoc-3.1.0-linux-x86_64/bin/protoc

generate:
	$(PROTOC) -I../third_party -I. --gogofast_out=plugins=grpc:. helloworld/service/*.proto
	sed -i '/google\/api/d' helloworld/service/*.pb.go
	$(PROTOC) -I../third_party -I. -I../../gogo/protobuf --gogofast_out=plugins=grpc:. echo/api/*.proto
	sed -i '/google\/api/d' echo/api/*.pb.go
	sed -i '/gogoproto/d' echo/api/*.pb.go
