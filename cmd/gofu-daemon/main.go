package main

import (
	"flag"

	"github.com/olexnzarov/gofu/internal/grpc_server"
	"github.com/olexnzarov/gofu/internal/grpc_server/process_manager_server"
	"github.com/olexnzarov/gofu/internal/process_registry"
	"github.com/olexnzarov/gofu/internal/system_directory"
	"github.com/olexnzarov/gofu/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var (
	FX_NATIVE_LOGGER = flag.Bool("fx-native-logger", false, "use Fx's native logger for its internal logging")
)

func withLogger() fx.Option {
	// By default, the general zap logger will be used.
	// The native logger can be a bit clearer when developing locally.
	if *FX_NATIVE_LOGGER {
		return fx.Supply()
	}

	return fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	})
}

func provideDirectories() (*system_directory.Config, error) {
	config := system_directory.NewConfig("gofu")
	return config, config.EnsureDirectories()
}

func runServer(lc fx.Lifecycle, log *zap.Logger, server *grpc_server.Server) {
	lc.Append(fx.StartHook(server.Start))
	lc.Append(fx.StopHook(server.Stop))
	lc.Append(fx.StopHook(log.Sync))
}

func main() {
	flag.Parse()

	fx.New(
		// For the sake of tracking what dependencies are provided and required,
		// keep the comments with the actual constructors and return types up-to-date.
		fx.Provide(
			logger.New,            // provides *zap.Logger
			provideDirectories,    // provides *system_directory.Config
			grpc_server.NewConfig, // provides *grpc_server.Config

			// provides *process_registry.ProcessRegistry
			// requires *zap.Logger
			process_registry.New,

			// provides *process_manager_server.ProcessManagerServer
			// requires *zap.Logger, *system_directory.Config, *process_registry.ProcessRegistry
			process_manager_server.New,

			// provides *grpc_server.Server
			// requires *grpc_server.Config, *zap.Logger, *process_manager_server.ProcessManagerServer
			grpc_server.New,
		),
		withLogger(),
		fx.Invoke(runServer),
	).Run()
}
