package system_directory

import (
	"os"
	"testing"
)

func TestEnsureDirectories(t *testing.T) {
	config := NewConfig("gofu-test")

	err := config.EnsureDirectories()
	if err != nil {
		t.Fatalf("failed to ensure the directories: %s", err.Error())
		return
	}

	directoriesToCheck := []string{config.HomeDirectory, config.ApplicationDirectory, config.LogDirectory}
	for _, directory := range directoriesToCheck {
		_, err := os.Stat(directory)
		if err != nil && os.IsNotExist(err) {
			t.Fatalf("directory '%s' doesn't exist", directory)
		}
	}

	err = config.CleanupDirectories()
	if err != nil {
		t.Fatalf("failed to cleanup the directories: %s", err.Error())
		return
	}

	directoriesToCheck = []string{config.ApplicationDirectory, config.LogDirectory}
	for _, directory := range directoriesToCheck {
		_, err := os.Stat(directory)
		if err == nil || !os.IsNotExist(err) {
			t.Fatalf("directory '%s' exists after the cleanup", directory)
		}
	}
}
