package psql

import (
	"context"
	"errors"

	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresEmployeeRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

func ProvidePostgresEmployeeRepository(db *pgxpool.Pool) repositories.EmployeeRepository {
	return &PostgresEmployeeRepository{
		Queries: dbmodel.New(db),
		DB:      db,
	}
}

func (e *PostgresEmployeeRepository) Save(c *context.Context, employee *dbmodel.CreateEmployeeParams) (*dbmodel.CreateEmployeeRow, error) {
	selectedEmployee, err := e.Queries.CreateEmployee(*c, *employee)

	if err != nil {
		return nil, errors.New("creating employee error")
	}
	return &selectedEmployee, nil
}

func (e *PostgresEmployeeRepository) ListAll(c *context.Context) (*[]dbmodel.Employee, error) {
	selectedEmployees, err := e.Queries.GetAllEmployees(*c)

	if  err != nil {
		return nil, errors.New("listing All employees error")	
	}
	return &selectedEmployees, nil
}

func (e *PostgresEmployeeRepository) GetByID(c *context.Context, id *pgtype.UUID) (*dbmodel.Employee, error) {
	selectedEmployee, err := e.Queries.GetEmployeeByID(*c, *id)

	if err != nil {
		return nil, errors.New("getting employee by id error")
	}
	return &selectedEmployee, nil
}

func (e *PostgresEmployeeRepository) Update(c *context.Context, employee *dbmodel.UpdateEmployeeParams) (*dbmodel.UpdateEmployeeRow, error) {
	selectedEmployee, err := e.Queries.UpdateEmployee(*c, *employee)

	if err != nil {
		return nil, errors.New("updating employee error")
	}
	return &selectedEmployee, nil
}

func (e *PostgresEmployeeRepository) Delete(c *context.Context, id *pgtype.UUID) (error) {
	if err := e.Queries.DeleteEmployee(*c, *id); err != nil {
		return errors.New("deleting employee error")
	}
	return nil
}