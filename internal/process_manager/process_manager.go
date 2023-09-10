package process_manager

import (
	"github.com/olexnzarov/gofu/internal/system_directory"
	"go.uber.org/zap"
)

type ProcessManager struct {
	Processes   *ProcessList
	log         *zap.Logger
	directories *system_directory.Config
}

func New(log *zap.Logger, directories *system_directory.Config) *ProcessManager {
	return &ProcessManager{
		log:         log,
		directories: directories,
		Processes:   NewProcessList(),
	}
}

func (r *ProcessManager) Start(data *ProcessData) (*ManagedProcess, error) {
	r.log.Sugar().Infof("starting a process %s", data.Id)
	process := NewManagedProcess(r.log, r.directories, data)
	r.Processes.add(process)
	return process, process.Spawn()
}
