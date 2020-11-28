package nonce

import (
	"crypto/rand"
	"encoding/base64"
	"io"

	"go.uber.org/zap"
)

// Generator generates nonces
type Generator interface {
	// Generate returns a new Nonce
	Generate() (Nonce, error)
}

// NewGenerator returns a new Generator
func NewGenerator(configuration Configuration, logger zap.Logger) Generator {
	return &generator{
		configuration: configuration,
		logger:        logger,
	}
}

type generator struct {
	configuration Configuration
	logger        zap.Logger
}

func (g *generator) Generate() (Nonce, error) {
	nonceLength := g.configuration.Size()

	buffer := make([]byte, nonceLength)
	if _, err := io.ReadAtLeast(rand.Reader, buffer, nonceLength); err != nil {
		g.logger.Error("could not generate nonce", zap.Error(err))
		return "", err
	}

	nonce := base64.URLEncoding.EncodeToString(buffer[:nonceLength])
	return Nonce(nonce), nil
}
