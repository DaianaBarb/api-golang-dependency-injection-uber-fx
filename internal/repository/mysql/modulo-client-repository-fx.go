package mysql

import "go.uber.org/fx"

var Module = fx.Module("db",
	fx.Provide(
		ConnectDB,
		fx.Annotate(
			NewClient,
			fx.As(new(IClientRepository)),
		),
	),
)
