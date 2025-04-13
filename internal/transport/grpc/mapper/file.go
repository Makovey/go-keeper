package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func ToProtoFromFile(files []*model.ExtendedInfoFile) []*pb.UsersFile {
	var res []*pb.UsersFile

	for _, file := range files {
		res = append(res, &pb.UsersFile{
			FileId:    file.ID,
			FileName:  file.FileName,
			FileSize:  file.FileSize,
			CreatedAt: timestamppb.New(file.CreatedAt),
		})
	}

	return res
}
