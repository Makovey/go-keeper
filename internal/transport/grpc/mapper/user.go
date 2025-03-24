package mapper

import (
	"github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func ToUserFromProto(user *auth.User) *model.User {
	return &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToProtoFromUser(user *model.User) *auth.User {
	return &auth.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
