package gofu

import (
	"fmt"
	"os"
)

type Config struct {
	DaemonTarget string
}

func NewConfig(grpcDaemonTarget string) *Config {
	return &Config{
		DaemonTarget: grpcDaemonTarget,
	}
}

func DefaultConfig() (*Config, error) {
	directories := NewDirectories()
	bytes, err := os.ReadFile(fmt.Sprintf("%s/%s", directories.ApplicationDirectory, DaemonTargetFile))
	if err != nil {
		return nil, err
	}
	return NewConfig(string(bytes)), nil
}
