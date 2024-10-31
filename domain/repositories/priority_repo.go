package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type PriorityRepository interface{
	Save(c context.Context, priority *entities.Priority) (*entities.Priority, error) 
	GetByID(c context.Context, id string) (*entities.Priority, error) 
}