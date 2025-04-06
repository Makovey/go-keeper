package unary

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
	"RegisterUser",
	"LoginUser",
}

func JWTAuth(
	log logger.Logger,
	jwtUtils *jwt.Manager,
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fn := "interceptor.Auth"

		method := strings.Split(info.FullMethod, "/")[len(strings.Split(info.FullMethod, "/"))-1]
		if slices.Contains(exceptMethods, method) {
			return handler(ctx, req)
		}

		var token string
		var userID string

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		values := md.Get(jwtMetaName)
		if len(values) > 0 {
			token = values[0]
		}

		if len(token) != 0 {
			userID, err = jwtUtils.ParseUserID(token)
			if err != nil && errors.Is(err, jwt.ErrParseToken) {
				return nil, status.Error(codes.Internal, helper.InternalServerError)
			}

			if userID == "" {
				log.Warn(fmt.Sprintf("[%s]: userID is empty", fn))
				return nil, status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
			}
		}

		if len(token) == 0 || errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrInvalidToken) {
			return nil, status.Error(codes.Unauthenticated, helper.ReloginAndTryAgain)
		}

		ctx = context.WithValue(ctx, jwt.CtxUserIDKey, userID)
		return handler(ctx, req)
	}
}
