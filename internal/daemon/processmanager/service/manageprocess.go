package processmanagerservice

import (
	"github.com/olexnzarov/gofu/internal/daemon/processmanager/service/managedprocess"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/protomask"
	"google.golang.org/protobuf/proto"
)

func (s *Service) RestartProcess(nameOrPid string) (*managedprocess.ManagedProcess, error) {
	process, err := s.manager.Processes.Find(nameOrPid)
	if err != nil {
		return nil, err
	}
	return process, process.Restart()
}

func (s *Service) StopProcess(nameOrPid string) (*managedprocess.ManagedProcess, error) {
	process, err := s.manager.Processes.Find(nameOrPid)
	if err != nil {
		return nil, err
	}
	return process, process.Stop()
}

func (s *Service) RemoveProcess(nameOrPid string) error {
	process, err := s.manager.Processes.Find(nameOrPid)
	if err != nil {
		return err
	}
	data := process.GetData()
	if data.GetConfiguration().Persist {
		if err := s.storage.Delete(process.GetID()); err != nil {
			return err
		}
	}
	s.manager.Remove(process)
	return nil
}

func (s *Service) UpdateProcess(nameOrPid string, config *pb.ProcessConfiguration, mask protomask.FieldMask) (*managedprocess.ManagedProcess, error) {
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
				if err := s.storage.Upsert(&managedprocess.ProcessData{Id: process.GetID(), Configuration: updatedConfig}); err != nil {
					return nil, err
				}
			} else if currentConfig.Persist {
				if err := s.storage.Delete(process.GetID()); err != nil {
					return nil, err
				}
			}

			return updatedConfig, nil
		},
	)

	return process, err
}

func (s *Service) StartPersistent() error {
	s.log.Info("Trying to start the persistent processes...")

	processes, err := s.storage.List()
	if err != nil {
		return err
	}

	if len(processes) == 0 {
		s.log.Info("There are no persistent processes")
		return nil
	}

	s.log.Infof("Starting %d persistent processes...", len(processes))

	for _, data := range processes {
		s.log.Infof("Starting the process '%s' (%s)", data.GetConfiguration().Name, data.GetID())
		go s.manager.Create(data)
	}

	return nil
}
