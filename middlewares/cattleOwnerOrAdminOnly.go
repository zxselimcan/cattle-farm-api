package middlewares

import (
	"api/models"
	"api/types"
	"api/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CattleOwnerOrAdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {

		user := c.Locals("user").(*models.User)

		if user.IsAdmin {
			return c.Next()
		}

		cattle_uuid := c.Params("cattle_uuid")

		var cattle *models.Cattle
		tx := utils.DB.First(&cattle, models.Cattle{UUID: cattle_uuid})
		if tx.Error != nil {
			fmt.Println("1")
			return c.Status(403).JSON(types.ErrorResponse{Message: types.ERROR_UNAUTHORIZED})
		}

		if tx.RowsAffected < 1 {
			fmt.Println("2")
			return c.Status(403).JSON(types.ErrorResponse{Message: types.ERROR_UNAUTHORIZED})
		}

		if cattle.OwnerUUID != user.UUID {
			fmt.Println(cattle_uuid)
			fmt.Println(cattle)
			return c.Status(403).JSON(types.ErrorResponse{Message: types.ERROR_UNAUTHORIZED})
		}

		c.Locals("cattle", cattle)

		return c.Next()
	}
}
