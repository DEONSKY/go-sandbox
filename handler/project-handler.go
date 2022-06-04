package handler

import (
	"log"
	"net/http"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/gofiber/fiber/v2"
)

// AddProject is a function to insert new project into database
// @Summary Insert Project
// @Description Adds new project to database
// @Tags project
// @Accept json
// @Produce json
// @Param Project body request.ProjectCreateRequest true "Create Project"
// @Success 200 {object} helper.Response{data=model.Project}
// @Failure 400 {object} helper.Response{}
// @Security ApiKeyAuth
// @Router /api/project [post]
func InsertProject(context *fiber.Ctx) error {
	var projectCreateDTO request.ProjectCreateRequest
	log.Println("Here")
	errDTO := context.BodyParser(&projectCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	log.Println("Here")
	result := service.CreateProject(projectCreateDTO)
	response := helper.BuildResponse(true, "OK", result)
	return context.Status(http.StatusCreated).JSON(response)

}
