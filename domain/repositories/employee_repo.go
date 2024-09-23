package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
)

type EmployeeRepository interface {
	ListAll(c *context.Context) (*[]dbmodel.Employee, error)
	Save(c *context.Context, employee *dbmodel.CreateEmployeeParams) (*dbmodel.CreateEmployeeRow, error)
	GetByID(c *context.Context, id *pgtype.UUID) (*dbmodel.Employee, error)
	Update(c *context.Context, employee *dbmodel.UpdateEmployeeParams) (*dbmodel.UpdateEmployeeRow, error)
	Delete(c *context.Context, id *pgtype.UUID) (error)
}