package main

import (
	"log"
	"net"

	echo "github.com/api-gateway/example/echo/service"
	pb "github.com/api-gateway/example/helloworld/service"
	_ "github.com/api-gateway/types"
	"github.com/gogo/protobuf/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	conn, err := grpc.Dial("echo:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	reply, er := echo.NewEchoClient(conn).Ping(ctx, &types.Empty{})
	if er != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: "Hello " + in.Name + "-" + reply.String()}, nil
}

func (s *server) SayBye(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
