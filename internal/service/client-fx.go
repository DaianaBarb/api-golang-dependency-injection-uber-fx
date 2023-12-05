package service

import (
	"go.uber.org/fx"
)

var Module = fx.Module("service", fx.Provide(
	fx.Annotate(
		NewClientService,
		fx.As(new(IclientService)),
	),
))
