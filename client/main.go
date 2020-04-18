package main

import (
	"context"
	"fmt"
	sample "grpc-sample/app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("test-abc-1-czrw5lbr2a-uc.a.run.app:443", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := sample.NewHogeClient(conn)
	message := &sample.HelloRequest{Abc: "tama"}

	res, err := client.Hello(context.Background(), message)

	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
