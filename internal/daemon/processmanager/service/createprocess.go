package processmanagerservice

import (
	"github.com/google/uuid"
	"github.com/olexnzarov/gofu/internal/daemon/processmanager/service/managedprocess"
	"github.com/olexnzarov/gofu/pb"
)

func (s *Service) CreateProcess(config *pb.ProcessConfiguration) (*managedprocess.ManagedProcess, error) {
	s.SanitizeProcessConfig(config)
	if err := s.ValidateProcessConfig(config); err != nil {
		return nil, err
	}

	data := &managedprocess.ProcessData{
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
