package managedprocess

import "time"

func (m *Manager) onProcessExit(event *ProcessExitEvent) {
	process := event.Process
	exitedPid := event.Exited.Inner().Pid

	m.log.Infof("%s process pid=%d exited with code %d", process, exitedPid, event.ExitCode)

	process.processMutex.Lock()
	defer process.processMutex.Unlock()

	if process.GetPid() != exitedPid {
		return
	}

	if canProcessAutoRestart(process) {
		m.log.Infof("%s autorestart is enabled")

		for canProcessAutoRestart(process) {
			process.autoRestartTries.Add(1)

			data := process.GetData()
			delay := data.GetRestartDelay()

			m.log.Infof("%s trying to restart in %s (%d/%d)", process, delay, process.autoRestartTries.Load(), data.GetRestartPolicy().MaxRetries)

			if waitInterrupt(process, delay) {
				m.log.Infof("%s received interrupt, autorestart is canceled", process)
				return
			}

			m.log.Infof("%s autorestarting...", process)
			_, err := process.spawn()
			if err == nil {
				return
			}
		}

		// At this point, process had autorestarts enabled, but failed to restart all tries.
		m.log.Infof("%s failed to autorestart", process)
	}

	process.interrupt(true)
}

func canProcessAutoRestart(p *ManagedProcess) bool {
	config := p.GetData().GetConfiguration()
	if p.interrupted || config.RestartPolicy == nil || !config.RestartPolicy.AutoRestart {
		return false
	}
	return config.RestartPolicy.MaxRetries == 0 || p.autoRestartTries.Load() < config.RestartPolicy.MaxRetries
}

func waitInterrupt(process *ManagedProcess, duration time.Duration) bool {
	waitChannel := make(chan interface{})
	go func() {
		time.Sleep(duration)
		waitChannel <- nil
	}()

	select {
	case <-process.interruptChannel:
		return true
	case <-waitChannel:
		return false
	}
}
