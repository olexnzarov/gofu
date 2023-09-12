package main

import (
	"flag"

	"github.com/olexnzarov/gofu/internal/gofu_daemon"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/grpc_server"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/grpc_server/process_manager_server"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/process_manager"
	"github.com/olexnzarov/gofu/internal/logger"
	"github.com/olexnzarov/gofu/pkg/gofu"
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

// Creates an entrypoint of the application.
func runDaemon(lc fx.Lifecycle, daemon *gofu_daemon.Daemon) {
	lc.Append(fx.StartHook(daemon.Start))
	lc.Append(fx.StopHook(daemon.Stop))
}

func main() {
	flag.Parse()

	fx.New(
		fx.Provide(
			logger.New,
			gofu.NewDirectories,
			process_manager.New,
			process_manager_server.New,
			grpc_server.NewConfig,
			grpc_server.New,
			gofu_daemon.NewDaemon,
		),
		withLogger(),
		fx.Invoke(runDaemon),
	).Run()
}
