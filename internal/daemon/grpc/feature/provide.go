package grpcfeature

import (
	"os"

	"github.com/olexnzarov/gofu/internal/daemon/grpc"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"go.uber.org/fx"
)

func announce(lc fx.Lifecycle, directories gofu.Directories, server *grpc.Server) {
	lc.Append(fx.StartHook(func() error {
		return os.WriteFile(
			directories.DaemonTargetFile(),
			[]byte(server.Target()),
			0644,
		)
	}))
	lc.Append(fx.StopHook(func() error {
		return os.Remove(directories.DaemonTargetFile())
	}))
}

func runServer(lc fx.Lifecycle, directories gofu.Directories, server *grpc.Server) {
	directories.CreateAll()
	lc.Append(fx.StartHook(server.Start))
	lc.Append(fx.StopHook(server.Stop))
}

func Provide() fx.Option {
	return fx.Options(
		fx.Provide(
			grpc.DefaultServerConfig,
			grpc.NewServer,
		),
		fx.Invoke(announce, runServer),
	)
}
