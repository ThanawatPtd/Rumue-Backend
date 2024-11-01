package usecases

import (
	"context"
	"time"

	"github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/exceptions"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Login(ctx context.Context, user *entities.User) (string, error)
	Register(ctx context.Context, user *entities.User) error
}

type AuthService struct {
	userRepo repositories.UserRepository
	employeeRepo repositories.EmployeeRepository
	config   *config.Config
}

func ProvideAuthService(userRepo repositories.UserRepository, employeeRepo repositories.EmployeeRepository, config *config.Config) AuthUseCase {
	return &AuthService{
		userRepo: userRepo,
		employeeRepo: employeeRepo,
		config:   config,
	}
}

// Register implements UserUseCase.
func (u *AuthService) Register(ctx context.Context, user *entities.User) error {
	// Check email format
	if err := utils.ValidateEmail(&utils.RegexEmailValidator{}, user.Email); err != nil {
		return err
	}
	if err := utils.ValidatePassword(&utils.RegexPasswordValidator{}, user.Password); err != nil {
		return err
	}
	// Find user by email
	getUser, err := u.userRepo.GetIDPasswordByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	// Check if user already exist
	if getUser != nil {
		return exceptions.ErrDuplicatedEmail
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return u.userRepo.Save(ctx, user)
}

// Login implements UserUseCase.
func (u *AuthService) Login(ctx context.Context, user *entities.User) (string, error) {
	if err := utils.ValidateEmail(&utils.RegexEmailValidator{}, user.Email); err != nil {
		return "", err
	}

	// Find user by email
	getUser, err := u.userRepo.GetIDPasswordByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}
	// Check if user exist
	if getUser == nil { 
		return "", exceptions.ErrUserNotFound
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(user.Password)) != nil {
		return "", exceptions.ErrLoginFailed
	}

	role := "user"
	employee, err := u.employeeRepo.GetByID(ctx, getUser.ID)
	if err != nil {
		return "", err
	}
	if employee != nil {
		role = "employee"
	}

	// Generate JWT token
	expireAt := time.Now().Add(time.Hour * 100)

	claims := jwt.MapClaims{
		"id": getUser.ID,
		"role": role,
		"exp": expireAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(u.config.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
