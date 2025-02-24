package milking_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllMilkingRecords(c *fiber.Ctx) error {

	milking_records := &[]models.MilkingRecord{}
	err := utils.DB.Preload("Cattle").Find(milking_records, &models.MilkingRecord{
		Cattle: &models.Cattle{
			IsAlive: true,
		},
	}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message":         "SUCCESS",
		"milking_records": milking_records,
	})

}
