package main

import (
	grpcfeature "github.com/olexnzarov/gofu/internal/daemon/grpc/feature"
	loggerfeature "github.com/olexnzarov/gofu/internal/daemon/logger/feature"
	processmanagerfeature "github.com/olexnzarov/gofu/internal/daemon/processmanager/feature"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(gofu.NewDirectories),
		loggerfeature.Provide(),
		processmanagerfeature.Provide(),
		grpcfeature.Provide(),
	).Run()
}
