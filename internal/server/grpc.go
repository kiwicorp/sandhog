package server

import (
	"net"

	"github.com/hashicorp/go-hclog"
	sandhogpb "github.com/selftechio/sandhog/internal/api/sandhog"
	"google.golang.org/grpc"
)

var (
	log hclog.Logger
)

func init() {
	log = hclog.Default().Named("server")
}

// Start starts the sandhog server.
func Start() error {
	// fixme 23.11.2020: hardcoded listener address and port
	lis, err := net.Listen("tcp", "127.0.0.1:22222")
	if err != nil {
		return err
	}

	// fixme 23.11.2020: grpc server created with no options
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// fixme 23.11.2020: sandhog server created with no options
	sandhogpb.RegisterSandhogServer(grpcServer, newSandhogServer())

	// fixme 23.11.2020: sandhog server goroutine cannot be gracefully stopped
	// fixme 23.11.2020: sandhog server goroutine cannot gracefully stop sandhog in case of an
	//                   error
	go func() {
		log.Info("grpc server starting")

		err := grpcServer.Serve(lis)
		if err != nil {
			log.Error("grpc server encountered an error", "err", err)
		}
	}()

	return nil
}
