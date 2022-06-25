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

// Returns projects that the user is a member of, with subjects
// @Summary Returns projects that the user is a member of, with subjects
// @Description Returns projects that the user is a member of, with subjects
// @Tags project
// @Accept json
// @Produce json
// @Param user_id path uint64 true "User ID"
// @Success 200 {object} helper.Response{data=response.ProjectNavTreeResponse}
// @Failure 400 {object} helper.Response{data=[]helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /api/project/sidenav-options/{user_id} [get]
func GetProjectsByUserId(context *fiber.Ctx) error {
	user_id, err := strconv.ParseUint(context.Params("user_id"), 10, 64)
	log.Println(user_id)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong UserID Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	result, err := service.GetProjectsByUserId(user_id)
	if err != nil {
		res := helper.BuildErrorResponse("Something went wrong while search", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	response := helper.BuildResponse(true, "OK", result)
	return context.Status(http.StatusOK).JSON(response)
}
