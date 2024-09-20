package wire

import (
	_ "github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
	"github.com/google/wire"
)

// var InfraSet = wire.NewSet(
// 	db.ProvidePgxPool,
// )
//
// var ConfigSet = wire.NewSet(
// 	config.ProvideConfig,
// )
//
// var ServiceSet = wire.NewSet(
// 	usecases.ProvideUserService,
// )
//
// var RepositorySet = wire.NewSet(
// 	psql.ProvideUserPsqlRepository,
// )

var HandlerSet = wire.NewSet(
	rest.ProvideUserRestHandler,
)
