package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userUUID string, email string, isAdmin bool) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["UUID"] = userUUID
	claims["Email"] = email
	claims["IsAdmin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	// Generate the JWT token string
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
