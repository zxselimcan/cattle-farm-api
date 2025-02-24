package birth_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetBirthRecordsByCattleUUID(c *fiber.Ctx) error {
	cattle_uuid := c.Params("cattle_uuid")

	birth_records := &[]models.BirthRecord{}
	err := utils.DB.Preload("Mother").Find(birth_records, models.BirthRecord{MotherUUID: cattle_uuid}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message":       "SUCCESS",
		"birth_records": birth_records,
	})

}
