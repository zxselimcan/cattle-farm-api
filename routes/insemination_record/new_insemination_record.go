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
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewInseminationRecord(c *fiber.Ctx) error {
	cattle_uuid := c.Params("cattle_uuid")

	type RequestBody struct {
		FatherUUID       string `json:"FatherUUID"`
		InseminationDate string `json:"InseminationDate" validate:"required"`
		InseminationType string `json:"InseminationType" validate:"required,oneof=natural artifical"`
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
	err := utils.DB.First(cattle, map[string]interface{}{
		"UUID":             fmt.Sprintf("%v", cattle_uuid),
		"Gender":           "female",
		"pregnancy_status": "not-pregnant",
		"is_alive":         true,
	}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_FOUND_OR_ALREADY_PREGNANT})
	}

	father := &models.Cattle{}
	if body.FatherUUID != "" {
		err := utils.DB.First(father, map[string]interface{}{
			"UUID":     fmt.Sprintf("%v", cattle_uuid),
			"Gender":   "male",
			"is_alive": true,
		}).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_FOUND})
		}
	}

	inseminationDateParsed, err := time.Parse("02-01-2006", body.InseminationDate)
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_INSEMINATION_DATE_PARSING_FAILED})
	}

	cattle.PregnancyStatus = "inseminated"
	cattle.LastInseminationDate = inseminationDateParsed

	inseminationRecord := models.InseminationRecord{
		UUID:             uuid.NewString(),
		Date:             inseminationDateParsed,
		Type:             body.InseminationType,
		Status:           "uncertain",
		StatusUpdateDate: inseminationDateParsed,
	}

	if body.FatherUUID != "" {
		inseminationRecord.Father = father
	}

	err = utils.DB.Model(cattle).Association("InseminationRecords").Append(&inseminationRecord)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}
	cattle.CurrentInseminationRecord = &inseminationRecord

	err = utils.DB.Save(cattle).Error
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	return c.JSON(fiber.Map{
		"message":     "SUCCESS",
		"mother_uuid": cattle.UUID,
	})

}
