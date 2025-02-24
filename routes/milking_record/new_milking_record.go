package milking_record

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

func NewMilkingRecord(c *fiber.Ctx) error {
	cattle_uuid := c.Params("cattle_uuid")

	type RequestBody struct {
		Date       string `json:"Date" validate:"required"`
		MilkAmount uint   `json:"MilkAmount" validate:"required"`
	}

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BAD_REQUEST})
	}

	errs := routes.ValidateStruct(body)
	if errs != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: errs[0].Message})
	}

	cattle := &models.Cattle{}
	err := utils.DB.First(cattle, models.Cattle{UUID: fmt.Sprintf("%v", cattle_uuid), Gender: "female", IsAlive: true}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_FOUND_OR_NOT_FEMALE})
	}

	dateParsed, err := time.Parse("02-01-2006", body.Date)
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_MILKING_DATE_PARSING_FAILED})
	}

	err = utils.DB.Model(cattle).Association("MilkingRecords").Append(&models.MilkingRecord{
		UUID:       uuid.NewString(),
		Date:       dateParsed,
		MilkAmount: body.MilkAmount,
	})

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	return c.Status(200).JSON(types.ErrorResponse{Message: "SUCCESS"})
}
