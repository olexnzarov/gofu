package procservice

import (
	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
)

func (s *Service) FindProcess(nameOrPid string) (*procmanager.ManagedProcess, error) {
	return s.manager.Processes.Find(nameOrPid)
}
