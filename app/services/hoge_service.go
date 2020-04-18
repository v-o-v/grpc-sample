package services

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	sample "grpc-sample/app/proto"
)

type HogeService struct{}

func (h HogeService) Hello(c context.Context, in *sample.HelloRequest) (*sample.HelloResponse, error) {

	if in.Abc == "" {
		return &sample.HelloResponse{}, status.Error(codes.InvalidArgument, "parameter error")
	}
	fmt.Println("hoge")

	return &sample.HelloResponse{Xyz: "wakuwaku"}, nil
}
