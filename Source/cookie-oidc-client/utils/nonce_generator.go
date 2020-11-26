package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

type NonceGenerator interface {
	Generate() (string, error)
}

func NewNonceGenerator() NonceGenerator {
	return &nonceGenerator{}
}

type nonceGenerator struct{}

func (g *nonceGenerator) Generate() (string, error) {
	buf := make([]byte, 18)
	_, err := io.ReadAtLeast(rand.Reader, buf, 18)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(buf[0:18]), nil
}
