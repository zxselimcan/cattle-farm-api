package birth_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetBirthRecordByUUID(c *fiber.Ctx) error {

	uuid := c.Params("uuid")

	birth := &models.BirthRecord{}
	tx := utils.DB.Preload("Mother").Find(birth, models.BirthRecord{UUID: uuid})
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) || tx.RowsAffected != 1 {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message": "SUCCESS",
		"birth":   birth,
	})

}
