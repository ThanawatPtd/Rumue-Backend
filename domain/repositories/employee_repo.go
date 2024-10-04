package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

type EmployeeRepository interface {
	ListAll(c *context.Context) (*[]entities.Employee, error)
	Save(c *context.Context, employee *entities.Employee) (*entities.Employee, error)
	GetByID(c *context.Context, id *pgtype.UUID) (*entities.Employee, error)
	Update(c *context.Context, employee *entities.Employee) (*entities.Employee, error)
	Delete(c *context.Context, id *pgtype.UUID) (error)
}