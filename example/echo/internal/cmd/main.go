package main

import (
	"log"
	"net"

	pb "github.com/api-gateway/example/echo/service"
	"github.com/gogo/protobuf/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type simpleEchoServer struct {
}

func (s *simpleEchoServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Text: req.Text,
	}, nil
}

func (s *simpleEchoServer) Ping(ctx context.Context, req *types.Empty) (*types.Timestamp, error) {
	return types.TimestampNow(), nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &simpleEchoServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
