package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type MileRepository interface{
	Save(c context.Context, mile *entities.Mile) (*entities.Mile, error)
	GetByID(c context.Context, id string) (*entities.Mile, error)
}