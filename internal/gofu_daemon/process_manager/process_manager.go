package process_manager

import (
	"github.com/olexnzarov/gofu/pkg/gofu"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type ProcessManager struct {
	Processes   *ProcessList
	log         *zap.Logger
	directories *gofu.Directories
	storage     PersistentStorage
}

func New(log *zap.Logger, directories *gofu.Directories, storage PersistentStorage) *ProcessManager {
	manager := &ProcessManager{
		Processes:   NewProcessList(),
		log:         log,
		directories: directories,
		storage:     storage,
	}
	return manager
}

func (r *ProcessManager) StartPersistent() error {
	r.log.Info("Trying to start the persistent processes...")

	processes, err := r.storage.List()
	if err != nil {
		return err
	}

	if len(processes) == 0 {
		r.log.Info("There are no persistent processes")
		return nil
	}

	r.log.Sugar().Infof("Starting %d persistent processes...", len(processes))

	for _, data := range processes {
		r.log.Sugar().Infof("Starting process '%s' (%s)", data.Configuration.Name, data.Id)
		go r.Start(data)
	}

	return nil
}

// Start creates a new managed process and spawns it.
// May return a process and an error if it failed to spawn, but the managed process was created.
func (r *ProcessManager) Start(data *ProcessData) (*ManagedProcess, error) {
	r.log.Sugar().Infof("Starting a process %s", data.Id)

	process := NewManagedProcess(r.log, r.directories, data)
	if err := r.Processes.add(process); err != nil {
		return nil, err
	}

	err := process.Spawn()
	if data.Configuration.Persist {
		err = multierr.Append(err, r.storage.Upsert(data))
	}

	return process, err
}

func (r *ProcessManager) Remove(process *ManagedProcess) error {
	err := process.Stop()
	r.Processes.remove(process)
	if process.data.Configuration.Persist {
		err = multierr.Append(err, r.storage.Delete(process.data.Id))
	}
	return err
}
