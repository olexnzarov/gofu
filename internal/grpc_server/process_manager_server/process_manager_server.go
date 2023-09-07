package process_manager_server

import (
	"github.com/olexnzarov/gofu/internal/process_registry"
	"github.com/olexnzarov/gofu/internal/system_directory"
	"github.com/olexnzarov/gofu/pb"
	"go.uber.org/zap"
)

type ProcessManagerServer struct {
	pb.UnimplementedProcessManagerServer
	log             *zap.Logger
	directories     *system_directory.Config
	processRegistry *process_registry.ProcessRegistry
}

func New(
	log *zap.Logger,
	directories *system_directory.Config,
	processRegistry *process_registry.ProcessRegistry,
) *ProcessManagerServer {
	return &ProcessManagerServer{
		log:             log,
		directories:     directories,
		processRegistry: processRegistry,
	}
}

// GetExitState returns an exit state if process has exited, otherwise it returns nil.
func GetExitState(p *process_registry.ManagedProcess) *pb.ProcessInformation_ExitState {
	if exit, err := p.GetExitState(); err != nil {
		exitState := pb.ProcessInformation_ExitState{
			Code: int32(exit.State.ExitCode()),
		}

		if exit.Error != nil {
			errMessage := (*exit.Error).Error()
			exitState.Error = &errMessage
		} else if exitState.Code != 0 {
			errMessage := exit.State.String()
			exitState.Error = &errMessage
		}

		return &exitState
	}

	return nil
}
