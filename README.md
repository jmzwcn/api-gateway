API Gateway

A gateway to automaticly provide RESTful API for gRPC

Usage:

1.Go get https://github.com/jmzwcn/api-gateway.git;

2.Tell protos where you will RESTful: e.g. put them into service direcotry;

3.Run "make";

4.Execute "api-gateway" in CMD.

How to define RESTful in *.proto: Add a [custom option](https://cloud.google.com/service-management/reference/rpc/google.api#http)
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