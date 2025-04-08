package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

type Crypto interface {
	EncryptString(text string, secret string) (string, error)
	DecryptString(text string, secret string) (string, error)
}

type crypto struct {
}

func NewCrypto() Crypto {
	return &crypto{}
}

func (c *crypto) EncryptString(text string, secret string) (string, error) {
	hash := sha256.Sum256([]byte(secret))
	key := hash[:]
	textBytes := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(textBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], textBytes)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func (c *crypto) DecryptString(text string, secret string) (string, error) {
	hash := sha256.Sum256([]byte(secret))
	key := hash[:]

	ciphertext, err := base64.URLEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
