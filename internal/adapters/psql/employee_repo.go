package psql

import (
	"context"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresEmployeeRepository struct {
	Queries *dbmodel.Queries
}

func ProvidePostgresEmployeeRepository(db *pgxpool.Pool) repositories.EmployeeRepository {
	return &PostgresEmployeeRepository{
		Queries: dbmodel.New(db),
	}
}

func (e *PostgresEmployeeRepository) Save(c *context.Context, employee *dbmodel.CreateEmployeeParams) (*dbmodel.CreateEmployeeRow, error) {
	selectedEmployee, err := e.Queries.CreateEmployee(*c, *employee)

	if err != nil {
		return nil, err
	}
	return &selectedEmployee, nil
}

func (e *PostgresEmployeeRepository) ListAll(c *context.Context) (*[]dbmodel.Employee, error) {
	selectedEmployees, err := e.Queries.GetAllEmployees(*c)

	if  err != nil {
		return nil, err
	}
	return &selectedEmployees, nil
}

func (e *PostgresEmployeeRepository) GetByID(c *context.Context, id *pgtype.UUID) (*dbmodel.Employee, error) {
	selectedEmployee, err := e.Queries.GetEmployeeByID(*c, *id)

	if err != nil {
		return nil, err
	}
	return &selectedEmployee, nil
}

func (e *PostgresEmployeeRepository) Update(c *context.Context, employee *dbmodel.UpdateEmployeeParams) (*dbmodel.UpdateEmployeeRow, error) {
	selectedEmployee, err := e.Queries.UpdateEmployee(*c, *employee)

	if err != nil {
		return nil, err
	}
	return &selectedEmployee, nil
}

func (e *PostgresEmployeeRepository) Delete(c *context.Context, id *pgtype.UUID) (error) {
	if err := e.Queries.DeleteEmployee(*c, *id); err != nil {
		return err
	}
	return nil
}