package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`

	jwt.RegisteredClaims
}

func GenerateAccessToken(
	userID string,
	// sessionID string,
	firstName string,
	lastName string,
	email string,
	// role string,
	secretKey []byte,
	duration time.Duration,
) (string, error) {
	now := time.Now()

	claims := AccessTokenClaims{
		UserID: userID,
		// SessionID: sessionID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		// Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ValidateAccessToken(
	tokenString string,
	secret []byte,
) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&AccessTokenClaims{},
		func(token *jwt.Token) (any, error) {
			return secret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
