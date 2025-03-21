package keeper

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/repository/dbo"
	serviceErrors "github.com/Makovey/go-keeper/internal/service"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	"github.com/Makovey/go-keeper/internal/transport/grpc/auth"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

//go:generate mockgen -source=auth.go -destination=../../repository/mock/repository_mock.go -package=mock
type AuthRepository interface {
	RegisterUser(ctx context.Context, user *dbo.User) error
	GetUserInfo(ctx context.Context, email string) (*dbo.User, error)
}

type service struct {
	repo AuthRepository
	jwt  *jwt.Manager
	cfg  config.Config
	log  logger.Logger
}

func NewAuthService(
	repo AuthRepository,
	jwt *jwt.Manager,
	cfg config.Config,
	log logger.Logger,
) auth.Service {
	return &service{
		repo: repo,
		jwt:  jwt,
		cfg:  cfg,
		log:  log,
	}
}

func (s *service) RegisterUser(ctx context.Context, user *model.User) (string, error) {
	fn := "keeper.RegisterUser"

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, serviceErrors.ErrGeneratePassword)
	}

	dboUser := dbo.User{
		ID:           uuid.NewString(),
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: string(pass),
	}

	err = s.repo.RegisterUser(ctx, &dboUser)
	if err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, err)
	}

	token, err := s.jwt.AssembleNewJWT(dboUser.ID)
	if err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, err)
	}

	return token, nil
}

func (s *service) LoginUser(ctx context.Context, login *model.Login) (string, error) {
	fn := "keeper.LoginUser"

	user, err := s.repo.GetUserInfo(ctx, login.Email)
	if err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(login.Password))
	if err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, serviceErrors.ErrIncorrectPassword)
	}

	token, err := s.jwt.AssembleNewJWT(user.ID)
	if err != nil {
		return "", fmt.Errorf("[%s]: %w", fn, err)
	}

	return token, nil
}
