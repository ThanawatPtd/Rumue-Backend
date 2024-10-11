package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type UserRepository interface {
	ListAll(c context.Context) (*[]entities.User, error)
	Save(c context.Context, user *entities.User) (*entities.User, error)
	GetByEmail(c context.Context, email string) (*entities.User, error)
	GetByID(c *context.Context, id string) (*entities.User, error)
	Update(c *context.Context, user *entities.User) (*entities.User, error)
	Delete(c context.Context, id string) error
}
