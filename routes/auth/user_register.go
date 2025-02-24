package auth

import (
	"api/models"
	"api/routes"
	"api/types"
	"api/utils"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {

	type RequestBody struct {
		Email    string `json:"email" validate:"required,email,min=5,max=64"`
		Password string `json:"password" validate:"required,min=8,max=64"`
	}

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BAD_REQUEST})
	}

	errs := routes.ValidateStruct(body)
	if errs != nil {
		return c.Status(401).JSON(types.ErrorResponse{Message: errs[0].Message})
	}

	err := utils.DB.First(&models.User{}, models.User{Email: body.Email}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(401).JSON(types.ErrorResponse{Message: types.ERROR_EMAIL_ALREADY_IN_USE})
	}

	err = utils.DB.Create(&models.User{
		UUID:     uuid.NewString(),
		Email:    body.Email,
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte(body.Password)[:])),
		// EmailVerified: false,
		EmailVerified: true,
		IsAdmin:       false,
	}).Error

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	// defer utils.SendVerificationMail(body.Email)
	// if err != nil {
	// 	return c.Status(500).JSON(ErrorResponse{Message: types.ERROR_SMTP_ERROR})
	// }

	return c.JSON(types.ErrorResponse{
		// Message: "Success! Please Verify Your Email!",
		Message: "SUCCESS",
	})

}
