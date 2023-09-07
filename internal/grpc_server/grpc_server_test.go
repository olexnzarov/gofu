package grpc_server

import (
	"testing"

	"github.com/alexnzarov/gofu/internal/grpc_server/process_manager_server"
	"github.com/alexnzarov/gofu/internal/process_registry"
	"github.com/alexnzarov/gofu/internal/system_directory"
	"github.com/alexnzarov/gofu/logger"
)

func TestServer(t *testing.T) {
	config := NewConfig()
	log, _ := logger.New()
	process_registry := process_registry.New(log)

	directories := system_directory.NewConfig("gofu-test")
	t.Cleanup(func() { directories.CleanupDirectories() })

	pms := process_manager_server.New(log, directories, process_registry)

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
