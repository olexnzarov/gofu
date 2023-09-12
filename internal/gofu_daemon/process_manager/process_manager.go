package process_manager

import (
	"github.com/olexnzarov/gofu/pkg/gofu"
	"go.uber.org/zap"
)

type ProcessManager struct {
	Processes   *ProcessList
	log         *zap.Logger
	directories *gofu.Directories
}

func New(log *zap.Logger, directories *gofu.Directories) *ProcessManager {
	return &ProcessManager{
		log:         log,
		directories: directories,
		Processes:   NewProcessList(),
	}
}

// Start creates a new managed process and spawns it. Managed process will always be returned, even if spawn errored.
func (r *ProcessManager) Start(data *ProcessData) (*ManagedProcess, error) {
	r.log.Sugar().Infof("starting a process %s", data.Id)
	process := NewManagedProcess(r.log, r.directories, data)
	r.Processes.add(process)
	return process, process.Spawn()
}

func (r *ProcessManager) Remove(process *ManagedProcess) {
	process.Stop()
	r.Processes.remove(process)
}
