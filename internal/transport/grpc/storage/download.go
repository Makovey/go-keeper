package storage

import (
	"io"

	"github.com/dustin/go-humanize"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	helper "github.com/Makovey/go-keeper/internal/transport/grpc"
)

func (s *Server) DownloadFile(
	req *storage.DownloadRequest,
	stream grpc.ServerStreamingServer[storage.DownloadResponse],
) error {
	fn := "storage.DownloadFile"

	userID, err := helper.GetUserIDFromContext(stream.Context())
	if err != nil {
		return status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
	}

	file, err := s.service.DownloadFile(stream.Context(), userID, req.GetFileId())
	if err != nil {
		return status.Error(codes.InvalidArgument, helper.InternalServerError)
	}

	if err = stream.Send(&storage.DownloadResponse{
		FileName: file.FileName,
	}); err != nil {
		s.log.Errorf("[%s]: %v", fn, err)
		return status.Errorf(codes.Internal, "failed to send file name: %v", err)
	}

	buf := make([]byte, humanize.MByte)
	for {
		n, err := file.Data.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			s.log.Errorf("[%s]: %v", fn, err)
			return status.Errorf(codes.Internal, "failed to read file: %v", err)
		}

		if err := stream.Send(&storage.DownloadResponse{
			ChunkData: buf[:n],
		}); err != nil {
			s.log.Errorf("[%s]: %v", fn, err)
			return status.Errorf(codes.Internal, "failed to send chunk: %v", err)
		}
	}

	return nil
}
