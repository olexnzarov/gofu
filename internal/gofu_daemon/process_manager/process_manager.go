package process_manager

import (
	"fmt"

	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"github.com/olexnzarov/protomask"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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

func (r *ProcessManager) UpdateConfiguration(
	process *ManagedProcess,
	config *pb.ProcessConfiguration,
	updateMask *fieldmaskpb.FieldMask,
) error {
	process.dataMutex.Lock()
	defer process.dataMutex.Unlock()

	previousConfig := proto.Clone(process.data.Configuration).(*pb.ProcessConfiguration)
	err := protomask.Update(process.data.Configuration, config, updateMask)

	if err != nil && previousConfig.Persist != config.Persist {
		if config.Persist {
			err = r.storage.Upsert(process.data)
		} else {
			err = r.storage.Delete(process.data.Id)
		}
	}

	// Revert the changes if something went wrong.
	if err != nil {
		process.data.Configuration = previousConfig
		return fmt.Errorf("failed to update the process: %s", err.Error())
	}

	return nil
}

func (r *ProcessManager) Remove(process *ManagedProcess) error {
	if process.data.Configuration.Persist {
		err := r.storage.Delete(process.data.Id)
		if err != nil {
			return err
		}
	}

	process.Stop()
	r.Processes.remove(process)

	return nil
}
