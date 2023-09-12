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

// GetExitState returns an exit state if process has exited, otherwise it returns nil.
func GetExitState(p *process_manager.ManagedProcess) *pb.ProcessInformation_ExitState {
	code, err := p.ExitCode()
	if err != nil {
		return nil
	}
	return &pb.ProcessInformation_ExitState{
		Code: int64(code),
	}
}
