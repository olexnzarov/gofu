package procservice

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
		s.log.Infof("Starting the process '%s' (%s)", data.Configuration.Name, data.Id)
		go s.manager.Create(data)
	}

	return nil
}
