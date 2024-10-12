package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type EmployeeRepository interface {
	ListAll(c context.Context) ([]entities.Employee, error)
	Save(c context.Context, employee *entities.Employee) (*entities.Employee, error)
	GetByID(c context.Context, id string) (*entities.Employee, error)
	Update(c context.Context, employee *entities.Employee) (*entities.Employee, error)
	Delete(c context.Context, id string) (error)
}