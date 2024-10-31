package usecases

import (
	"context"
	"errors"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
)

type MileUseCase interface {
	GetMileByID(c context.Context, id string) (*entities.Mile, error)
	CreateMileByID(c context.Context, mile *entities.Mile) (*entities.Mile, error)
}

type MileService struct {
	repo repositories.MileRepository
}

func ProvideMileService(repo repositories.MileRepository) MileUseCase {
	return &MileService{
		repo: repo,
	}
}

func (m *MileService) CreateMileByID(c context.Context, mile *entities.Mile) (*entities.Mile, error) {
	selected, err := m.repo.GetByID(c, mile.ID)

	if err != nil {
		return nil, err
	}
	
	if selected != nil {
		return nil, errors.New("this mile is already exists")
	}

	created, err := m.repo.Save(c, mile)

	if err != nil{
		return nil, err
	}

	return created, nil
}

func (m *MileService) GetMileByID(c context.Context, id string) (*entities.Mile, error) {
	selected, err := m.repo.GetByID(c, id)

	if err != nil{
		return nil, err
	}

	return selected, err
}