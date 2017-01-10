package main

import (
	"context"

	"google.golang.org/grpc"
)

func Forward(ctx context.Context, in *interface{}, opts ...grpc.CallOption) (*interface{}, error) {
	out := new(interface{})
	conn, err := grpc.Dial("endpoint", []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		return nil, err
	}
	err = grpc.Invoke(ctx, "/settings.Message/Operate", in, out, conn, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
