package birth_record

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

func NewBirthRecord(c *fiber.Ctx) error {
	cattle_uuid := c.Params("cattle_uuid")

	user := c.Locals("user").(*models.User)

	type RequestChildrenInfo struct {
		Gender  string `json:"Gender" validate:"required,oneof=male female"`
		IsAlive *bool  `json:"IsAlive" validate:"required"`
	}

	type RequestBody struct {
		Birthday  string                `json:"Birthday" validate:"required"`
		Children  []RequestChildrenInfo `json:"Children" validate:"required"`
		BirthType string                `json:"BirthType" validate:"required,oneof=natural caesarean"`
	}

	type ResponseBodyChildrenInfo struct {
		ChildUUID string `json:"child_uuid"`
		BirthUUID string `json:"birth_uuid"`
	}

	type ResponseBody struct {
		ChildrenInfo []ResponseBodyChildrenInfo `json:"children_info"`
		MotherUUID   string
		Message      string `json:"message"`
	}

	response := ResponseBody{
		Message: "SUCCESS",
	}

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BAD_REQUEST})
	}

	errs := routes.ValidateStruct(body)
	if errs != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: errs[0].Message})
	}

	if len(body.Children) < 1 {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BAD_REQUEST})
	}

	if len(body.Children) > 3 {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BAD_REQUEST})
	}

	mother := &models.Cattle{}
	err := utils.DB.Preload("CurrentInseminationRecord").First(mother, models.Cattle{UUID: fmt.Sprintf("%v", cattle_uuid), Gender: "female", PregnancyStatus: "pregnant"}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_FOUND_OR_NOT_PREGNANT})
	}

	if mother.PregnancyStatus == "not-pregnant" {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_CATTLE_NOT_PREGNANT})
	}

	insemination := &models.InseminationRecord{}

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

	birthdayParsed, err := time.Parse("02-01-2006", body.Birthday)
	if err != nil {
		return c.Status(400).JSON(types.ErrorResponse{Message: types.ERROR_BIRTHDAY_PARSING_FAILED})
	}

	mother.PregnancyStatus = "not-pregnant"
	mother.LastGiveBirthDate = birthdayParsed
	mother.ChildrenCount += uint(len(body.Children))
	mother.InseminationRecordUUID = ""
	mother.CurrentInseminationRecord = &models.InseminationRecord{}
	insemination.Status = "done"
	insemination.StatusUpdateDate = birthdayParsed

	err = utils.DB.Save(mother).Error
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	err = utils.DB.Save(insemination).Error
	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	response.MotherUUID = mother.UUID

	for i := 0; i < len(body.Children); i++ {
		child_item := body.Children[i]

		if *child_item.IsAlive {
			child := &models.Cattle{
				UUID:      uuid.NewString(),
				Birthday:  birthdayParsed,
				Mother:    mother,
				Gender:    child_item.Gender,
				IsAlive:   *child_item.IsAlive,
				OwnerUUID: user.UUID,
			}

			err = utils.DB.Create(child).Error
			if err != nil {
				return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
			}

			birth := &models.BirthRecord{
				UUID:          uuid.NewString(),
				Mother:        mother,
				Child:         child,
				Date:          birthdayParsed,
				ChildrenCount: uint(len(body.Children)),
				Type:          body.BirthType,
			}
			err = utils.DB.Create(birth).Error

			response.ChildrenInfo = append(response.ChildrenInfo, ResponseBodyChildrenInfo{
				ChildUUID: child.UUID,
				BirthUUID: birth.UUID,
			})
		}

	}

	if err != nil {
		return c.Status(500).JSON(types.ErrorResponse{Message: types.ERROR_SERVER_ERROR})
	}

	return c.JSON(response)

}
