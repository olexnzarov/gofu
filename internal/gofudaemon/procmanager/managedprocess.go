package procmanager

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/olexnzarov/gofu/internal/logger"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"github.com/olexnzarov/gofu/pkg/process"
)

const (
	STATUS_RUNNING    = "running"
	STATUS_RESTARTING = "restarting"
	STATUS_STOPPED    = "stopped"
	STATUS_FAILED     = "failed"
)

type ManagedProcess struct {
	log              logger.Logger
	outOptions       process.OutOptions
	data             *ProcessData
	dataMutex        *sync.RWMutex
	process          *process.Process
	processMutex     *sync.Mutex
	interrupted      bool
	interruptChannel chan interface{}
	interruptMutex   *sync.Mutex
	autoRestartTries atomic.Uint32
}

func NewManagedProcess(log logger.Logger, directories *gofu.Directories, data *ProcessData) *ManagedProcess {
	return &ManagedProcess{
		log:              log,
		data:             data,
		processMutex:     &sync.Mutex{},
		dataMutex:        &sync.RWMutex{},
		interruptMutex:   &sync.Mutex{},
		interruptChannel: make(chan interface{}),
		outOptions:       process.NewOutOptions(directories.LogDirectory, data.Id),
	}
}

func (p *ManagedProcess) String() string {
	return fmt.Sprintf("Process '%s'", p.data.Configuration.Name)
}

func (p *ManagedProcess) GetRestarts() uint32 {
	return p.autoRestartTries.Load()
}

func (p *ManagedProcess) GetStatus() string {
	if p.IsRunning() {
		return STATUS_RUNNING
	}
	if p.canAutoRestart() {
		return STATUS_RESTARTING
	}
	if code, _ := p.GetExitCode(); code > 0 {
		return STATUS_FAILED
	}
	return STATUS_STOPPED
}

func (p *ManagedProcess) GetRunningProcess() (*process.Process, error) {
	if p.process == nil {
		return nil, errors.New("process is not running")
	}
	return p.process, nil
}

func (p *ManagedProcess) GetInnerRunningProcess() (*os.Process, error) {
	process, err := p.GetRunningProcess()
	if err != nil {
		return nil, err
	}
	return process.Inner(), nil
}

func (p *ManagedProcess) GetProcessId() int {
	process, err := p.GetInnerRunningProcess()
	if err != nil {
		return -1
	}
	return process.Pid
}

func (p *ManagedProcess) GetId() string {
	return p.data.Id
}

func (p *ManagedProcess) GetStdoutPath() string {
	return p.outOptions.Stdout
}

func (p *ManagedProcess) GetData() *ProcessData {
	p.dataMutex.RLock()
	defer p.dataMutex.RUnlock()
	return p.data
}

func (p *ManagedProcess) GetExitCode() (int, error) {
	if p.process == nil {
		return 0, nil
	}
	return p.process.ExitCode()
}

func (p *ManagedProcess) IsRunning() bool {
	_, err := p.GetExitCode()
	return err != nil
}

func (p *ManagedProcess) waitWithInterrupt(duration time.Duration) {
	waitChannel := make(chan interface{})
	go func() {
		time.Sleep(duration)
		waitChannel <- nil
	}()

	select {
	case <-p.interruptChannel:
	case <-waitChannel:
	}
}

func (p *ManagedProcess) canAutoRestart() bool {
	config := p.GetData().Configuration
	if p.interrupted || config.RestartPolicy == nil || !config.RestartPolicy.AutoRestart {
		return false
	}
	return config.RestartPolicy.MaxRetries == 0 || p.autoRestartTries.Load() < config.RestartPolicy.MaxRetries
}

func (p *ManagedProcess) onProcessExit(process *process.Process, exitCode int) {
	p.log.Infof("%s: pid=%d exited with code %d", p, process.Inner().Pid, exitCode)

	p.processMutex.Lock()
	defer p.processMutex.Unlock()

	// Make sure that it's possible to manually kill the process and replace it with a new one.
	// For example, it's necessary for the Restart function.
	if p.process.Inner().Pid != process.Inner().Pid {
		return
	}

	if p.canAutoRestart() {
		p.log.Infof("%s: automatic restart is enabled", p)

		for p.canAutoRestart() {
			p.autoRestartTries.Add(1)

			delay := p.data.GetRestartDelay()
			p.log.Infof(
				"%s: trying to restart in %s (%d/%d)",
				p, delay, p.autoRestartTries.Load(),
				p.data.GetRestartPolicy().MaxRetries,
			)

			p.waitWithInterrupt(delay)

			if p.interrupted {
				p.log.Infof("%s: process was interrupted, canceling the autorestart", p)
				return
			}

			err := p.spawn()
			if err == nil {
				return
			}
		}

		p.interrupt(true)
		p.log.Infof("%s: failed to automatically restart")
	} else {
		p.interrupt(true)
	}
}

func (p *ManagedProcess) spawn() error {
	if p.IsRunning() {
		return errors.New("process is already running")
	}

	p.log.Infof("%s: spawning a process...", p)

	p.dataMutex.RLock()
	startOptions := process.StartOptions{
		Out:         p.outOptions,
		Command:     p.data.Configuration.Command,
		Arguments:   p.data.Configuration.Arguments,
		Environment: p.data.Configuration.Environment,
	}
	p.dataMutex.RUnlock()

	process, exit, err := process.Start(startOptions)
	if err != nil {
		p.log.Infof("%s: failed to spawn a process: %s", p, err)
		return err
	}

	p.log.Infof("%s: spawned a process pid=%d", p, process.Inner().Pid)
	p.process = process

	go func() {
		exitCode := <-exit
		p.onProcessExit(process, exitCode)
	}()

	return nil
}

// safeSpawn is a thread-safe version of "spawn".
func (p *ManagedProcess) safeSpawn() error {
	p.processMutex.Lock()
	defer p.processMutex.Unlock()
	return p.spawn()
}

func (p *ManagedProcess) interrupt(interrupted bool) {
	p.interruptMutex.Lock()
	defer p.interruptMutex.Unlock()
	if interrupted {
		if p.interruptChannel != nil {
			select {
			case p.interruptChannel <- nil:
			default:
			}
			close(p.interruptChannel)
			p.interruptChannel = nil
		}
	} else {
		p.interruptChannel = make(chan interface{})
	}
	p.interrupted = interrupted
}

func (p *ManagedProcess) Stop() error {
	p.interrupt(true)
	p.processMutex.Lock()
	defer p.processMutex.Unlock()
	if !p.IsRunning() {
		return nil
	}
	p.log.Infof("%s: stopping the process", p)
	return p.process.Close()
}

func (p *ManagedProcess) Restart() error {
	p.interrupt(true)
	p.processMutex.Lock()
	defer p.processMutex.Unlock()
	p.log.Infof("%s: restarting the process", p)
	p.autoRestartTries.Store(0)
	p.interrupt(false)
	if p.IsRunning() {
		p.log.Infof("%s: killing pid=%d", p, p.GetProcessId())
		p.process.Close()
		p.process.Inner().Wait()
	}
	return p.spawn()
}
