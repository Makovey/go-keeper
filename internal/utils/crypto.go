package utils

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/dustin/go-humanize"
)

//go:generate mockgen -source=crypto.go -destination=mock/crypto_mock.go -package=mock
type Crypto interface {
	EncryptString(text string, secret string) (string, error)
	DecryptString(text string, secret string) (string, error)
	EncryptReader(reader *bufio.Reader, secret string) (*bufio.Reader, error)
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

func (c *crypto) EncryptReader(reader *bufio.Reader, secret string) (*bufio.Reader, error) {
	fn := "utils.EncryptReader"

	var b strings.Builder
	buf := make([]byte, humanize.MByte)

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if n == 0 {
			break
		}

		b.Write(buf[:n])
	}

	encrypted, err := c.EncryptString(b.String(), secret)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	return bufio.NewReader(strings.NewReader(encrypted)), nil
}
