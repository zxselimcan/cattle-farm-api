package milking_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMilkingRecordsByCattleByUUID(c *fiber.Ctx) error {

	cattle_uuid := c.Params("cattle_uuid")
	milking_records := &[]models.MilkingRecord{}
	tx := utils.DB.Find(milking_records, models.MilkingRecord{CattleUUID: cattle_uuid})
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message":         "SUCCESS",
		"milking_records": milking_records,
	})

}
