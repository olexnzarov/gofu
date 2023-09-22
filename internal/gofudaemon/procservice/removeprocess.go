package procservice

func (s *Service) RemoveProcess(nameOrPid string) error {
	process, err := s.manager.Processes.Find(nameOrPid)
	if err != nil {
		return err
	}
	data := process.GetData()
	if data.Configuration.Persist {
		if err := s.storage.Delete(process.GetId()); err != nil {
			return err
		}
	}
	s.manager.Remove(process)
	return nil
}
