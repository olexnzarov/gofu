package procservice

import (
	"github.com/google/uuid"
	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
	"github.com/olexnzarov/gofu/pb"
)

func (s *Service) CreateProcess(config *pb.ProcessConfiguration) (*procmanager.ManagedProcess, error) {
	s.SanitizeProcessConfig(config)
	if err := s.ValidateProcessConfig(config); err != nil {
		return nil, err
	}

	data := &procmanager.ProcessData{
		Id:            uuid.New().String(),
		Configuration: config,
	}

	process, err := s.manager.Create(data)
	if err != nil {
		data.Configuration.Persist = false
		return nil, err
	}

	if data.Configuration.Persist {
		if err := s.storage.Upsert(data); err != nil {
			data.Configuration.Persist = false
			return nil, err
		}
	}

	return process, nil
}
