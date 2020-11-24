package server

import (
	"context"
	"os"

	uuid "github.com/satori/go.uuid"
	sandhogpb "github.com/selftechio/sandhog/internal/api/sandhog"
)

type sandhogServer struct {
	sandhogpb.UnimplementedSandhogServer
}

func newSandhogServer() *sandhogServer {
	return new(sandhogServer)
}

func (s *sandhogServer) StartTunnelNegotiations(ctx context.Context, req *sandhogpb.StartTunnelNegotiationsRequest) (*sandhogpb.StartTunnelNegotiationsResponse, error) {
	log.Info("starting tunnel negotiations", "tunnel-name", req.GetTunnelName(), "tunnel-uuid", req.GetTunnelUuid())

	hostname, err := os.Hostname()
	if err != nil {
		log.Warn("failed to start tunnel negotiations: failed to get the machine hostname", "err", err)
		return nil, err
	}

	resp := new(sandhogpb.StartTunnelNegotiationsResponse)
	resp.SandhogHostname = hostname
	resp.SandhogUuid = uuid.NewV4().String()

	return resp, nil
}
