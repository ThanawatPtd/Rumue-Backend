package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	Queries *dbmodel.Queries
	DB      *pgxpool.Pool
}

func ProvidePostgresUserRepository(db *pgxpool.Pool) repositories.UserRepository {
	return &PostgresUserRepository{
		Queries: dbmodel.New(db),
		DB:      db,
	}
}

func (u *PostgresUserRepository) Save(c context.Context, user *entities.User) error {
	paramsUser := dbmodel.CreateUserParams{}
	if err := utils.MappingParser(user, &paramsUser); err != nil {
		return err
	}
	return u.Queries.CreateUser(c, paramsUser)
}

func (u *PostgresUserRepository) ListAll(c context.Context) ([]entities.User, error) {
	selectedUsers, err := u.Queries.GetAllUsers(c)

	if err != nil {
		return nil, err
	}

	if selectedUsers == nil {
		return []entities.User{}, nil
	}

	var users []entities.User
	for _, value := range selectedUsers {
		user := entities.User{}
		if err := utils.MappingParser(&value, &user); err != nil {
			return nil, err
		}
		users = append(users, user)

	}

	return users, nil
}

// Delete implements repositories.UserRepository.
func (u *PostgresUserRepository) Delete(c context.Context, id string) error {
	ID := convert.StringToUUID(id)
	err := u.Queries.DeleteUser(c, ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *PostgresUserRepository) GetIDPasswordByEmail(c context.Context, email string) (*entities.User, error) {
	idPassword, err := u.Queries.GetUserIDPasswordByEmail(c, email)
	if errors.Is(err, sql.ErrNoRows) {
        return nil, nil
    }
	if err != nil {
		return nil, err
	}

	user := &entities.User{}
	if err := utils.MappingParser(&idPassword, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetByID implements repositories.UserRepository.
func (u *PostgresUserRepository) GetByID(c context.Context, id string) (*entities.User, error) {
	ID := convert.StringToUUID(id)
	getUser, err := u.Queries.GetUserByID(c, ID)
	if errors.Is(err, sql.ErrNoRows) {
        return nil, nil
    }

	if err != nil {
		return nil, err
	}

	user := &entities.User{}
	if err := utils.MappingParser(&getUser, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Update implements repositories.UserRepository.
func (u *PostgresUserRepository) Update(c context.Context, user *entities.User) (*entities.User, error) {
	paramsUser := &dbmodel.UpdateUserParams{}
	if err := utils.MappingParser(user, paramsUser); err != nil {
		return nil, err
	}

	updateUser, err := u.Queries.UpdateUser(c, *paramsUser)
	if err != nil {
		return nil, err
	}

	user = &entities.User{}
	if err := utils.MappingParser(&updateUser, user); err != nil {
		return nil, err
	}

	return user, nil
}
