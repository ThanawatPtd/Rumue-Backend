package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
)

type EmployeeManagementRepository interface {
	ListAll(c *context.Context) (*[]dbmodel.EmployeeManagement, error)
	Save(c *context.Context, management *dbmodel.CreateEmployeeManagementParams) (*dbmodel.CreateEmployeeManagementRow, error)
	GetByEmployeeID(c *context.Context, id *pgtype.UUID) (*[]dbmodel.EmployeeManagement, error)
	GetByAdminID(c *context.Context, id *pgtype.UUID) (*[]dbmodel.EmployeeManagement, error)
	Update(c *context.Context, management *dbmodel.UpdateEmployeeManagementParams) (*dbmodel.UpdateEmployeeManagementRow, error) 
	Delete(c *context.Context, management *dbmodel.DeleteEmployeeManagementParams) (error)
}