package auth

import (
	"api/models"
	"api/routes"
	"api/types"
	"api/utils"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func EmailVerification(c *fiber.Ctx) error {
	type RequestBody struct {
		Token string `json:"token" validate:"required,min=128,max=128"`
	}
	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BAD_REQUEST})
	}

	errs := routes.ValidateStruct(body)
	if errs != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: errs[0].Message})
	}

	email, err := utils.Redis.Get(context.Background(), body.Token).Result()
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	fmt.Println(email)

	err = utils.DB.Where(&models.User{Email: email}).UpdateColumns(models.User{EmailVerified: true}).Error
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	err = utils.Redis.Del(context.Background(), body.Token).Err()
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	return c.SendStatus(200)

}
