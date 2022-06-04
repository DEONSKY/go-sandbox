package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/gofiber/fiber/v2"
)

// InsertSubject is a function to insert new project into database
// @Summary Insert Subject
// @Description Adds new subject to database
// @Tags Subject
// @Accept json
// @Produce json
// @Param Project body request.SubjectCreateRequest true "Create Subject"
// @Success 200 {object} helper.Response{data=model.Project}
// @Failure 400 {object} helper.Response{}
// @Security ApiKeyAuth
// @Router /api/subject [post]
func InsertSubject(context *fiber.Ctx) error {
	var subjectCreateDTO request.SubjectCreateRequest
	errDTO := context.BodyParser(&subjectCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	result, err := service.CreateSubject(subjectCreateDTO)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	response := helper.BuildResponse(true, "OK", result)
	return context.Status(http.StatusCreated).JSON(response)

}

func InsertUserToSubject(context *fiber.Ctx) error {
	subject_id, err := strconv.ParseUint(context.Params("subject_id"), 10, 64)
	log.Println(subject_id)
	user_id, err := strconv.ParseUint(context.Params("user_id"), 10, 64)
	log.Println(user_id)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	log.Println("here")
	result := service.InsertUserToSubject(
		subject_id,
		user_id)

	response := helper.BuildResponse(true, "OK", result)
	return context.Status(http.StatusCreated).JSON(response)
}
