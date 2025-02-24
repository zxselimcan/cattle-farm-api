package auth

import (
	"api/models"
	"api/types"
	"api/utils"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BAD_REQUEST})
	}

	user := models.User{}
	err := utils.DB.First(&user, models.User{
		Email:    body.Email,
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte(body.Password)[:])),
	}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(401).JSON(types.ErrorResponse{Message: types.ERROR_WRONG_CREDENTIALS})
	}

	if !user.EmailVerified {
		return c.Status(401).JSON(types.ErrorResponse{Message: types.ERROR_EMAIL_NOT_VERIFIED})
	}

	token, err := utils.GenerateJWT(user.UUID, user.Email, user.IsAdmin)
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})

	}

	// Send the JWT token to the client
	return c.JSON(fiber.Map{
		"token": token,
	})
}
