package grpc_server

import (
	"fmt"
	"net"

	"github.com/alexnzarov/gofu/internal/grpc_server/process_manager_server"
	"github.com/alexnzarov/gofu/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	config *Config
	log    *zap.Logger
	inner  *grpc.Server
}

func New(
	log *zap.Logger,
	config *Config,
	pms *process_manager_server.ProcessManagerServer,
) *Server {
	server := grpc.NewServer()

	pb.RegisterProcessManagerServer(server, pms)

	return &Server{
		log:    log,
		config: config,
		inner:  server,
	}
}

// Start establishes a TCP listener and serves the gRPC server on it.
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return err
	}
	s.log.Sugar().Infof("Listening on :%d", s.config.Port)
	go s.inner.Serve(listener)
	return nil
}

// Stop tries to gracefully stop the gRPC server.
func (s *Server) Stop() error {
	s.log.Info("Gracefully stopping the server...")
	s.inner.GracefulStop()
	return nil
}
