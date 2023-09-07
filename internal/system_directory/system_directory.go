package system_directory

import (
	"fmt"
	"os"
)

type Config struct {
	HomeDirectory        string
	ApplicationDirectory string
	LogDirectory         string
}

func NewConfig(appName string) *Config {
	homeDirectory, _ := os.UserHomeDir()
	return &Config{
		HomeDirectory:        homeDirectory,
		ApplicationDirectory: fmt.Sprintf("%s/.%s", homeDirectory, appName),
		LogDirectory:         fmt.Sprintf("%s/.%s/logs", homeDirectory, appName),
	}
}

// EnsureDirectories creates all necessary directories for the given config.
func (c *Config) EnsureDirectories() error {
	return os.MkdirAll(c.LogDirectory, 0755)
}

// CleanupDirectories removes the application directory.
func (c *Config) CleanupDirectories() error {
	return os.RemoveAll(c.ApplicationDirectory)
}
