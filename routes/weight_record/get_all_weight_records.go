package weight_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllWeightRecords(c *fiber.Ctx) error {

	weight_records := &[]models.WeightRecord{}
	err := utils.DB.Preload("Cattle").Find(weight_records, &models.WeightRecord{
		Cattle: &models.Cattle{
			IsAlive: true,
		},
	}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message":        "SUCCESS",
		"weight_records": weight_records,
	})

}
