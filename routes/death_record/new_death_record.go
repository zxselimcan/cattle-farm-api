package death_record

import (
	"api/models"
	"api/routes"
	"api/types"
	"api/utils"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewDeathRecord(c *fiber.Ctx) error {

	type RequestBody struct {
		CattleUUID string `json:"CattleUUID" validate:"required"`
		Date       string `json:"Date" validate:"required"`
		Cause      string `json:"Cause" validate:"required"`
	}

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_NOT_FOUND})
	}

	errs := routes.ValidateStruct(body)
	if errs != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: errs[0].Message})
	}

	cattle := &models.Cattle{}
	err := utils.DB.First(cattle, models.Cattle{UUID: fmt.Sprintf("%v", body.CattleUUID), IsAlive: true}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_FOUND})
	}

	deathDateParsed, err := time.Parse("02-01-2006", body.Date)
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_DEATH_DATE_PARSING_FAILED})
	}

	cattle.IsAlive = false

	err = utils.DB.Save(cattle).Error
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	death := &models.DeathRecord{
		UUID:   uuid.NewString(),
		Cattle: cattle,
		Date:   deathDateParsed,
		Cause:  body.Cause,
	}

	err = utils.DB.Create(death).Error
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	return c.JSON(fiber.Map{
		"message":     "SUCCESS",
		"cattle_uuid": cattle.UUID,
		"death_uuid":  death.UUID,
	})

}
