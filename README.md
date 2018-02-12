# API Gateway
A native and easy gateway to provide RESTful API for gRPC, inspired by [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway), but fully automation.

## Prepare:
* Download&&Install golang/protoc/docker and Env setting.

## Usage:
* Get code: [git clone https://github.com/jmzwcn/api-gateway.git];
* Tell your macro-services directory in Makefile;
* "make run";
* Try the URL in your proto.


How to define RESTful in *.proto: [[custom option](https://cloud.google.com/service-management/reference/rpc/google.api#http)]
   
   your_service.proto:
   ```diff
    syntax = "proto3";
    package example;
   +
   +import "google/api/annotations.proto";
   +
    message StringMessage {
      string value = 1;
    }
    
    service YourService {
   -  rpc Echo(StringMessage) returns (StringMessage) {}
   +  rpc Echo(StringMessage) returns (StringMessage) {
   +    option (google.api.http) = {
   +      post: "/v1/example/echo"
   +      body: "*"
   +    };
   +  }
    }
   ```
   
Enjoy it!

<hr/>
For custom setting, please refer to example&&Makefile.
