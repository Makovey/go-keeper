package stream

import (
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/Makovey/go-keeper/internal/logger"
)

func Logger(log logger.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		method := strings.Split(info.FullMethod, "/")[len(strings.Split(info.FullMethod, "/"))-1]
		if method == "ServerReflectionInfo" {
			return nil
		}

		log.Infof("[%s]: stream started", info.FullMethod)

		startTime := time.Now()

		err := handler(srv, ss)

		duration := time.Since(startTime)
		if err != nil {
			log.Errorf("[%s]: stream ended with error: %v with duration: %v", info.FullMethod, err, duration)
		} else {
			log.Infof("[%s]: stream ended successfully with duration: %v", info.FullMethod, duration)
		}

		return err
	}
}
