package mgmt

import (
	"context"

	"github.com/caos/zitadel/pkg/grpc/management"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	service management.ManagementServiceServer
}

func New(svc management.ManagementServiceServer) *Server {
	return &Server{
		service: svc,
	}
}

func (s *Server) RegisterGRPC(srv *grpc.Server) {
	management.RegisterManagementServiceServer(srv, s.service)
}

func (s *Server) RegisterRESTGateway(ctx context.Context, grpcMux *runtime.ServeMux) error {
	conn, err := grpc.Dial(":50002", grpc.WithInsecure())
	if err != nil {
		return err
	}

	return management.RegisterManagementServiceHandler(ctx, grpcMux, conn)
}

func (s *Server) ServicePrefix() string {
	return "/management/v1"
}
