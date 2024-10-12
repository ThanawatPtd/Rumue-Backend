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

type AuthUseCase interface {
	Login(ctx context.Context, user *entities.User) (*entities.User, string, error)
	Register(ctx context.Context, user *entities.User) (*entities.User, error)
}

type AuthService struct {
	userRepo repositories.UserRepository
	config   *config.Config
}

func ProvideAuthService(userRepo repositories.UserRepository, config *config.Config) AuthUseCase {
	return &AuthService{
		userRepo: userRepo,
		config:   config,
	}
}

// Register implements UserUseCase.
func (u *AuthService) Register(ctx context.Context, user *entities.User) (*entities.User, error) {
	// Find user by email
	getUser, err := u.userRepo.GetByEmail(ctx, user.Email)

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
func (u *AuthService) Login(ctx context.Context, user *entities.User) (*entities.User, string, error) {
	// Find user by email
	getUser, err := u.userRepo.GetByEmail(ctx, user.Email)

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
		"id":    getUser.ID,
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
