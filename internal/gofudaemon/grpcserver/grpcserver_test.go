package grpcserver

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olexnzarov/gofu/internal/gofudaemon/grpcserver/processmanagerserver"
	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
	"github.com/olexnzarov/gofu/internal/gofudaemon/procservice"
	"github.com/olexnzarov/gofu/internal/gofudaemon/sqlitestorage"
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
	storage := sqlitestorage.New(log, db)
	manager := procmanager.New(log, directories)
	service := procservice.New(log, manager, storage)

	pms := processmanagerserver.New(log, directories, service)

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
