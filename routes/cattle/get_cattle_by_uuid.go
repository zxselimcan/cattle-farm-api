package cattle

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetCattleByUUID(c *fiber.Ctx) error {

	cattle_uuid := c.Params("cattle_uuid")

	cattle := &models.Cattle{}
	tx := utils.DB.Preload("Mother").Preload("Father").Preload("CurrentInseminationRecord").Find(cattle, models.Cattle{IsAlive: true, UUID: cattle_uuid})
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) || tx.RowsAffected != 1 {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message": "SUCCESS",
		"cattle":  cattle,
	})

}
