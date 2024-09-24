package psql

import (
	"context"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresEmployeeManagementRepository struct {
	Queries *dbmodel.Queries
}

func ProvidePostgresEmployeeManagementRepository(db *pgxpool.Pool) repositories.EmployeeManagementRepository {
	return &PostgresEmployeeManagementRepository{
		Queries: dbmodel.New(db),
	}
}

func (a *PostgresEmployeeManagementRepository) ListAll(c *context.Context) (*[]dbmodel.EmployeeManagement, error) {
	selectedManagement, err := a.Queries.GetAllEmployeeManagement(*c)

	if err != nil {
		return nil, err
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) Save(c *context.Context, management *dbmodel.CreateEmployeeManagementParams) (*dbmodel.CreateEmployeeManagementRow, error) {
	selectedManagement, err := a.Queries.CreateEmployeeManagement(*c, *management)

	if err != nil {
		return nil, err
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) GetByEmployeeID(c *context.Context, id *pgtype.UUID) (*[]dbmodel.EmployeeManagement, error) {
	selectedManagement, err := a.Queries.GetEmployeeManagementsByEmployeeID(*c, *id)

	if err != nil {
		return nil, err
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) GetByAdminID(c *context.Context, id *pgtype.UUID) (*[]dbmodel.EmployeeManagement, error) {
	selectedManagement, err := a.Queries.GetAllEmployeeManagementsByAdminID(*c, *id)

	if err != nil {
		return nil, err
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) Update(c *context.Context, management *dbmodel.UpdateEmployeeManagementParams) (*dbmodel.UpdateEmployeeManagementRow, error) {
	selectedManagement, err := a.Queries.UpdateEmployeeManagement(*c, * management)

	if err != nil {
		return nil, err
	}
	return &selectedManagement, nil
}

func (a *PostgresEmployeeManagementRepository) Delete(c *context.Context, management *dbmodel.DeleteEmployeeManagementParams) (error) {
	if err := a.Queries.DeleteEmployeeManagement(*c, *management); err != nil {
		return err
	}
	return nil
}