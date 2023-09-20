package process_manager_server

import (
	"github.com/olexnzarov/gofu/internal/gofu_daemon/process_manager"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"go.uber.org/zap"
)

type ProcessManagerServer struct {
	pb.UnimplementedProcessManagerServer
	log            *zap.Logger
	directories    *gofu.Directories
	processManager *process_manager.ProcessManager
}

func New(
	log *zap.Logger,
	directories *gofu.Directories,
	processManager *process_manager.ProcessManager,
) *ProcessManagerServer {
	return &ProcessManagerServer{
		log:            log,
		directories:    directories,
		processManager: processManager,
	}
}
