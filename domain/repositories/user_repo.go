package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
)

type UserRepository interface {
	// ListAll(c *context.Context) (*[]dbmodel.GetAllUsersRow, error)
	Save(c context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error)
}
