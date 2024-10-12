package wire

import (
	"github.com/ThanawatPtd/SAProject/config"
	_ "github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/internal/adapters/psql"
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
	usecases.ProvideVehicleService,
	usecases.ProvideEmployeeService,
	usecases.ProvideInvoiceService,
	usecases.ProvideTransactionService,
	usecases.ProvideAuthService,
)

var RepositorySet = wire.NewSet(
	psql.ProvidePostgresUserRepository,
	psql.ProvidePostgresEmployeeRepository,
	psql.ProvidePostgresVehicleRepository,
	psql.ProvidePostgresVehicleOwnerRepository,
	psql.ProvidePostgresInvoiceRepository,
	psql.ProvidePostgresTransactionRepository,
)

var HandlerSet = wire.NewSet(
	rest.ProvideUserRestHandler,
	rest.ProvideVehicleHandler,
	rest.ProvideEmployeeRestHandler,
	rest.ProvideInvoiceRestHandler,
	rest.ProvideTransactionRestHandler,
	rest.ProvideAuthRestHandler,
)
