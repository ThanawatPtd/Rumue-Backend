//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
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
