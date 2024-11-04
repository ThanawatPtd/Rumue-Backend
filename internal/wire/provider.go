package wire

import (
	"github.com/ThanawatPtd/SAProject/config"
	_ "github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	_ "github.com/ThanawatPtd/SAProject/internal/adapters/mongo"
	"github.com/ThanawatPtd/SAProject/internal/adapters/psql"
	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/mongoDB"
	"github.com/google/wire"
)

var InfraSet = wire.NewSet(
	db.ProvidePgxPool,
	mongodb.GetClient,
)

var ConfigSet = wire.NewSet(
	config.ProvideConfig,
)

var ServiceSet = wire.NewSet(
	usecases.ProvideUserService,
	usecases.ProvideVehicleService,
	usecases.ProvideEmployeeService,
	usecases.ProvideTransactionService,
	usecases.ProvideAuthService,
	usecases.ProvideInsuranceService,
	usecases.ProvideMileService,
	usecases.ProvidePriorityService,
	usecases.ProvideEmailService,
)

var RepositorySet = wire.NewSet(
	psql.ProvidePostgresUserRepository,
	psql.ProvidePostgresEmployeeRepository,
	psql.ProvidePostgresVehicleRepository,
	psql.ProvidePostgresVehicleOwnerRepository,
	psql.ProvidePostgresTransactionRepository,
	psql.ProvidePostgresPriorityRepository,
	psql.ProvidePostgresMileRepository,
	psql.ProvideInsuranceRepository,
	//mongo.ProvideMongoUserRepository,
)

var HandlerSet = wire.NewSet(
	rest.ProvideUserRestHandler,
	rest.ProvideVehicleHandler,
	rest.ProvideEmployeeRestHandler,
	rest.ProvideTransactionRestHandler,
	rest.ProvideAuthRestHandler,
	rest.ProvideInsuranceRestHandler,
	rest.ProvideMileRestHandler,
	rest.ProvidePriorityRestHandler,
	rest.ProvideEmailRestHandler,
)
