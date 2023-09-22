package procservice

import "github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"

func (s *Service) RestartProcess(nameOrPid string) (*procmanager.ManagedProcess, error) {
	process, err := s.manager.Processes.Find(nameOrPid)
	if err != nil {
		return nil, err
	}
	return process, process.Restart()
}
