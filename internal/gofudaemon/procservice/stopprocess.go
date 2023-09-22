package procservice

import "github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"

func (s *Service) StopProcess(nameOrPid string) (*procmanager.ManagedProcess, error) {
	process, err := s.manager.Processes.Find(nameOrPid)
	if err != nil {
		return nil, err
	}
	return process, process.Stop()
}
