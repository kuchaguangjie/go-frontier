package server

import (
	"net"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/singchia/frontier/api/controlplane/frontlas/v1"
)

func NewGRPCServer(ln net.Listener, svc v1.ClusterServiceServer) *grpc.Server {
	// new server
	opts := []grpc.ServerOption{
		grpc.Middleware(recovery.Recovery()),
		grpc.Listener(ln),
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterClusterServiceServer(srv, svc)
	return srv
}
