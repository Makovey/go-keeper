package grpc

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/Makovey/go-keeper/internal/service/jwt"
)

const (
	InternalServerError = "internal server error"
	ReloginAndTryAgain  = "please, relogin again, to get access to this resource"
)

func GetUserIDFromContext(ctx context.Context) (string, error) {
	if ctx.Value(jwt.CtxUserIDKey) == nil {
		return "", errors.New("token doesn't exist")
	}

	userID := ctx.Value(jwt.CtxUserIDKey).(string)
	if userID == "" {
		return "", errors.New("token is empty")
	}

	if _, err := uuid.Parse(userID); err != nil {
		return "", errors.New("invalid user id")
	}

	return userID, nil
}
