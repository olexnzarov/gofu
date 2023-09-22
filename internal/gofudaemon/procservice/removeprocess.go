package procservice

func (s *Service) RemoveProcess(nameOrPid string) error {
	process, err := s.manager.Processes.Find(nameOrPid)
	if err != nil {
		return err
	}
	data := process.Data()
	if data.Configuration.Persist {
		if err := s.storage.Delete(process.Id()); err != nil {
			return err
		}
	}
	s.manager.Remove(process)
	return nil
}
