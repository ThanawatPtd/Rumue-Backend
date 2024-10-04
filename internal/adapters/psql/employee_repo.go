package psql

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/ThanawatPtd/SAProject/utils"
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

func (e *PostgresEmployeeRepository) Save(c *context.Context, employee *entities.Employee) (*entities.Employee, error) {
	paramsEmployee := &dbmodel.CreateEmployeeParams{}
	if err := utils.MappingParser(employee, paramsEmployee); err != nil{
		return nil, err
	}

	selectedEmployee, err := e.Queries.CreateEmployee(*c, *paramsEmployee)
	if err != nil {
		return nil, err
	}

	employee = &entities.Employee{}
	if err := utils.MappingParser(&selectedEmployee, employee); err != nil{
		return nil, err
	}	
	return employee, nil
}

func (e *PostgresEmployeeRepository) ListAll(c *context.Context) (*[]entities.Employee, error) {
	selectedEmployees, err := e.Queries.GetAllEmployees(*c)
	if  err != nil {
		return nil, err
	}

	employees := []entities.Employee{}
	for _, value := range selectedEmployees{
		employee := &entities.Employee{}	
		if err := utils.MappingParser(&value, employee); err != nil {
			return nil, err
		}
		employees = append(employees, *employee) 
	}
	return &employees, nil
}

func (e *PostgresEmployeeRepository) GetByID(c *context.Context, id *pgtype.UUID) (*entities.Employee, error) {
	selectedEmployee, err := e.Queries.GetEmployeeByID(*c, *id)
	if err != nil {
		return nil, err
	}

	employee := entities.Employee{}
	if err := utils.MappingParser(&selectedEmployee, &employee); err != nil {
		return nil, err
	}
	return &employee, nil
}

func (e *PostgresEmployeeRepository) Update(c *context.Context, employee *entities.Employee) (*entities.Employee, error) {
	paramsEmployee := dbmodel.UpdateEmployeeParams{}
	if err := utils.MappingParser(employee, &paramsEmployee); err != nil {
		return nil, err
	}
	selectedEmployee, err := e.Queries.UpdateEmployee(*c, paramsEmployee)
	if err != nil {
		return nil, err
	}

	employee = &entities.Employee{}
	if err := utils.MappingParser(&selectedEmployee, employee); err != nil {
		return nil, err
	}
	return employee, nil
}

func (e *PostgresEmployeeRepository) Delete(c *context.Context, id *pgtype.UUID) (error) {
	if err := e.Queries.DeleteEmployee(*c, *id); err != nil {
		return err
	}
	return nil
}