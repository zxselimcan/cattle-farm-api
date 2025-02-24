package death_record

import (
	"api/models"
	"api/types"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMyDeadCattles(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	cattles := &[]models.Cattle{}
	err := utils.DB.Find(cattles, map[string]interface{}{"is_alive": false, "owner_uuid": user.UUID}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}
	return c.JSON(fiber.Map{
		"message": "SUCCESS",
		"cattles": cattles,
	})

}
