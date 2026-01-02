package tracer

import (
	"github.com/not-for-prod/observer/tracer"
	"go.uber.org/fx"
)

func NewProviderFx(options ...tracer.Option) fx.Option {
	return fx.Invoke(
		func(LC fx.Lifecycle) *tracer.Provider {
			provider := tracer.NewProvider(options...)
			LC.Append(fx.Hook{
				OnStart: provider.Start,
				OnStop:  provider.Stop,
			})

			return provider
		},
	)
}
