package services

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-sample/app/proto"
)

type HogeService struct{}

func (h HogeService) Hello(context.Context, *sample.HelloRequest) (*sample.HelloResponse, error) {

	fmt.Println("hoge")

	return &sample.HelloResponse{}, status.Error(codes.NotFound, "error")
}
