package psql

import (
	"context"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

func ProvidePostgresUserRepository(db *pgxpool.Pool) repositories.UserRepository {
	return &PostgresUserRepository{
		Queries: dbmodel.New(db),
	}
}

func (u *PostgresUserRepository) Save(c *context.Context, user *dbmodel.CreateUserParams) (*dbmodel.CreateUserRow, error) {
	selectedUser, err := u.Queries.CreateUser(*c, *user)

	if err != nil {
		return nil, err 
	}

	return &selectedUser, nil
}

func (u *PostgresUserRepository) ListAll(c *context.Context) (*[]dbmodel.GetAllUsersRow, error) {
	selectedUsers, err := u.Queries.GetAllUsers(*c)

	if err != nil {
		return nil, err 
	}

	return &selectedUsers, nil
}

func (u *PostgresUserRepository) GetByEmail(c *context.Context, email *string) (*dbmodel.GetUserByEmailRow, error) {
	selectedUser, err := u.Queries.GetUserByEmail(*c, *email)

	if err != nil {
		return nil, err 
	}

	return &selectedUser, nil 
}

func (u *PostgresUserRepository) GetByID(c *context.Context, id *pgtype.UUID) (*dbmodel.GetUserByIDRow, error) {
	selectedUser, err := u.Queries.GetUserByID(*c, *id)
	
	if err != nil {
		return nil, err 
	}

	return &selectedUser, nil
}

func (u *PostgresUserRepository) Update(c *context.Context, user *dbmodel.UpdateUserParams) (*dbmodel.UpdateUserRow, error) {
	selectedUser, err := u.Queries.UpdateUser(*c, *user)

	if err != nil {
		return nil, err
	}
	return &selectedUser, nil
}

func (u *PostgresUserRepository) Delete(c *context.Context, id *pgtype.UUID) (error) {
	if err := u.Queries.DeleteUser(*c, *id); err != nil {
	return err
	}
	return nil
}
