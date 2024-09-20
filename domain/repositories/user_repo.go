package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
)

type UserRepository interface {
	// ListAll(c *context.Context) ([]responses.UserResponse, error)
	Save(c *context.Context, user *dbmodel.User) (*dbmodel.User, error)
}