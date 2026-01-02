package logger

import (
	"context"

	"go.uber.org/fx"

	"github.com/not-for-prod/observer/logger"
)

func NewLoggerFx(
	log logger.Logger,
) fx.Option {
	return fx.Invoke(
		func(lc fx.Lifecycle) {
			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						logger.SetLogger(log)
						return nil
					},
					OnStop: logger.Stop,
				},
			)
		},
	)
}
