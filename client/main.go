package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	sample "grpc-sample/app/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	grpcEndpoint = flag.String("grpc_endpoint", "", "The gRPC Endpoint of the Server")
)

func main() {

	flag.Parse()
	if *grpcEndpoint == "" {
		log.Fatal("[main] unable to start client without gRPC endpoint to server")
	}

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	log.Printf("Connecting to gRPC Service [%s]", *grpcEndpoint)
	conn, err := grpc.Dial(*grpcEndpoint, opts...)

	//sampleなのでwithInsecure
	// conn, err := grpc.Dial("test-abc-1-czrw5lbr2a-uc.a.run.app:443", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := sample.NewHogeClient(conn)
	message := &sample.HelloRequest{Abc: "tama"}

	res, err := client.Hello(context.Background(), message)

	if err != nil {
		fmt.Printf("error::%#v \n", err)
	}

	fmt.Printf("result:%#v \n", res.Xyz)
}
