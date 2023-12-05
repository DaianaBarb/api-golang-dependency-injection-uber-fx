package rest

import (
	"go.uber.org/fx"
)

var Module = fx.Module("handler", fx.Provide(
	fx.Annotate(
		NewCLientHandler,
		fx.As(new(IClientHandler)),
	),
))
