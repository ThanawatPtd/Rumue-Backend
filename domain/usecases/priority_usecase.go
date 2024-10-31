package usecases

import (
	"context"
	"errors"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
)

type PriorityUseCase interface {
	CreatePriorityByID(c context.Context, priority *entities.Priority) (*entities.Priority, error)
	GetPriorityByID(c context.Context, id string) (*entities.Priority, error)
}

type PriorityService struct {
	repo repositories.PriorityRepository
}

func ProvidePriorityService(repo repositories.PriorityRepository) PriorityUseCase {
	return &PriorityService{
		repo: repo,
	}
}

func (p *PriorityService) CreatePriorityByID(c context.Context, priority *entities.Priority) (*entities.Priority, error) {
	selected, err := p.repo.GetByID(c, priority.ID)

	if err != nil {
		return nil, err
	}
	
	if selected != nil {
		return nil, errors.New("this mile is already exists")
	}

	created, err := p.repo.Save(c, priority)

	if err != nil{
		return nil, err
	}

	return created, nil
}


func (p *PriorityService) GetPriorityByID(c context.Context, id string) (*entities.Priority, error) {
	selected, err := p.repo.GetByID(c, id)

	if err != nil{
		return nil, err
	}

	return selected, err
}