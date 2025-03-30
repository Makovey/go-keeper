package grpc

import (
	"context"

	"github.com/Makovey/go-keeper/internal/service/jwt"
)

const (
	InternalServerError = "internal server error"
	uuidLength          = 36
	ReloginAndTryAgain  = "please, relogin again, to get access to this resource"
)

func GetUserIDFromContext(ctx context.Context) (string, error) {
	if ctx.Value(jwt.CtxUserIDKey) == nil {
		return "", nil
	}

	userID := ctx.Value(jwt.CtxUserIDKey).(string)
	if userID == "" || len(userID) != uuidLength {
		return "", nil
	}

	return userID, nil
}
