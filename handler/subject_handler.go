package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/DEONSKY/go-sandbox/utils"
	"github.com/gofiber/fiber/v2"
)

// InsertSubject is a function to insert new project into database
// @Summary Insert Subject
// @Description Adds new subject to database
// @Tags Subject
// @Accept json
// @Produce json
// @Param Subject body request.SubjectCreateRequest true "Create Subject"
// @Success 200 {object} helper.Response{data=model.Subject}
// @Failure 400 {object} helper.Response{}
// @Security ApiKeyAuth
// @Router /api/subject [post]
func InsertSubject(context *fiber.Ctx) error {
	var subjectCreateDTO request.SubjectCreateRequest

	err := context.BodyParser(&subjectCreateDTO)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Request DTO Parse Problem", []string{err.Error()})
	}
	userID := context.Locals("user_id").(uint64)
	subjectCreateDTO.TeamLeaderID = userID
	result, err := service.CreateSubject(subjectCreateDTO, userID)
	if err != nil {
		return err
	}
	response := helper.BuildResponse("OK", result)
	return context.Status(http.StatusCreated).JSON(response)

}

// Creates subject - user many2many association
// @Summary Creates subject - user many2many association
// @Description Creates subject - user many2many association
// @Tags Subject
// @Accept json
// @Produce json
// @Param subject_id path string true "Subject ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} helper.Response{data=model.Subject}
// @Failure 400 {object} helper.Response{}
// @Security ApiKeyAuth
// @Router /api/subject/{subject_id}/{user_id} [put]
func InsertUserToSubject(context *fiber.Ctx) error {
	subject_id, err := strconv.ParseUint(context.Params("subject_id"), 10, 64)
	log.Println(subject_id)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Wrong SubjectID Parameter", []string{err.Error()})
	}
	user_id, err := strconv.ParseUint(context.Params("user_id"), 10, 64)
	log.Println(user_id)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Wrong UserID Parameter", []string{err.Error()})
	}
	result, err := service.InsertUserToSubject(subject_id, user_id)
	if err != nil {
		return err
	}

	response := helper.BuildResponse("OK", result)
	return context.Status(http.StatusCreated).JSON(response)
}
