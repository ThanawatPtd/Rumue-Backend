package wire

import (
	"github.com/google/wire"
)

func InitializeHandler() *rest.Handler {
	wire.Build(
		ConfigSet,
		InfraSet,
		ServiceSet,
		RepositorySet,
		ProvideContext,
		HandlerSet,
		rest.ProvideHandler,
	)

	return &rest.Handler{}
}
