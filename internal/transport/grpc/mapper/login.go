package mapper

import (
	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func ToLoginFromProto(user *auth.LoginRequest) *model.Login {
	return &model.Login{
		Email:    user.Email,
		Password: user.Password,
	}
}
