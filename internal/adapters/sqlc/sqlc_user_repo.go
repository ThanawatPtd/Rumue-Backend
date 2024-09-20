package sqlc

import (
	"context"
	"errors"

	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SqlcUserRepository struct {
	Queries *dbmodel.Queries
	DB *pgxpool.Pool
}

func NewSqlcUserRepository(db *pgxpool.Pool) repositories.UserRepository {
	return &SqlcUserRepository{
		Queries: dbmodel.New(db),
		DB: db,
	}
}

func (r *SqlcUserRepository) Save(c *context.Context, user *dbmodel.User) (*dbmodel.User, error) {
	seletedUser, err := r.Queries.CreateUser(*c, dbmodel.CreateUserParams{
		Email: user.Email,
		Fname: user.Fname,
		Lname: user.Lname,
		Password: user.Password,
		PhoneNumber: user.PhoneNumber,
		Address: user.Address,
	})

	if err != nil {
		print(err.Error())
		return nil, errors.New("Creating user error.")
	}

	return &dbmodel.User{
		ID: seletedUser.ID,
		Email: seletedUser.Email,
		Fname: seletedUser.Fname,
		CreatedAt: seletedUser.CreatedAt,
		UpdatedAt: seletedUser.UpdatedAt,
	},
	nil
}