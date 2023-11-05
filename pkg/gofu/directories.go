package gofu

import (
	"fmt"
	"os"

	"go.uber.org/multierr"
)

type Directories struct {
	HomeDirectory        string
	ApplicationDirectory string
	LogDirectory         string
	DataDirectory        string
}

func NewDirectories() Directories {
	homeDirectory, _ := os.UserHomeDir()
	return Directories{
		HomeDirectory:        homeDirectory,
		ApplicationDirectory: fmt.Sprintf("%s/.%s", homeDirectory, ApplicationName),
		LogDirectory:         fmt.Sprintf("%s/.%s/logs", homeDirectory, ApplicationName),
		DataDirectory:        fmt.Sprintf("%s/.%s/data", homeDirectory, ApplicationName),
	}
}

// CreateAll creates all necessary directories.
func (d *Directories) CreateAll() error {
	return multierr.Combine(
		os.MkdirAll(d.LogDirectory, 0755),
		os.MkdirAll(d.DataDirectory, 0755),
	)
}

// RemoveApplicationDirectory removes the application directory.
func (d *Directories) RemoveApplicationDirectory() error {
	return os.RemoveAll(d.ApplicationDirectory)
}

// DaemonTargetFile returns a path to a daemon target file.
func (d *Directories) DaemonTargetFile() string {
	return fmt.Sprintf("%s/%s", d.ApplicationDirectory, DaemonTargetFile)
}
