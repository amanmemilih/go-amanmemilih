package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey []byte
}

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// New creates a new JWTManager with a provided secret key
func New(secret string) *JWTManager {
	return &JWTManager{
		secretKey: []byte(secret),
	}
}

// GenerateJWT generates a JWT token for a given user ID
func (j *JWTManager) GenerateJWT(userID int) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

// ValidateJWT validates the JWT token and returns the claims
func (j *JWTManager) ValidateJWT(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
