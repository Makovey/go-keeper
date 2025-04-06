package mapper

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func ToProtoFromFile(files []*model.ExtendedInfoFile) []*storage.UsersFile {
	var res []*storage.UsersFile

	for _, file := range files {
		res = append(res, &storage.UsersFile{
			FileId:    file.ID,
			FileName:  file.ID,
			FileSize:  file.FileSize,
			CreatedAt: timestamppb.New(file.CreatedAt),
		})
	}

	return res
}
