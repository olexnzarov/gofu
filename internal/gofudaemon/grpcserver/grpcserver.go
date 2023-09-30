package grpcserver

import (
	"fmt"
	"net"

	"github.com/olexnzarov/gofu/internal/gofudaemon/grpcserver/processmanagerserver"
	"github.com/olexnzarov/gofu/internal/logger"
	"github.com/olexnzarov/gofu/pb"
	"google.golang.org/grpc"
)

type Server struct {
	config *Config
	log    logger.Logger
	inner  *grpc.Server
}

func New(
	log logger.Logger,
	config *Config,
	pms *processmanagerserver.Server,
) *Server {
	server := grpc.NewServer()

	pb.RegisterProcessManagerServer(server, pms)

	return &Server{
		log:    log,
		config: config,
		inner:  server,
	}
}

func (s *Server) Target() string {
	return fmt.Sprintf(":%d", s.config.Port)
}

// Start establishes a TCP listener and serves the gRPC server on it.
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.Target())
	if err != nil {
		return err
	}
	s.log.Infof("Listening on :%d", s.config.Port)
	go s.inner.Serve(listener)
	return nil
}

// Stop tries to gracefully stop the gRPC server.
func (s *Server) Stop() error {
	s.log.Info("Gracefully stopping the server...")
	s.inner.GracefulStop()
	return nil
}
