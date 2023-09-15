package grpc_server

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/grpc_server/process_manager_server"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/process_manager"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/process_storage"
	"github.com/olexnzarov/gofu/internal/logger"
	"github.com/olexnzarov/gofu/pkg/gofu"
)

func TestServer(t *testing.T) {
	directories := gofu.NewDirectories()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create a database instace: %s", err)
		return
	}

	config := NewConfig()
	log, _ := logger.New()
	storage := process_storage.New(log, db)
	processManager := process_manager.New(log, directories, storage)

	pms := process_manager_server.New(log, directories, processManager)

	server := New(log, config, pms)
	t.Cleanup(server.inner.Stop)

	err = server.Start()
	if err != nil {
		t.Fatalf("Failed to start the server: %s", err)
		return
	}

	err = server.Stop()
	if err != nil {
		t.Fatalf("Failed to stop the server: %s", err)
	}
}
