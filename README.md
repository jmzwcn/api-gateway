# API Gateway

A native and easiest gateway to automaticly provide RESTful API for gRPC


## Usage:

1.Get code: [git clone https://github.com/jmzwcn/api-gateway.git];

2."make run" in api-gateway directory;

3.Try the URL in your proto.


How to define RESTful in *.proto: [[custom option](https://cloud.google.com/service-management/reference/rpc/google.api#http)]
   
   your_service.proto:
   ```protobuf
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