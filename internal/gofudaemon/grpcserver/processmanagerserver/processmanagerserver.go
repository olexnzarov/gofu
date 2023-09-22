package processmanagerserver

import (
	"github.com/olexnzarov/gofu/internal/gofudaemon/procservice"
	"github.com/olexnzarov/gofu/internal/logger"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/gofu"
)

type ProcessManagerServer struct {
	pb.UnimplementedProcessManagerServer
	log         logger.Logger
	directories *gofu.Directories
	service     *procservice.Service
}

func New(
	log logger.Logger,
	directories *gofu.Directories,
	service *procservice.Service,
) *ProcessManagerServer {
	return &ProcessManagerServer{
		log:         log,
		directories: directories,
		service:     service,
	}
}
