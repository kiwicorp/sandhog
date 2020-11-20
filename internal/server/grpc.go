package server

import (
	"net"

	sandhogpb "github.com/selftechio/sandhog/internal/api/sandhog"
	"google.golang.org/grpc"
)

// Start starts the sandhog server.
func Start() error {
	lis, err := net.Listen("tcp", "127.0.0.1:22222")
	if err != nil {
		return err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	sandhogpb.RegisterSandhogServer(grpcServer, newSandhogServer())

	return grpcServer.Serve(lis)
}
