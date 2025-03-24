package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Makovey/go-keeper/internal/logger"
)

func Logger(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		message := info.FullMethod

		if _, ok := req.(*emptypb.Empty); !ok {
			message = fmt.Sprintf("%s. Message: %s", info.FullMethod, req)
		}

		log.Info(message)
		return handler(ctx, req)
	}
}
