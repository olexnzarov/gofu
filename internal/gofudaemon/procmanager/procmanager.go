package procmanager

import (
	"github.com/olexnzarov/gofu/internal/logger"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/gofu"
)

type Manager struct {
	Processes   *ProcessList
	log         logger.Logger
	directories *gofu.Directories
}

func New(log logger.Logger, directories *gofu.Directories) *Manager {
	manager := &Manager{
		Processes:   NewProcessList(),
		log:         log,
		directories: directories,
	}
	return manager
}

func (m *Manager) Create(data *ProcessData) (*ManagedProcess, error) {
	m.log.Infof("Creating a process: %s", data.Id)

	process := NewManagedProcess(m.log, m.directories, data)
	if err := m.Processes.add(process); err != nil {
		return nil, err
	}
	process.safeSpawn()

	return process, nil
}

func (r *Manager) Remove(process *ManagedProcess) {
	process.Stop()
	r.Processes.remove(process)
}

func (r *Manager) UpdateConfiguration(
	process *ManagedProcess,
	update func(currentConfig *pb.ProcessConfiguration) (*pb.ProcessConfiguration, error),
) error {
	process.dataMutex.Lock()
	defer process.dataMutex.Unlock()

	config, err := update(process.data.Configuration)
	if err != nil {
		return err
	}

	process.data.Configuration = config
	return nil
}