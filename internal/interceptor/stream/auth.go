package stream

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/service/jwt"
	helper "github.com/Makovey/go-keeper/internal/transport/grpc"
)

const (
	jwtMetaName = "jwt"
)

var exceptMethods = []string{
	"ServerReflectionInfo", // reflection debug mode
}

func JWTAuth(
	log logger.Logger,
	jwtUtils *jwt.Manager,
) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		fn := "interceptor.Auth"

		method := strings.Split(info.FullMethod, "/")[len(strings.Split(info.FullMethod, "/"))-1]
		if slices.Contains(exceptMethods, method) {
			return handler(srv, ss)
		}

		var token string
		var userID string
		var err error

		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		values := md.Get(jwtMetaName)
		if len(values) > 0 {
			token = values[0]
		}

		if len(token) != 0 {
			userID, err = jwtUtils.ParseUserID(token)
			if err != nil && errors.Is(err, jwt.ErrParseToken) {
				return status.Error(codes.Internal, helper.InternalServerError)
			}

			if userID == "" {
				log.Warn(fmt.Sprintf("[%s]: userID is empty", fn))
				return status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
			}
		}

		if len(token) == 0 || errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrInvalidToken) {
			return status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
		}

		ctx := context.WithValue(ss.Context(), jwt.CtxUserIDKey, userID)

		wrappedStream := &wrappedServerStream{
			ServerStream: ss,
			ctx:          ctx,
		}

		return handler(srv, wrappedStream)
	}
}

type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}
