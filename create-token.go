package pkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generate a valid jwt token for 10 minutes
func GenerateToken(userID int64, secretKey []byte) (string, error) {
	if secretKey == nil || userID < 1 {
		return "", errors.New("Invalid secret key or userID")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(10 * time.Minute).Unix(), // expire in 10 minutes
	})
	return token.SignedString(secretKey)
}
