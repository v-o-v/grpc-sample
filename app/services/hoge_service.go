package services

import (
	"context"
	"fmt"
	sample "grpc-sample/app/proto"
)

type HogeService struct{}

func (h HogeService) Hello(context.Context, *sample.HelloRequest) (*sample.HelloResponse, error) {

	fmt.Println("hoge")

	return &sample.HelloResponse{Xyz: "wakuwaku"}, nil
}
