package process

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/olexnzarov/gofu/pkg/envfmt"
	"go.uber.org/multierr"
)

var (
	ErrProcessRunning = errors.New("process is still running")
)

type OutOptions struct {
	Stdin  string
	Stdout string
	Stderr string
}

type StartOptions struct {
	Out              OutOptions
	Command          string
	Arguments        []string
	Environment      map[string]string
	WorkingDirectory string
}

type Process struct {
	stdin     *os.File
	stdout    *os.File
	stderr    *os.File
	inner     *os.Process
	options   *StartOptions
	exitCode  *int
	startTime time.Time
	stopTime  time.Time
}

func NewOutOptions(outDirectory string, id string) OutOptions {
	return OutOptions{
		Stdin:  os.DevNull,
		Stdout: fmt.Sprintf("%s/%s.log", outDirectory, id),
		Stderr: fmt.Sprintf("%s/%s.log", outDirectory, id),
	}
}

// prepareProcess opens necessary file descriptors and prepares the process struct. It doesn't spawn the inner process, it will be nil.
// stdout and stderr are read-only for everyone except the owner and the group.
func prepareProcess(options *StartOptions) (*Process, error) {
	stdin, err := os.Open(options.Out.Stdin)
	if err != nil {
		return nil, err
	}

	stdout, err := os.OpenFile(options.Out.Stdout, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0664)
	if err != nil {
		stdin.Close()
		return nil, err
	}

	stderr, err := os.OpenFile(options.Out.Stderr, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0664)
	if err != nil {
		stdin.Close()
		stderr.Close()
		return nil, err
	}

	return &Process{
		stdin:    stdin,
		stdout:   stdout,
		stderr:   stderr,
		options:  options,
		exitCode: nil,
	}, nil
}

func (p *Process) Inner() *os.Process {
	return p.inner
}

func (p *Process) Options() *StartOptions {
	return p.options
}

func (p *Process) StartedAt() time.Time {
	return p.startTime
}

func (p *Process) StoppedAt() (time.Time, error) {
	if p.exitCode == nil {
		return time.Time{}, ErrProcessRunning
	}
	return p.stopTime, nil
}

func (p *Process) ExitCode() (int, error) {
	if p.exitCode == nil {
		return 0, ErrProcessRunning
	}
	return *p.exitCode, nil
}

// Close kills the process and releases all associated resources.
func (p *Process) Close() error {
	return multierr.Append(
		p.inner.Kill(),
		p.release(),
	)
}

// Release closes all file descriptors.
func (p *Process) release() error {
	return multierr.Combine(
		p.stdin.Close(),
		p.stderr.Close(),
		p.stdout.Close(),
	)
}

// Start starts a process with given options.
func Start(options StartOptions) (*Process, <-chan int, error) {
	bin, err := exec.LookPath(options.Command)
	if err != nil {
		return nil, nil, err
	}

	// argv should start with the program name
	program := strings.Fields(options.Command)[0]
	arguments := append([]string{program}, options.Arguments...)

	// Combine environment variables
	env := os.Environ()
	if options.Environment != nil {
		optEnv := envfmt.ToKeyValueArray(options.Environment)
		env = append(env, optEnv...)
	}

	if options.Arguments == nil {
		options.Arguments = []string{}
	}

	process, err := prepareProcess(&options)
	if err != nil {
		return nil, nil, err
	}

	attributes := os.ProcAttr{
		Dir: options.WorkingDirectory,
		Env: env,
		Files: []*os.File{
			process.stdin,
			process.stdout,
			process.stderr,
		},
		// These attributes are system-specific.
		// Check sys_proc_attr_<os>.go files to see how it handles children on different systems.
		Sys: newSysProcAttr(),
	}

	p, err := os.StartProcess(bin, arguments, &attributes)
	if err != nil {
		return nil, nil, err
	}
	process.inner = p
	process.startTime = time.Now()

	exitChannel := make(chan int)
	go func() {
		state, err := p.Wait()
		code := -1
		if err == nil {
			code = state.ExitCode()
		}
		process.exitCode = &code
		process.stopTime = time.Now()
		exitChannel <- *process.exitCode
		close(exitChannel)
	}()

	return process, exitChannel, nil
}
