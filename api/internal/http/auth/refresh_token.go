package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"
)

type RefreshToken struct {
	PlainText string
	Hash      string
	ExpiresAt time.Time
}

func GenerateRefreshToken(ttl time.Duration) (RefreshToken, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return RefreshToken{}, err
	}

	token := base64.RawURLEncoding.EncodeToString(bytes)

	hash := sha256.Sum256([]byte(token))

	return RefreshToken{
		PlainText: token,
		Hash:      hex.EncodeToString(hash[:]),
		ExpiresAt: time.Now().Add(ttl),
	}, nil
}
