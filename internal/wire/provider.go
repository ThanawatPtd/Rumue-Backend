package wire

import (
	"github.com/ThanawatPtd/SAProject/config"
	_ "github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/internal/adapters/repositories/psql"
	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db"
	"github.com/google/wire"
)

var InfraSet = wire.NewSet(
	db.ProvidePgxPool,
)

var ConfigSet = wire.NewSet(
	config.ProvideConfig,
)

var ServiceSet = wire.NewSet(
	usecases.ProvideUserService,
)

var RepositorySet = wire.NewSet(
	psql.ProvidePostgresUserRepository,
)

var HandlerSet = wire.NewSet(
	rest.ProvideUserRestHandler,
)
