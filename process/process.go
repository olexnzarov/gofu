package process

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"go.uber.org/multierr"
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
	isAwaited bool
}

func NewOutOptions(outDirectory string, id string) OutOptions {
	return OutOptions{
		Stdin:  os.DevNull,
		Stdout: fmt.Sprintf("%s/%s-out.log", outDirectory, id),
		Stderr: fmt.Sprintf("%s/%s-err.log", outDirectory, id),
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
		stdin:   stdin,
		stdout:  stdout,
		stderr:  stderr,
		options: options,
	}, nil
}

func (p *Process) Inner() *os.Process {
	return p.inner
}

func (p *Process) Options() *StartOptions {
	return p.options
}

func (p *Process) IsAwaited() bool {
	return p.isAwaited
}

func (p *Process) Close() error {
	return multierr.Append(
		p.inner.Kill(),
		p.release(),
	)
}

func (p *Process) Wait() (*os.ProcessState, error) {
	p.isAwaited = true
	return p.inner.Wait()
}

// Release closes all file descriptors and releases the resources associated with the process.
func (p *Process) release() error {
	ioerror := multierr.Combine(
		p.stdin.Close(),
		p.stderr.Close(),
		p.stdout.Close(),
	)

	if p.inner != nil {
		ioerror = multierr.Append(ioerror, p.inner.Release())
	}

	return ioerror
}

// Start starts a process with given options.
func Start(options StartOptions) (*Process, error) {
	bin, err := exec.LookPath(options.Command)
	if err == nil && bin == "" {
		return nil, err
	}

	// Combine environment variables
	env := os.Environ()
	if options.Environment != nil {
		optEnv := make([]string, 0, len(options.Environment))
		for key, value := range options.Environment {
			optEnv = append(optEnv, fmt.Sprintf("%s=%s", key, value))
		}
		env = append(env, optEnv...)
	}

	if options.Arguments == nil {
		options.Arguments = []string{}
	}

	process, err := prepareProcess(&options)
	if err != nil {
		return nil, err
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

	// argv should start with the program name
	program := strings.Fields(options.Command)[0]
	arguments := append([]string{program}, options.Arguments...)

	p, err := os.StartProcess(bin, arguments, &attributes)
	if err != nil {
		return nil, err
	}
	process.inner = p

	return process, nil
}
