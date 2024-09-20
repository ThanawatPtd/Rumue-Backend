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

func (u *UserRepository) Save(c context.Context, user *dbmodel.CreateUserParams) (*dbmodel.User, error) {
	newUser, err := u.Queries.CreateUser(c, *user)

	if err != nil {
		print(err.Error())
		return nil, errors.New("Creating user error.")
	}

	return &dbmodel.User{
			ID:        newUser.ID,
			Email:     newUser.Email,
			Fname:     newUser.Fname,
			Lname:     newUser.Lname,
			CreatedAt: newUser.CreatedAt,
			UpdatedAt: newUser.UpdatedAt,
		},
		nil
}
