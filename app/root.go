package app

import (
	sample "grpc-sample/app/proto"
	"grpc-sample/app/services"
	"log"
	"net"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// InitRoot 初期ルート
func InitRoot() {

	port := os.Getenv("PORT")
	if port == "" {
		port = ""
	}
	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Router()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func Router() *grpc.Server {

	grpcServer := grpc.NewServer()
	sample.RegisterHogeServer(grpcServer, &services.HogeService{})

	// listの取得
	reflection.Register(grpcServer)
	return grpcServer
}
