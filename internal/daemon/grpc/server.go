package grpc

import (
	"fmt"
	"net"

	"github.com/olexnzarov/gofu/internal/daemon/logger"
	"google.golang.org/grpc"
)

type Server struct {
	config ServerConfig
	log    logger.Logger
	server *grpc.Server
}

func NewServer(
	log logger.Logger,
	config ServerConfig,
) *Server {
	server := grpc.NewServer()

	// TODO:
	//pb.RegisterProcessManagerServer(server, pms)

	return &Server{
		log:    log,
		config: config,
		server: server,
	}
}

func (s *Server) Raw() *grpc.Server {
	return s.server
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
	s.log.Infof("Listening on %s", s.Target())
	go s.server.Serve(listener)
	return nil
}

// Stop tries to gracefully stop the gRPC server.
func (s *Server) Stop() error {
	s.log.Info("Gracefully stopping the server...")
	s.server.GracefulStop()
	return nil
}
