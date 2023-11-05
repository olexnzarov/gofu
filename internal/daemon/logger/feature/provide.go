package loggerfeature

import (
	"github.com/olexnzarov/gofu/internal/daemon/logger"
	"go.uber.org/fx"
)

func sync(lc fx.Lifecycle, log logger.Logger) {
	lc.Append(fx.StopHook(log.Sync))
}

func Provide() fx.Option {
	return fx.Options(
		fx.Provide(logger.New),
		fx.Invoke(sync),
	)
}
