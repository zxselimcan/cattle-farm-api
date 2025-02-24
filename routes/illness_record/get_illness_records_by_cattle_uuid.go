package illness_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetIllnessRecordsByCattleByUUID(c *fiber.Ctx) error {

	cattle_uuid := c.Params("cattle_uuid")
	illness_records := &[]models.IllnessRecord{}
	tx := utils.DB.Find(illness_records, models.IllnessRecord{CattleUUID: cattle_uuid})
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message":         "SUCCESS",
		"illness_records": illness_records,
	})

}
