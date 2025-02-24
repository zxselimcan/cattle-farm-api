package insemination_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetInseminationRecordsByCattleUUID(c *fiber.Ctx) error {
	cattle_uuid := c.Params("cattle_uuid")

	insemination_records := &[]models.InseminationRecord{}
	err := utils.DB.Preload("Mother").Find(insemination_records, models.InseminationRecord{MotherUUID: cattle_uuid}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message":              "SUCCESS",
		"insemination_records": insemination_records,
	})

}
