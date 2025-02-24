package cattle

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

func NewCattle(c *fiber.Ctx) error {

	user := c.Locals("user").(*models.User)

	type RequestBody struct {
		TagNumber            string `json:"TagNumber" validate:"required"`
		Birthday             string `json:"Birthday" validate:"required"`
		Gender               string `json:"Gender" validate:"required,oneof=male female"`
		LastInseminationDate string `json:"LastInseminationDate"`
		LastGiveBirthDate    string `json:"LastGiveBirthDate"`
		PregnancyStatus      string `json:"PregnancyStatus" validate:"oneof=pregnant not-pregnant inseminated ''"`
	}

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BAD_REQUEST})
	}

	errs := routes.ValidateStruct(body)
	if errs != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: errs[0].Message})
	}

	err := utils.DB.First(&models.Cattle{}, models.Cattle{TagNumber: fmt.Sprintf("%v", body.TagNumber)}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_TAG_NUMBER_ALREADY_IN_USE})
	}

	birthdayParsed, err := time.Parse("02-01-2006", body.Birthday)
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BIRTHDAY_PARSING_FAILED})
	}

	data := &models.Cattle{
		UUID:      uuid.NewString(),
		Birthday:  birthdayParsed,
		TagNumber: fmt.Sprintf("%v", body.TagNumber),
		Gender:    body.Gender,
		Owner:     user,
	}

	if body.Gender == "female" {

		if body.PregnancyStatus == "" {
			return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_SEND_PREGNANCY_STATUS})
		}

		data.PregnancyStatus = body.PregnancyStatus
		if body.PregnancyStatus != "not-pregnant" && body.LastInseminationDate == "" {
			return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_SEND_LAST_INSEMINATION_DATE})
		}

		if body.LastInseminationDate != "" {
			LastInseminationDateParsed, err := time.Parse("02-01-2006", body.LastInseminationDate)
			if err != nil {
				return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_LAST_INSEMINATION_DATE_PARSING_FAILED})
			}
			data.LastInseminationDate = LastInseminationDateParsed
		}

		if body.LastGiveBirthDate != "" {
			LastGiveBirthDateParsed, err := time.Parse("02-01-2006", body.LastGiveBirthDate)
			if err != nil {
				return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_LAST_GIVE_BIRTH_DATE_PARSING_FAILED})
			}
			data.LastGiveBirthDate = LastGiveBirthDateParsed
		}

	}

	err = utils.DB.Create(data).Error
	// err = utils.DB.Model(user).Association("Cattles").Append(data)

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	return c.JSON(fiber.Map{
		"message":     "SUCCESS",
		"cattle_uuid": data.UUID,
	})

}
