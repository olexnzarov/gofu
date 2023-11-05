package grpc

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olexnzarov/gofu/internal/daemon/logger"
	processmanagerserver "github.com/olexnzarov/gofu/internal/daemon/processmanager/server"
	processmanagerservice "github.com/olexnzarov/gofu/internal/daemon/processmanager/service"
	processmanagerstorage "github.com/olexnzarov/gofu/internal/daemon/processmanager/storage"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/gofu"
)

func TestServer(t *testing.T) {
	directories := gofu.NewDirectories()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create a database instace: %s", err)
		return
	}

	config := DefaultServerConfig()
	log, _ := logger.New()
	storage := processmanagerstorage.NewSQL(log, &processmanagerstorage.SQLDB{DB: db})
	service := processmanagerservice.New(log, directories, storage)

	pms := processmanagerserver.New(log, directories, service)

	server := NewServer(log, config)
	pb.RegisterProcessManagerServer(server.Raw(), pms)
	t.Cleanup(server.server.Stop)

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
