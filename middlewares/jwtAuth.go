package middlewares

import (
	"fmt"
	"os"

	"api/models"
	"api/types"
	"api/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type JWTClaims struct {
	jwt.StandardClaims
	IsAdmin bool   `json:"IsAdmin"`
	UUID    string `json:"UUID"`
	Email   string `json:"Email"`
}

func VerifyJwtKey(adminOnly bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(401).JSON(types.ErrorResponse{Message: types.ERROR_UNAUTHORIZED})
		}

		// Parse the JWT token
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Check that the signing method is correct
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return c.Status(401).JSON(types.ErrorResponse{Message: types.ERROR_UNAUTHORIZED})
		}

		// Check that the token is valid
		if !token.Valid {
			return c.Status(401).JSON(types.ErrorResponse{Message: types.ERROR_UNAUTHORIZED})
		}

		parsedClaims := token.Claims.(*JWTClaims)

		if parsedClaims.IsAdmin != adminOnly {
			return c.Status(401).JSON(types.ErrorResponse{Message: types.ERROR_UNAUTHORIZED})
		}

		var user *models.User
		tx := utils.DB.Where(models.User{UUID: parsedClaims.UUID}).First(&user)
		if tx.Error != nil {
			return c.Status(401).JSON(types.ErrorResponse{Message: types.ERROR_UNAUTHORIZED})
		}

		if tx.RowsAffected != 1 {
			return c.Status(401).JSON(types.ErrorResponse{Message: types.ERROR_UNAUTHORIZED})
		}

		c.Locals("user", user)

		return c.Next()
	}
}
