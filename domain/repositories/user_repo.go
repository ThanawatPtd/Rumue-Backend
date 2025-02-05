package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type UserRepository interface {
	ListAll(c context.Context) ([]entities.User, error)
	Save(c context.Context, user *entities.User) error
	GetIDPasswordByEmail(c context.Context, email string) (*entities.User, error)
	GetIDPasswordByID(c context.Context, id string) (*entities.User, error)
	GetByID(c context.Context, id string) (*entities.User, error)
	GetUserProfileByID(c context.Context, id string) (*entities.UserProfile, error)
	Update(c context.Context, user *entities.User) (*entities.User, error)
	UpdatePassword(c context.Context, user *entities.User) error
	Delete(c context.Context, id string) error
}
