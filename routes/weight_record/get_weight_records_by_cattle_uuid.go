package weight_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetWeightRecordsByCattleByUUID(c *fiber.Ctx) error {

	cattle_uuid := c.Params("cattle_uuid")
	weight_records := &[]models.WeightRecord{}
	tx := utils.DB.Find(weight_records, models.WeightRecord{CattleUUID: cattle_uuid})
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message":        "SUCCESS",
		"weight_records": weight_records,
	})

}
