package process_manager

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/olexnzarov/gofu/pkg/gofu"
	"github.com/olexnzarov/gofu/pkg/process"
	"go.uber.org/zap"
)

const (
	STATUS_RUNNING    = "running"
	STATUS_RESTARTING = "restarting"
	STATUS_STOPPED    = "stopped"
	STATUS_FAILED     = "failed"
)

type ManagedProcess struct {
	log              *zap.Logger
	directories      *gofu.Directories
	data             *ProcessData
	dataMutex        *sync.RWMutex
	process          *process.Process
	processMutex     *sync.Mutex
	interrupted      bool
	autoRestartTries atomic.Uint32
}

func NewManagedProcess(log *zap.Logger, directories *gofu.Directories, data *ProcessData) *ManagedProcess {
	return &ManagedProcess{
		log:          log,
		data:         data,
		directories:  directories,
		processMutex: &sync.Mutex{},
		dataMutex:    &sync.RWMutex{},
	}
}

func (p *ManagedProcess) String() string {
	return fmt.Sprintf("process=%s", p.data.Id)
}

func (p *ManagedProcess) Status() string {
	if p.IsRunning() {
		return STATUS_RUNNING
	}
	if p.canAutoRestart() {
		return STATUS_RESTARTING
	}
	if code, _ := p.ExitCode(); code > 0 {
		return STATUS_FAILED
	}
	return STATUS_STOPPED
}

func (p *ManagedProcess) Process() (*process.Process, error) {
	if p.process == nil {
		return nil, errors.New("process is not running")
	}
	return p.process, nil
}

func (p *ManagedProcess) Inner() (*os.Process, error) {
	process, err := p.Process()
	if err != nil {
		return nil, err
	}
	return process.Inner(), nil
}

func (p *ManagedProcess) Pid() int {
	process, err := p.Inner()
	if err != nil {
		return -1
	}
	return process.Pid
}

func (p *ManagedProcess) Data() *ProcessData {
	p.dataMutex.RLock()
	defer p.dataMutex.RUnlock()
	return p.data
}

func (p *ManagedProcess) ExitCode() (int, error) {
	if p.process == nil {
		return 0, nil
	}
	return p.process.ExitCode()
}

func (p *ManagedProcess) IsRunning() bool {
	_, err := p.ExitCode()
	return err != nil
}

func (p *ManagedProcess) canAutoRestart() bool {
	config := p.Data().Configuration
	if p.interrupted || config.RestartPolicy == nil || !config.RestartPolicy.AutoRestart {
		return false
	}
	return config.RestartPolicy.MaxRetries == 0 || p.autoRestartTries.Load() < config.RestartPolicy.MaxRetries
}

func (p *ManagedProcess) onProcessExit(process *process.Process, exitCode int) {
	p.log.Sugar().Infof("%s: pid=%d exited with code %d", p, process.Inner().Pid, exitCode)

	p.processMutex.Lock()
	defer p.processMutex.Unlock()

	// Make sure that it's possible to manually kill the process and replace it with a new one.
	// For example, it's necessary for the Restart function.
	if p.process.Inner().Pid != process.Inner().Pid {
		return
	}

	if p.canAutoRestart() {
		p.log.Sugar().Infof("%s: automatic restart is enabled", p)

		defer func() { p.autoRestartTries.Store(0) }()

		for p.canAutoRestart() {
			p.autoRestartTries.Add(1)

			delay := p.data.AutoRestartDelay()
			p.log.Sugar().Infof(
				"%s: trying to restart in %s (%d/%d)",
				p, delay, p.autoRestartTries.Load(),
				p.data.Configuration.RestartPolicy.MaxRetries,
			)
			time.Sleep(delay)

			if p.interrupted {
				p.log.Sugar().Infof("%s: process was interrupted, canceling the autorestart", p)
				return
			}

			err := p.spawn()
			if err == nil {
				return
			}
		}

		p.interrupted = true
		p.log.Sugar().Infof("%s: failed to automatically restart")
	}
}

func (p *ManagedProcess) spawn() error {
	if p.IsRunning() {
		return errors.New("process is already running")
	}

	p.log.Sugar().Infof("%s: spawning a process...", p)

	p.dataMutex.RLock()
	startOptions := process.StartOptions{
		Out:         process.NewOutOptions(p.directories.LogDirectory, p.data.Id),
		Command:     p.data.Configuration.Command,
		Arguments:   p.data.Configuration.Arguments,
		Environment: p.data.Configuration.Environment,
	}
	p.dataMutex.RUnlock()

	process, exit, err := process.Start(startOptions)
	if err != nil {
		p.log.Sugar().Infof("%s: failed to spawn a process: %s", p, err.Error())
		return err
	}

	p.log.Sugar().Infof("%s: spawned a process pid=%d", p, process.Inner().Pid)
	p.process = process

	go func() {
		exitCode := <-exit
		p.onProcessExit(process, exitCode)
	}()

	return nil
}

func (p *ManagedProcess) Spawn() error {
	p.processMutex.Lock()
	defer p.processMutex.Unlock()
	p.interrupted = false
	return p.spawn()
}

func (p *ManagedProcess) Stop() error {
	// It's important to set interrupted before we acquire the process lock.
	// onProcessExit function can keep a lock when autorestarts are enabled.
	// Setting "interrupted = true" is a way to cancel the autorestarts.
	p.interrupted = true
	p.processMutex.Lock()
	defer p.processMutex.Unlock()
	if !p.IsRunning() {
		return errors.New("process is not running")
	}
	p.log.Sugar().Infof("%s: stopping the process", p)
	p.process.Close()
	return nil
}

func (p *ManagedProcess) Restart() error {
	p.processMutex.Lock()
	defer p.processMutex.Unlock()
	p.log.Sugar().Infof("%s: restarting the process", p)
	p.interrupted = false
	if p.IsRunning() {
		p.log.Sugar().Infof("%s: killing pid=%d", p, p.Pid())
		p.process.Close()
		p.process.Inner().Wait()
	}
	return p.spawn()
}
