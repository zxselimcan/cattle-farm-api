package illness_record

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

func NewIllnessRecord(c *fiber.Ctx) error {
	cattle_uuid := c.Params("cattle_uuid")

	type RequestBody struct {
		StartDate           string `json:"StartDate" validate:"required"`
		EndDate             string `json:"EndDate"`
		Name                string `json:"Name" validate:"required"`
		AreAntibioticsUsing *bool  `json:"AreAntibioticsUsing" validate:"required"`
		BlocksMilking       *bool  `json:"BlocksMilking" validate:"required"`
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
	err := utils.DB.First(cattle, models.Cattle{UUID: fmt.Sprintf("%v", cattle_uuid), IsAlive: true}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_FOUND_OR_NOT_FEMALE})
	}

	illnessRecord := &models.IllnessRecord{
		UUID:                uuid.NewString(),
		Name:                body.Name,
		AreAntibioticsUsing: *body.AreAntibioticsUsing,
		BlocksMilking:       *body.BlocksMilking,
	}

	illnessRecord.StartDate, err = time.Parse("02-01-2006", body.StartDate)
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_START_DATE_PARSING_FAILED})
	}

	if body.EndDate != "" {
		illnessRecord.EndDate, err = time.Parse("02-01-2006", body.EndDate)
		if err != nil {
			return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_END_DATE_PARSING_FAILED})
		}
	}

	err = utils.DB.Model(cattle).Association("IllnessRecords").Append(illnessRecord)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	return c.Status(200).JSON(types.ErrorResponse{Message: "SUCCESS"})
}
