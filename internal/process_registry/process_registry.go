package process_registry

import (
	"regexp"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/zap"
)

type ProcessRegistry struct {
	log            *zap.Logger
	processes      []*ManagedProcess
	processesMutex *sync.RWMutex
}

func New(log *zap.Logger) *ProcessRegistry {
	return &ProcessRegistry{
		log:            log,
		processes:      []*ManagedProcess{},
		processesMutex: &sync.RWMutex{},
	}
}

// Processes returns the registry list.
func (r *ProcessRegistry) Processes() *[]*ManagedProcess {
	return &r.processes
}

// tryAutoRestart tries to restart the process until successful, or until it exhausts all possible autorestarts.
// Returns whether it succeeded in restarting the process.
func (r *ProcessRegistry) tryAutoRestart(p *ManagedProcess) bool {
	for p.data.HasAutoRestart() {
		delay := p.data.AutoRestartDelay()

		r.log.Sugar().Infof("%s: trying to restart the process in %s", p, delay)
		err := p.autoRestart(delay)

		if err != nil {
			r.log.Sugar().Errorf("%s: failed to restart the process\nerror: %s", p, err)
		} else {
			r.log.Sugar().Infof("%s: process was restarted", p)
			return true
		}
	}

	return false
}

func (r *ProcessRegistry) startWatchLoop(p *ManagedProcess) {
	r.log.Sugar().Infof("%s: watch loop started", p)

	for {
		exitChannel, err := p.Wait()

		if err != nil {
			// In theory, this should never happen, but better be safe than create a bunch of zombies.
			// We kill the process and try to restart it to get the control of it.
			r.log.Sugar().Errorf("%s: process became detached: %s", p, err.Error())
			p.process.Close()
		} else {
			exit := <-exitChannel
			r.log.Sugar().Infof("%s: process exited with code %d (%s)", p, exit.State.ExitCode(), exit.State)
		}

		if !r.tryAutoRestart(p) {
			break
		}
	}

	p.process.Close()
	r.log.Sugar().Infof("%s: watch loop ended", p)
}

// Watch adds the process to the registry and waits for it, returns a channel that closes when the process finishes. Thread-safe.
func (r *ProcessRegistry) Watch(p *ManagedProcess) {
	r.processesMutex.Lock()
	defer r.processesMutex.Unlock()

	r.processes = append(r.processes, p)

	go r.startWatchLoop(p)
}

// Get returns the process with given pid. If process is not in the registry, returns nil. Thread-safe.
func (r *ProcessRegistry) Get(pid int) *ManagedProcess {
	r.processesMutex.RLock()
	defer r.processesMutex.RUnlock()

	for _, p := range r.processes {
		if p.process.Inner().Pid == pid {
			return p
		}
	}

	return nil
}

// GetById returns the process with given internal id. If process is not in the registry, returns nil. Thread-safe.
func (r *ProcessRegistry) GetById(id string) *ManagedProcess {
	r.processesMutex.RLock()
	defer r.processesMutex.RUnlock()

	for _, p := range r.processes {
		if p.data.Id == id {
			return p
		}
	}

	return nil
}

// Find returns the process with given name or pid. If process is not in the registry, returns nil. Thread-safe.
func (r *ProcessRegistry) Find(name string) *ManagedProcess {
	r.processesMutex.RLock()
	defer r.processesMutex.RUnlock()

	name = strings.TrimSpace(name)

	for _, p := range r.processes {
		if p.data.Configuration.Name == name {
			return p
		}
	}

	if regexp.MustCompile(`^\d+$`).MatchString(name) {
		pid, err := strconv.ParseInt(name, 10, 0)
		if err != nil {
			return nil
		}
		return r.Get(int(pid))
	}

	return nil
}
