package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository interface {
	ListAll(c *context.Context) (*[]dbmodel.GetAllUsersRow, error)
	Save(c *context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error)
	GetByEmail(c *context.Context, email *string) (*dbmodel.GetUserByEmailRow, error)
	GetByID(c *context.Context, id *pgtype.UUID) (*dbmodel.GetUserByIDRow, error)
	Update(c *context.Context, user *dbmodel.UpdateUserParams) (*dbmodel.UpdateUserRow, error)
	Delete(c *context.Context, id *pgtype.UUID) (error)
}
