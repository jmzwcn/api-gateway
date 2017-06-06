package loader
import _ "github.com/api-gateway/example/echo/api"
import _ "github.com/api-gateway/example/helloworld/service"
const PROTO_JSON = "[{\"Package\":\"helloworld\",\"Service\":\"Greeter\",\"Method\":{\"name\":\"SayHello\",\"input_type\":\".helloworld.HelloRequest\",\"output_type\":\".helloworld.HelloReply\",\"options\":{}},\"Pattern\":{\"Verb\":\"POST\",\"Path\":\"/hello/{name}\",\"Body\":\"*\"}}]"