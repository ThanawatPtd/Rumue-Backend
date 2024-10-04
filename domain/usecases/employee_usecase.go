package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	// "github.com/jackc/pgx/v5/pgtype"
)

type EmployeeUseCase interface{
	// ListAll(ctx context.Context) (*[]dbmodel.Employee, error)
	Save(ctx context.Context, employee *entities.Employee) (*entities.Employee, error)
	// GetByID(ctx context.Context, id *pgtype.UUID) (*dbmodel.Employee, error)
	// Update(ctx context.Context, employee *dbmodel.UpdateEmployeeParams) (*dbmodel.UpdateEmployeeRow, error)
	// Delete(ctx context.Context, id *pgtype.UUID) (error)
}

type EmployeeService struct{
	repo repositories.EmployeeRepository
}


func ProvideEmployeeService(repo repositories.EmployeeRepository) EmployeeUseCase{
	return &EmployeeService{
		repo: repo,
	}
}

func (es *EmployeeService) Save(ctx context.Context, employee *entities.Employee) (*entities.Employee, error){
	response, err := es.repo.Save(&ctx, employee)

	if err != nil{
		return nil, err
	}

	return response, nil
}