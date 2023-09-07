package process_registry

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/alexnzarov/gofu/process"
)

type ManagedProcess struct {
	mutex        *sync.RWMutex
	data         *ProcessData
	isRestarting bool

	// When changing the process, don't forget to remove the process_exit value.
	process     *process.Process
	processExit *ProcessExit
}

const (
	STATUS_RUNNING    = "running"
	STATUS_STOPPED    = "stopped"
	STATUS_ERROR      = "error"
	STATUS_RESTARTING = "restarting"
)

func NewManagedProcess(process *process.Process, data *ProcessData) *ManagedProcess {
	return &ManagedProcess{
		process: process,
		data:    data,
		mutex:   &sync.RWMutex{},
	}
}

func (mp *ManagedProcess) String() string {
	return fmt.Sprintf("'%s' (%d)", mp.data.Configuration.Name, mp.Pid())
}

func (mp *ManagedProcess) Data() *ProcessData {
	return mp.data
}

// Inner returns the underlying os.Process. Thread-safe.
func (mp *ManagedProcess) Inner() *os.Process {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()

	return mp.process.Inner()
}

// HasError returns whether the application exited with an error.
func (mp *ManagedProcess) HasError() bool {
	return mp.processExit != nil && (mp.processExit.State.ExitCode() > 0 || mp.processExit.Error != nil)
}

func (mp *ManagedProcess) GetExitState() (*ProcessExit, error) {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()

	if mp.processExit == nil {
		return nil, errors.New("process is still running")
	}

	return mp.processExit, nil
}

// Status returns a readable status of the process. See STATUS_* constants for possible values.
// It doesn't store the status, it's always calculated on runtime based on number of internal states.
func (mp *ManagedProcess) Status() string {
	if mp.Pid() > 0 {
		return STATUS_RUNNING
	}
	if mp.isRestarting {
		return STATUS_RESTARTING
	}
	if mp.HasError() {
		return STATUS_ERROR
	}
	return STATUS_STOPPED
}

// Pid returns the inner process' pid. Returns -1 if process isn't running.
func (mp *ManagedProcess) Pid() int {
	inner := mp.Inner()
	if inner == nil {
		return -1
	}
	return inner.Pid
}

// Restart kills the process, waits for it to exit, and spawns a new one with the same options.
// If process was previously awaited, the exit channel will receive a value and the channel will be closed.
// This method also resets the autoRestarts counter.
// Thread-safe.
func (mp *ManagedProcess) Restart(delay time.Duration) error {
	return mp.restart(delay, false)
}

// restarts does the same thing as Restart, with the exception that it resets the autoRestarts counter only if internal is false.
func (mp *ManagedProcess) restart(delay time.Duration, internal bool) error {
	mp.process.Close()
	mp.process.Inner().Wait()

	mp.isRestarting = true

	time.Sleep(delay)

	mp.mutex.Lock()
	defer mp.mutex.Unlock()

	p, err := process.Start(*mp.process.Options())
	if err != nil {
		mp.isRestarting = false
		return err
	}

	// Reset the autoRestarts counter if user manually restarted the process.
	if !internal {
		mp.data.autoRestarts = 0
	}

	mp.data.Restarts++
	mp.process = p
	mp.processExit = nil
	mp.isRestarting = false

	return nil
}

// autoRestart calls Restart and increases autoRestarts counter. Thread-safe.
func (mp *ManagedProcess) autoRestart(delay time.Duration) error {
	mp.mutex.Lock()
	mp.data.autoRestarts++
	mp.mutex.Unlock()

	return mp.restart(delay, true)
}

// Wait awaits the process, returning an exit channel. If the process was already awaited, returns an error.
// Thread-safe.
func (mp *ManagedProcess) Wait() (<-chan *ProcessExit, error) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()

	if mp.process.IsAwaited() {
		return nil, errors.New("process is already awaited")
	}

	exit := make(chan *ProcessExit)

	go func() {
		state, err := mp.process.Wait()
		mp.processExit = NewProcessExit(state, err)
		exit <- mp.processExit
		close(exit)
	}()

	return exit, nil
}
