package procmanager

func (m *Manager) onProcessSpawn(event *ProcessSpawnEvent) {
	if event.Error != nil {
		m.log.Infof("%s failed to spawn a process: %s", event.Process, event.Error)
		return
	}

	m.log.Infof("%s spawned a process pid=%d", event.Process, event.Spawned.Inner().Pid)
}
