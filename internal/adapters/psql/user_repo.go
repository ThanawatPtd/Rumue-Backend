package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
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

func (u *PostgresUserRepository) Save(c context.Context, user *entities.User) (*entities.User, error) {
	var createUser dbmodel.CreateUserParams
	createUser = dbmodel.CreateUserParams{
		Email:       user.Email,
		Fname:       user.Fname,
		Lname:       user.Lname,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}

	selectedUser, err := u.Queries.CreateUser(c, createUser)

	if err != nil {
		return nil, err
	}

	var createdUser entities.User
	createdUser = entities.User{
		ID:      convert.UUIDToString(selectedUser.ID),
		Email:       selectedUser.Email,
		Fname:       selectedUser.Fname,
		Lname:       selectedUser.Lname,
		PhoneNumber: selectedUser.PhoneNumber,
		Address:     selectedUser.Address,
	}

	return &createdUser, nil
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
	for _, user := range selectedUsers {
		mapUser := entities.User{
			ID:      convert.UUIDToString(user.ID),
			Email:       user.Email,
			Fname:       user.Fname,
			Lname:       user.Lname,
			Password:    user.Password,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
		}
		users = append(users, mapUser)

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

// GetByEmail implements repositories.UserRepository.
func (u *PostgresUserRepository) GetByEmail(c context.Context, email *string) (*entities.User, error) {
	getUser, err := u.Queries.GetUserByEmail(c, *email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var user entities.User
	user = entities.User{
		ID:      convert.UUIDToString(getUser.ID),
		Email:       getUser.Email,
		Fname:       getUser.Fname,
		Lname:       getUser.Lname,
		Password:    getUser.Password,
		PhoneNumber: getUser.PhoneNumber,
		Address:     getUser.Address,
	}
	return &user, nil
}

// GetByID implements repositories.UserRepository.
func (u *PostgresUserRepository) GetByID(c context.Context, id string) (*entities.User, error) {
	ID := convert.StringToUUID(id)
	getUser, err := u.Queries.GetUserByID(c, ID)

	if err != nil {
		return nil, err
	}
	var user entities.User
	user = entities.User{
		ID:      id,
		Email:       getUser.Email,
		Fname:       getUser.Fname,
		Lname:       getUser.Lname,
		Password:    getUser.Password,
		PhoneNumber: getUser.PhoneNumber,
		Address:     getUser.Address,
	}
	return &user, nil
}

// Update implements repositories.UserRepository.
func (u *PostgresUserRepository) Update(c context.Context, user *entities.User) (*entities.User, error) {
	var updateUserParams dbmodel.UpdateUserParams
	updateUserParams = dbmodel.UpdateUserParams{
		ID:          convert.StringToUUID(user.ID),
		Email:       user.Email,
		Fname:       user.Fname,
		Lname:       user.Lname,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		Address:     user.Address,
	}
	getUser, err := u.Queries.UpdateUser(c, updateUserParams)
	if err != nil {
		return nil, err
	}
	var mapUser entities.User
	mapUser = entities.User{
		ID:      convert.UUIDToString(getUser.ID),
		Email:       getUser.Email,
		Fname:       getUser.Fname,
		Lname:       getUser.Lname,
		Password:    getUser.Password,
		PhoneNumber: getUser.PhoneNumber,
		Address:     getUser.Address,
	}
	return &mapUser, nil
}
