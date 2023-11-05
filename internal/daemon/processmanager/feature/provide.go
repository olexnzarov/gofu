package processmanagerfeature

import (
	"github.com/olexnzarov/gofu/internal/daemon/grpc"
	"github.com/olexnzarov/gofu/internal/daemon/processmanager"
	processmanagerserver "github.com/olexnzarov/gofu/internal/daemon/processmanager/server"
	processmanagerservice "github.com/olexnzarov/gofu/internal/daemon/processmanager/service"
	processmanagerstorage "github.com/olexnzarov/gofu/internal/daemon/processmanager/storage"
	"github.com/olexnzarov/gofu/pb"
	"go.uber.org/fx"
)

func initializeStorage(storage processmanager.Storage) {
	storage.Initialize()
}

func initializeService(lc fx.Lifecycle, service *processmanagerservice.Service) {
	lc.Append(fx.StartHook(func() {
		go service.StartPersistent()
	}))
}

func registerServer(server *grpc.Server, pms pb.ProcessManagerServer) {
	pb.RegisterProcessManagerServer(server.Raw(), pms)
}

func Provide() fx.Option {
	return fx.Options(
		fx.Provide(
			processmanagerstorage.WithSQLite,
			processmanagerstorage.NewSQL,
			processmanagerservice.New,
			processmanagerserver.New,
		),
		fx.Invoke(initializeStorage, initializeService, registerServer),
	)
}
