package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/exceptions"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Login(ctx context.Context, user *entities.User) (*entities.User, string, error)
	Register(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserByID(ctx context.Context, id string) (*entities.User, error)
	DeleteByID(ctx context.Context, id string) error
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUsers(ctx context.Context) ([]entities.User, error)
	UpdateUser(c context.Context, id string, user *entities.User) (*entities.User, error)
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

// Register implements UserUseCase.
func (u *UserService) Register(ctx context.Context, user *entities.User) (*entities.User, error) {
	// Find user by email
	getUser, err := u.userRepo.GetByEmail(ctx, &user.Email)

	if err != nil {
		return nil, err
	}

	// Check if user already exist
	if getUser != nil {
		return nil, exceptions.ErrDuplicatedEmail
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	return u.userRepo.Save(ctx, user)
}

// Login implements UserUseCase.
func (u *UserService) Login(ctx context.Context, user *entities.User) (*entities.User, string, error) {
	// Find user by email
	getUser, err := u.userRepo.GetByEmail(ctx, &user.Email)

	if err != nil {
		return nil, "", err
	}

	// Check if user exist
	if getUser == nil {
		return nil, "", exceptions.ErrLoginFailed
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(user.Password)) != nil {
		fmt.Println("fail password not match")
		return nil, "", exceptions.ErrLoginFailed
	}

	// Generate JWT token
	expireAt := time.Now().Add(time.Hour * 1)

	claims := jwt.MapClaims{
		"id":    getUser.UserId,
		"name":  getUser.Fname + " " + getUser.Lname,
		"email": getUser.Email,
		"exp":   expireAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(u.config.JWTSecret))
	if err != nil {
		return nil, "", err
	}

	return getUser, tokenString, nil
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

// GetByEmail implements UserUseCase.
func (u *UserService) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	getUser, err := u.userRepo.GetByEmail(ctx, &email)
	if getUser == nil {
		return nil, exceptions.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return getUser, nil
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
func (u *UserService) UpdateUser(c context.Context, id string, user *entities.User) (*entities.User, error) {
	getUser, err := u.userRepo.GetByID(c, id)
	if getUser == nil {
		return nil, exceptions.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	var updateUser entities.User
	updateUser = entities.User{
		UserId:      id,
		Email:       user.Email,
		Fname:       user.Fname,
		Lname:       user.Lname,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}

	return u.userRepo.Update(c, &updateUser)
}
