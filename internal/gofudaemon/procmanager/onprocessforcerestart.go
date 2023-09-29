package procmanager

func (m *Manager) onProcessForceRestart(event *ProcessForceRestartEvent) {
	var previousState string
	if event.WasRunning {
		previousState = "running"
	} else {
		previousState = "stopped"
	}

	if event.Error != nil {
		m.log.Infof("%s failed to forcefully restart %s process: %s", event.Process, previousState, event.Error)
		return
	}

	m.log.Infof("%s forcefully restarted %s process, new process pid=%d", event.Process, previousState, event.Spawned.Inner().Pid)
}
