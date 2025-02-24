package birth_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMyBirthRecords(c *fiber.Ctx) error {

	birth_records := &[]models.BirthRecord{}
	err := utils.DB.Preload("Mother").Find(birth_records, models.BirthRecord{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message":       "SUCCESS",
		"birth_records": birth_records,
	})

}
