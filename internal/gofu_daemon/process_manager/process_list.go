package process_manager

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// ProcessList is a structure that contains a list of processes.
type ProcessList struct {
	processes      []*ManagedProcess
	processesMutex *sync.RWMutex
}

func NewProcessList() *ProcessList {
	return &ProcessList{
		processes:      []*ManagedProcess{},
		processesMutex: &sync.RWMutex{},
	}
}

// All returns a list of all processes. Thread-safe.
func (l *ProcessList) All() *[]*ManagedProcess {
	l.processesMutex.RLock()
	defer l.processesMutex.RUnlock()
	return &l.processes
}

// Add adds a process to the list. If process is already in the list, returns an error.
func (l *ProcessList) add(process *ManagedProcess) error {
	l.processesMutex.Lock()
	defer l.processesMutex.Unlock()

	// process duplicates, sanity check
	if _, err := l.getById(process.data.Id); err == nil {
		return fmt.Errorf("duplicate of a process '%s'", process.data.Id)
	}

	// unique names
	if _, err := l.find(process.data.Configuration.Name); err == nil {
		return fmt.Errorf("process with name '%s' already exists", process.data.Configuration.Name)
	}

	l.processes = append(l.processes, process)
	return nil
}

// Remove removes a process from the list. Thread-safe.
func (l *ProcessList) remove(process *ManagedProcess) {
	l.processesMutex.Lock()
	defer l.processesMutex.Unlock()

	for i, p := range l.processes {
		if p.Data().Id == process.Data().Id {
			// It may be not memory or speed-efficient, but it preserves the order.
			// I don't think this array will be of size that will have a problem with this.
			l.processes = append(l.processes[:i], l.processes[i+1:]...)
			break
		}
	}
}

// Get returns the process with given pid. If process is not in the registry, returns nil. Thread-safe.
func (l *ProcessList) GetByPid(pid int) (*ManagedProcess, error) {
	l.processesMutex.RLock()
	defer l.processesMutex.RUnlock()
	return l.getByPid(pid)
}

// GetById returns the process with given internal id. If process is not in the registry, returns nil. Thread-safe.
func (l *ProcessList) GetById(id string) (*ManagedProcess, error) {
	l.processesMutex.RLock()
	defer l.processesMutex.RUnlock()
	return l.getById(id)
}

// Find returns the process with given name or pid. If process is not in the registry, returns nil. Thread-safe.
func (l *ProcessList) Find(name string) (*ManagedProcess, error) {
	l.processesMutex.RLock()
	defer l.processesMutex.RUnlock()
	return l.find(name)
}

func (l *ProcessList) find(name string) (*ManagedProcess, error) {
	name = strings.TrimSpace(name)

	for _, p := range l.processes {
		if p.Data().Configuration.Name == name {
			return p, nil
		}
	}

	if regexp.MustCompile(`^\d+$`).MatchString(name) {
		pid, err := strconv.ParseInt(name, 10, 0)
		if err == nil {
			return l.GetByPid(int(pid))
		}
	}

	return nil, fmt.Errorf("unknown process '%s'", name)
}

func (l *ProcessList) getById(id string) (*ManagedProcess, error) {
	for _, p := range l.processes {
		if p.Data().Id == id {
			return p, nil
		}
	}

	return nil, fmt.Errorf("unknown process '%s'", id)
}

func (l *ProcessList) getByPid(pid int) (*ManagedProcess, error) {
	for _, p := range l.processes {
		if inner, err := p.Inner(); err == nil && inner.Pid == pid {
			return p, nil
		}
	}

	return nil, fmt.Errorf("unknown process '%d'", pid)
}
