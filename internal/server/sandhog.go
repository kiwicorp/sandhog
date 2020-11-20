package server

import (
	"context"

	sandhogpb "github.com/selftechio/sandhog/internal/api/sandhog"
	"github.com/selftechio/sandhog/internal/tunnel"
)

type sandhogServer struct {
	sandhogpb.UnimplementedSandhogServer
}

func newSandhogServer() *sandhogServer {
	return new(sandhogServer)
}

func (s *sandhogServer) StartTunnel(ctx context.Context, req *sandhogpb.StartTunnelRequest) (*sandhogpb.StartTunnelResponse, error) {
	resp := new(sandhogpb.StartTunnelResponse)

	_, err := tunnel.NewTunnel(req.GetAddress(), req.GetName())
	if err != nil {
		resp.Result = sandhogpb.Result_NOTOK
	} else {
		resp.Result = sandhogpb.Result_OK
		resp.Tunnel = &sandhogpb.StartTunnelResponse_TunnelSpecs{
			TunnelSpecs: &sandhogpb.SandhogPeer{
				PublicKey:  "abc",
				AllowedIps: make([]string, 1),
				ListenPort: 51820,
				Endpoint: &sandhogpb.SandhogPeer_EndpointValue{
					EndpointValue: "abc",
				},
				PersistentKeepalive: &sandhogpb.SandhogPeer_PersistentKeepaliveValue{
					PersistentKeepaliveValue: 25,
				},
			},
		}
	}

	return resp, nil
}
