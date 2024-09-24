package psql

import (
	"context"
	"errors"

	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresEmployeeManagementRepository struct {
	Queries *dbmodel.Queries
	DB *pgxpool.Pool
}

func ProvidePostgresEmployeeManagementRepository(db *pgxpool.Pool) repositories.EmployeeManagementRepository {
	return &PostgresEmployeeManagementRepository{
		Queries: dbmodel.New(db),
		DB: db,
	}
}

func (a *PostgresEmployeeManagementRepository) ListAll(c *context.Context) (*[]dbmodel.EmployeeManagement, error) {
	selectedManagement, err := a.Queries.GetAllEmployeeManagement(*c)

	if err != nil {
		return nil, errors.New("listing all employee management error")
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) Save(c *context.Context, management *dbmodel.CreateEmployeeManagementParams) (*dbmodel.CreateEmployeeManagementRow, error) {
	selectedManagement, err := a.Queries.CreateEmployeeManagement(*c, *management)

	if err != nil {
		return nil, errors.New("creating employee management error")
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) GetByEmployeeID(c *context.Context, id *pgtype.UUID) (*[]dbmodel.EmployeeManagement, error) {
	selectedManagement, err := a.Queries.GetEmployeeManagementsByEmployeeID(*c, *id)

	if err != nil {
		return nil, errors.New("getting employee management error")
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) GetByAdminID(c *context.Context, id *pgtype.UUID) (*[]dbmodel.EmployeeManagement, error) {
	selectedManagement, err := a.Queries.GetAllEmployeeManagementsByAdminID(*c, *id)

	if err != nil {
		return nil, errors.New("getting employee management error")
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) Update(c *context.Context, management *dbmodel.UpdateEmployeeManagementParams) (*dbmodel.UpdateEmployeeManagementRow, error) {
	selectedManagement, err := a.Queries.UpdateEmployeeManagement(*c, * management)

	if err != nil {
		return nil, errors.New("updating employee management error")
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) Delete(c *context.Context, management *dbmodel.DeleteEmployeeManagementParams) (error) {
	if err := a.Queries.DeleteEmployeeManagement(*c, *management); err != nil {
		return errors.New("deleting admin error")
	}
	return nil
}