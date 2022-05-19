package rpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"guanyu.dev/arithmetic/internal/adapters/framework/left/grpc/pb"
	"guanyu.dev/arithmetic/internal/ports"
)

type Adapter struct {
	api ports.APIPort
	pb.UnimplementedArithmeticServiceServer
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
