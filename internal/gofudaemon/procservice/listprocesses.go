package procservice

import "github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"

type ProcessListFilter struct{}

func (s *Service) ListProcesses(filter *ProcessListFilter) []*procmanager.ManagedProcess {
	return s.manager.Processes.All()
}
