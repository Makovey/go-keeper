package storage

import (
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	grpcErr "github.com/Makovey/go-keeper/internal/transport/grpc"
	"github.com/Makovey/go-keeper/internal/transport/grpc/model"
)

func (s *Server) UploadFile(req grpc.ClientStreamingServer[storage.UploadRequest, storage.UploadResponse]) error {
	fn := "storage.UploadFile"

	var f model.File
	for {
		r, err := req.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.log.Errorf("[%s]: %v", fn, err)
			return status.Error(codes.InvalidArgument, "something went wrong, try another file")
		}

		f.Data = append(f.Data, r.ChunkData...)
		f.FileSize += len(r.ChunkData)
		if r.Filename != "" {
			f.FileName = r.Filename
		}
	}

	fileId, err := s.service.UploadFile(req.Context(), f, "userID")
	if err != nil {
		s.log.Errorf("[%s]: %v", fn, err)
		return status.Error(codes.Internal, grpcErr.InternalServerError)
	}

	return req.SendAndClose(&storage.UploadResponse{FileId: fileId})
}
