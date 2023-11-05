package managedprocess

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/olexnzarov/gofu/internal/daemon/processmanager"
	"github.com/olexnzarov/gofu/internal/process"
)

const (
	STATUS_RUNNING    = "running"
	STATUS_RESTARTING = "restarting"
	STATUS_STOPPED    = "stopped"
	STATUS_FAILED     = "failed"
)

type ManagedProcess struct {
	manager    *Manager
	outOptions process.OutOptions

	data      processmanager.ProcessData
	dataMutex *sync.RWMutex

	currentProcess *process.Process
	processMutex   *sync.Mutex

	interrupted      bool
	interruptChannel chan interface{}
	interruptMutex   *sync.Mutex

	autoRestartTries atomic.Uint32
}

func NewManagedProcess(manager *Manager, data processmanager.ProcessData) *ManagedProcess {
	return &ManagedProcess{
		manager:          manager,
		data:             data,
		processMutex:     &sync.Mutex{},
		dataMutex:        &sync.RWMutex{},
		interruptMutex:   &sync.Mutex{},
		interruptChannel: make(chan interface{}),
		outOptions:       process.NewOutOptions(manager.directories.LogDirectory, data.GetID()),
	}
}

func (mp *ManagedProcess) String() string {
	return fmt.Sprintf("ManagedProcess{name:'%s'}", mp.data.GetConfiguration().Name)
}

func (mp *ManagedProcess) GetRestarts() uint32 {
	return mp.autoRestartTries.Load()
}

func (mp *ManagedProcess) GetStatus() string {
	if mp.IsRunning() {
		return STATUS_RUNNING
	}
	if canProcessAutoRestart(mp) {
		return STATUS_RESTARTING
	}
	if code, _, _ := mp.GetExitState(); code > 0 {
		return STATUS_FAILED
	}
	return STATUS_STOPPED
}

func (mp *ManagedProcess) GetRunningProcess() (*process.Process, error) {
	if mp.currentProcess == nil {
		return nil, errors.New("process is not running")
	}
	return mp.currentProcess, nil
}

func (mp *ManagedProcess) GetInnerRunningProcess() (*os.Process, error) {
	process, err := mp.GetRunningProcess()
	if err != nil {
		return nil, err
	}
	return process.Inner(), nil
}

func (mp *ManagedProcess) GetPid() int {
	process, err := mp.GetInnerRunningProcess()
	if err != nil {
		return -1
	}
	return process.Pid
}

func (mp *ManagedProcess) GetId() string {
	return mp.data.GetID()
}

func (mp *ManagedProcess) GetStdoutPath() string {
	return mp.outOptions.Stdout
}

func (mp *ManagedProcess) GetData() processmanager.ProcessData {
	mp.dataMutex.RLock()
	defer mp.dataMutex.RUnlock()
	return mp.GetData()
}

func (mp *ManagedProcess) GetExitState() (int, time.Time, error) {
	if mp.currentProcess == nil {
		return 0, time.Time{}, nil
	}
	return mp.currentProcess.GetExitState()
}

func (mp *ManagedProcess) IsRunning() bool {
	_, _, err := mp.GetExitState()
	return err != nil
}

func (mp *ManagedProcess) spawn() (*process.Process, error) {
	if mp.IsRunning() {
		return nil, errors.New("process is already running")
	}

	mp.dataMutex.RLock()
	config := mp.data.GetConfiguration()
	startOptions := process.StartOptions{
		Out:         mp.outOptions,
		Command:     config.Command,
		Arguments:   config.Arguments,
		Environment: config.Environment,
	}
	mp.dataMutex.RUnlock()

	process, exit, err := process.Start(startOptions)
	mp.manager.Event.OnProcessSpawn.EmitAsync(ProcessSpawnEvent{Process: mp, Spawned: process, Error: err})
	if err != nil {
		return nil, err
	}
	mp.currentProcess = process

	go func() {
		exitCode := <-exit
		mp.manager.Event.OnProcessExit.Emit(ProcessExitEvent{Process: mp, Exited: process, ExitCode: exitCode})
	}()

	return process, nil
}

// safeSpawn is a thread-safe version of "spawn"
func (mp *ManagedProcess) safeSpawn() (*process.Process, error) {
	mp.processMutex.Lock()
	defer mp.processMutex.Unlock()
	return mp.spawn()
}

func (mp *ManagedProcess) interrupt(interrupted bool) {
	mp.interruptMutex.Lock()
	defer mp.interruptMutex.Unlock()

	if interrupted {
		if mp.interruptChannel != nil {
			select {
			case mp.interruptChannel <- nil:
			default:
			}
			close(mp.interruptChannel)
			mp.interruptChannel = nil
		}
	} else {
		mp.interruptChannel = make(chan interface{})
	}

	mp.interrupted = interrupted
}

func (mp *ManagedProcess) Stop() error {
	mp.interrupt(true)
	mp.processMutex.Lock()
	defer mp.processMutex.Unlock()

	if !mp.IsRunning() {
		return nil
	}

	return mp.currentProcess.Close()
}

func (mp *ManagedProcess) Restart() error {
	mp.interrupt(true)
	mp.processMutex.Lock()
	defer mp.processMutex.Unlock()

	mp.autoRestartTries.Store(0)
	mp.interrupt(false)

	event := ProcessForceRestartEvent{
		Process:    mp,
		WasRunning: mp.IsRunning(),
	}
	if event.WasRunning {
		mp.currentProcess.Close()
		mp.currentProcess.Inner().Wait()
	}
	event.Spawned, event.Error = mp.spawn()
	go mp.manager.Event.OnProcessForceRestart.Emit(event)

	return event.Error
}
