package processmanagerservice

import (
	"github.com/olexnzarov/gofu/internal/daemon/processmanager/service/managedprocess"
)

type ProcessListFilter struct{}

func (s *Service) ListProcesses(filter *ProcessListFilter) []*managedprocess.ManagedProcess {
	return s.manager.Processes.All()
}

func (s *Service) FindProcess(nameOrPid string) (*managedprocess.ManagedProcess, error) {
	return s.manager.Processes.Find(nameOrPid)
}
