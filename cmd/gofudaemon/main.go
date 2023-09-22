package main

import (
	"database/sql"
	"flag"

	"github.com/olexnzarov/gofu/internal/gofudaemon"
	"github.com/olexnzarov/gofu/internal/gofudaemon/grpcserver"
	"github.com/olexnzarov/gofu/internal/gofudaemon/grpcserver/processmanagerserver"
	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
	"github.com/olexnzarov/gofu/internal/gofudaemon/procservice"
	"github.com/olexnzarov/gofu/internal/gofudaemon/sqlitestorage"
	"github.com/olexnzarov/gofu/internal/logger"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"go.uber.org/fx"
)

// An entrypoint of the application.
func runDaemon(lc fx.Lifecycle, log logger.Logger, daemon *gofudaemon.Daemon) {
	lc.Append(fx.StartHook(daemon.Start))
	lc.Append(fx.StopHook(daemon.Stop))
	lc.Append(fx.StopHook(log.Sync))
}

func persistentStorage(log logger.Logger, sql *sql.DB) (procservice.Storage, error) {
	storage := sqlitestorage.New(log, sql)
	if err := storage.Initialize(); err != nil {
		return nil, err
	}
	return storage, nil
}

func initialize(service *procservice.Service) {
	go service.StartPersistent()
}

func main() {
	flag.Parse()

	fx.New(
		fx.Provide(
			logger.New,
			gofu.NewDirectories,
			procmanager.New,
			procservice.New,
			processmanagerserver.New,
			persistentStorage,
			grpcserver.NewConfig,
			grpcserver.New,
			gofudaemon.NewDatabase,
			gofudaemon.New,
		),
		fx.Invoke(runDaemon),
		fx.Invoke(initialize),
	).Run()
}
