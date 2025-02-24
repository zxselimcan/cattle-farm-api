package illness_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllIllnessRecords(c *fiber.Ctx) error {

	illness_records := &[]models.IllnessRecord{}
	err := utils.DB.Preload("Cattle").Find(illness_records, &models.IllnessRecord{
		Cattle: &models.Cattle{
			IsAlive: true,
		},
	}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message":         "SUCCESS",
		"illness_records": illness_records,
	})

}
