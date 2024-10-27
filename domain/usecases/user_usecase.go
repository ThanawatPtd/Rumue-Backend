package usecases

import (
	"context"

	"github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/exceptions"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	GetUserByID(ctx context.Context, id string) (*entities.User, error)
	DeleteByID(ctx context.Context, id string) error
	GetUsers(ctx context.Context) ([]entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	UpdatePassword(c context.Context, id string, oldPassword string, newPassword string) error
}

type UserService struct {
	userRepo repositories.UserRepository
	config   *config.Config
}

func ProvideUserService(userRepo repositories.UserRepository, config *config.Config) UserUseCase {
	return &UserService{
		userRepo: userRepo,
		config:   config,
	}
}

// DeleteByID implements UserUseCase.
func (u *UserService) DeleteByID(ctx context.Context, id string) error {
	findUser, err := u.GetUserByID(ctx, id)
	if findUser == nil {
		return exceptions.ErrUserNotFound
	}

	if err != nil {
		return err
	}

	err = u.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID implements UserUseCase.
func (u *UserService) GetUserByID(ctx context.Context, id string) (*entities.User, error) {
	getUser, err := u.userRepo.GetByID(ctx, id)
	if getUser == nil {
		return nil, exceptions.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return getUser, nil
}

// GetUsers implements UserUseCase.
func (u *UserService) GetUsers(ctx context.Context) ([]entities.User, error) {
	users, err := u.userRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser implements UserUseCase.
func (u *UserService) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	selectUser, err := u.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	selectUser.ID = user.ID
	if err := utils.MappingParser(user, selectUser); err != nil {
		return nil, err
	}
	return u.userRepo.Update(ctx, selectUser)
}

func (u *UserService) UpdatePassword(ctx context.Context, id string, oldPassword string, newPassword string) error {
	selectUser, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(selectUser.Password), []byte(oldPassword)); err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(selectUser.Password), []byte(newPassword)); err == nil {
		return nil
	}
	if err := utils.ValidatePassword(&utils.RegexPasswordValidator{}, newPassword); err != nil {
		return err
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	selectUser.Password = string(hashPassword)

	_, err = u.userRepo.Update(ctx, selectUser)
	if err != nil {
		return err
	}
	return nil
}
