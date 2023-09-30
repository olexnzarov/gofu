package processmanagerserver

import (
	"github.com/olexnzarov/gofu/internal/gofudaemon/procservice"
	"github.com/olexnzarov/gofu/internal/logger"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/gofu"
)

type Server struct {
	pb.UnimplementedProcessManagerServer
	log         logger.Logger
	directories *gofu.Directories
	service     *procservice.Service
}

func New(
	log logger.Logger,
	directories *gofu.Directories,
	service *procservice.Service,
) *Server {
	return &Server{
		log:         log,
		directories: directories,
		service:     service,
	}
}
