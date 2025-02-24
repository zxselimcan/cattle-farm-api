package weight_record

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

func NewWeightRecord(c *fiber.Ctx) error {
	cattle_uuid := c.Params("cattle_uuid")

	type RequestBody struct {
		Date   string `json:"Date" validate:"required"`
		Weight uint   `json:"Weight" validate:"required"`
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
	err := utils.DB.First(cattle, models.Cattle{UUID: fmt.Sprintf("%v", cattle_uuid)}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_FOUND_OR_NOT_PREGNANT})
	}

	dateParsed, err := time.Parse("02-01-2006", body.Date)
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_WEIGHT_DATE_PARSING_FAILED})
	}

	err = utils.DB.Model(cattle).Association("WeightRecords").Append(&models.WeightRecord{
		UUID:   uuid.NewString(),
		Date:   dateParsed,
		Weight: body.Weight,
	})

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	return c.Status(200).JSON(types.ErrorResponse{Message: "SUCCESS"})
}
