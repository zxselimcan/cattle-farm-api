package cattle

import (
	"api/lib"
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMyMilkableCattles(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	cattles := []models.Cattle{}
	err := utils.DB.Find(&cattles, models.Cattle{IsAlive: true, OwnerUUID: user.UUID}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}

	milkable_cattles := []models.Cattle{}

	for _, cattle := range cattles {
		if lib.GetCattleMilkablePeriod(&cattle) == "MILKABLE" {
			milkable_cattles = append(milkable_cattles, cattle)
		}
		// cattles[i].Period = GetCattleMilkablePeriod(&cattle)
	}

	return c.JSON(fiber.Map{
		"message": "SUCCESS",
		"cattles": milkable_cattles,
	})

}
