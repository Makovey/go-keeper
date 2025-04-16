package mapper

import (
	pb "github.com/Makovey/go-keeper/internal/gen/auth"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func ToUserFromProto(user *pb.User) *model.User {
	return &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToProtoFromUser(user *model.User) *pb.User {
	return &pb.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
