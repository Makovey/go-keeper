package storage

import (
	"bytes"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	helper "github.com/Makovey/go-keeper/internal/transport/grpc"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func (s *Server) UploadFile(req grpc.ClientStreamingServer[storage.UploadRequest, storage.UploadResponse]) error {
	fn := "storage.UploadFile"

	userID, err := helper.GetUserIDFromContext(req.Context())
	if err != nil {
		return status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
	}

	var buf bytes.Buffer
	var fileName string
	var fileSize int

	for {
		r, err := req.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.log.Errorf("[%s]: %v", fn, err)
			return status.Error(codes.InvalidArgument, "something went wrong, try another file")
		}

		if _, err = buf.Write(r.ChunkData); err != nil {
			s.log.Errorf("[%s]: %v", fn, err)
			return status.Error(codes.Internal, helper.InternalServerError)
		}
		fileSize += len(r.ChunkData)
		if r.Filename != "" {
			fileName = r.Filename
		}
	}

	f := model.File{
		Data:     *bytes.NewReader(buf.Bytes()),
		FileName: fileName,
		FileSize: fileSize,
	}

	fileId, err := s.service.UploadFile(req.Context(), f, userID)
	if err != nil {
		s.log.Errorf("[%s]: %v", fn, err)
		return status.Error(codes.Internal, helper.InternalServerError)
	}

	return req.SendAndClose(&storage.UploadResponse{FileId: fileId})
}
