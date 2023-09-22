package procservice

import (
	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/protomask"
	"google.golang.org/protobuf/proto"
)

func (s *Service) UpdateProcess(nameOrPid string, config *pb.ProcessConfiguration, mask protomask.FieldMask) (*procmanager.ManagedProcess, error) {
	process, err := s.manager.Processes.Find(nameOrPid)
	if err != nil {
		return nil, err
	}

	err = s.manager.UpdateConfiguration(
		process,
		func(currentConfig *pb.ProcessConfiguration) (*pb.ProcessConfiguration, error) {
			updatedConfig := proto.Clone(currentConfig).(*pb.ProcessConfiguration)
			if err := protomask.Update(updatedConfig, config, mask); err != nil {
				return nil, err
			}

			if err := s.ValidateProcessConfig(updatedConfig); err != nil {
				return nil, err
			}
			s.SanitizeProcessConfig(updatedConfig)

			if updatedConfig.Persist {
				if err := s.storage.Upsert(&procmanager.ProcessData{Id: process.GetId(), Configuration: updatedConfig}); err != nil {
					return nil, err
				}
			} else if currentConfig.Persist {
				if err := s.storage.Delete(process.GetId()); err != nil {
					return nil, err
				}
			}

			return updatedConfig, nil
		},
	)

	return process, err
}
