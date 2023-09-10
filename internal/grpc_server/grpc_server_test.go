package grpc_server

import (
	"testing"

	"github.com/olexnzarov/gofu/internal/grpc_server/process_manager_server"
	"github.com/olexnzarov/gofu/internal/process_manager"
	"github.com/olexnzarov/gofu/internal/system_directory"
	"github.com/olexnzarov/gofu/logger"
)

func TestServer(t *testing.T) {
	directories := system_directory.NewConfig("gofu-test")
	t.Cleanup(func() { directories.CleanupDirectories() })

	config := NewConfig()
	log, _ := logger.New()
	processManager := process_manager.New(log, directories)

	pms := process_manager_server.New(log, directories, processManager)

	server := New(log, config, pms)
	t.Cleanup(server.inner.Stop)

	err := server.Start()
	if err != nil {
		t.Fatalf("failed to start the server: %s", err.Error())
		return
	}

	err = server.Stop()
	if err != nil {
		t.Fatalf("failed to stop the server: %s", err.Error())
	}
}
