package insemination_record

import (
	"api/models"
	"api/routes"
	"api/types"
	"api/utils"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FailedPregnancy(c *fiber.Ctx) error {

	cattle_uuid := c.Params("cattle_uuid")

	type RequestBody struct {
		StatusUpdateDate string `json:"StatusUpdateDate" validate:"required"`
	}

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BAD_REQUEST})
	}

	errs := routes.ValidateStruct(body)
	if errs != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: errs[0].Message})
	}

	statusUpdateDateParsed, err := time.Parse("02-01-2006", body.StatusUpdateDate)
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_STATUS_DATE_PARSING_FAILED})
	}

	mother := &models.Cattle{}
	insemination := &models.InseminationRecord{}

	err = utils.DB.Preload("CurrentInseminationRecord").
		First(
			mother,
			map[string]interface{}{
				"UUID":   fmt.Sprintf("%v", cattle_uuid),
				"Gender": "female",
			},
		).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_FOUND_OR_NOT_INSEMINATED})
	}

	if mother.PregnancyStatus == "not-pregnant" {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_PREGNANT})
	}

	err = utils.DB.First(
		insemination,
		models.InseminationRecord{
			UUID: mother.InseminationRecordUUID,
		},
	).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_INSEMINATION_NOT_FOUND})
	}

	if insemination.Status == "done" || insemination.Status == "failed" {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_INSEMINATION_ALREADY_DONE_OR_FAILED})
	}

	mother.PregnancyStatus = "not-pregnant"
	mother.InseminationRecordUUID = ""
	mother.CurrentInseminationRecord = &models.InseminationRecord{}
	insemination.Status = "failed"
	insemination.StatusUpdateDate = statusUpdateDateParsed

	err = utils.DB.Save(mother).Error
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	err = utils.DB.Save(insemination).Error
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	return c.JSON(fiber.Map{
		"message":     "SUCCESS",
		"mother_uuid": mother.UUID,
	})

}
