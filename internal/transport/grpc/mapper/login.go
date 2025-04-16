package mapper

import (
	pb "github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func ToLoginFromProto(user *pb.LoginRequest) *model.Login {
	return &model.Login{
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToProtoFromLogin(login *model.Login) *pb.LoginRequest {
	return &pb.LoginRequest{
		Email:    login.Email,
		Password: login.Password,
	}
}
