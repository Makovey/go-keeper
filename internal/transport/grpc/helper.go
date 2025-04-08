package grpc

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/Makovey/go-keeper/internal/gen/storage"
	"github.com/Makovey/go-keeper/internal/service/jwt"
)

const (
	InternalServerError = "internal server error"
	ReloginAndTryAgain  = "please, relogin again, to get access to this resource"
	InvalidArgument     = "invalid argument"
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

type TextSecure string

const (
	Secure   TextSecure = "secure"
	Unsecure TextSecure = "unsecure"
)

func MapProtoToLocalTextSecure(textType storage.TextType) TextSecure {
	switch textType {
	case storage.TextType_secure:
		return Secure
	case storage.TextType_unsecure:
		return Unsecure
	}

	return Unsecure
}
