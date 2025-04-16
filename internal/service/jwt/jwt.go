package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Makovey/go-keeper/internal/config"
)

type Key string

const (
	CtxUserIDKey Key = "UserID"
	tokenExp         = time.Hour * 24
)

var (
	ErrSigningMethod = errors.New("unexpected signing method")
	ErrParseToken    = errors.New("failed to parse token")
	ErrInvalidToken  = errors.New("invalid token")
	ErrTokenExpired  = errors.New("token is expired")
)

type Manager struct {
	cfg config.Config
}

func NewManager(
	cfg config.Config,
) *Manager {
	return &Manager{
		cfg: cfg,
	}
}

type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

func (m *Manager) AssembleNewJWT(userID string) (string, error) {
	fn := "jwt.AssembleNewJWT"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExp)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv(m.cfg.SecretKey())))
	if err != nil {
		return "", fmt.Errorf("[%s]: %v", fn, err)
	}

	return tokenString, nil
}

func (m *Manager) ParseUserID(tokenString string) (string, error) {
	fn := "jwt.ParseUserID"

	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[%s]: %v", fn, ErrSigningMethod)
		}

		return []byte(os.Getenv(m.cfg.SecretKey())), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", fmt.Errorf("[%s]: %v", fn, ErrTokenExpired)
		}
		return "", fmt.Errorf("[%s]: %v", fn, ErrParseToken)
	}

	if !token.Valid {
		return "", fmt.Errorf("[%s]: %v", fn, ErrInvalidToken)
	}

	return claims.UserID, nil
}
