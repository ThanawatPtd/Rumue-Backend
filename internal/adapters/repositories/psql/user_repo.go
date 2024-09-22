package psql

import (
	"context"
	"errors"

	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

func ProvideUserRepository(db *pgxpool.Pool) repositories.UserRepository {
	return &UserRepository{
		Queries: dbmodel.New(db),
		DB:      db,
	}
}

func (u *UserRepository) Save(c context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error) {
	newUser, err := u.Queries.CreateUser(c, *user)

	if err != nil {
		print(err.Error())
		return nil, errors.New("Creating user error.")
	}

	return &newUser,nil
}

// func (u *UserRepository) ListAll(c context.Context) (*[]dbmodel.GetAllUsersRow) {
// 	users, err := u.Queries.GetAllEmployees(c)

// 	if err != nil {

// 	}
// }
