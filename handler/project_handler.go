package handler

import (
	"net/http"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/DEONSKY/go-sandbox/utils"
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

	userID := context.Locals("user_id").(uint64)
	err := context.BodyParser(&projectCreateDTO)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request DTO Parse Problem", []string{err.Error()})
	}

	projectCreateDTO.ProjectLeaderID = userID
	result, err := service.CreateProject(projectCreateDTO)
	if err != nil {
		return err
	}
	response := helper.BuildResponse("OK", result)
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
	userID := context.Locals("user_id").(uint64)
	result, err := service.GetProjectsByUserId(userID)

	if err != nil {
		return err
	}

	response := helper.BuildResponse("OK", result)
	return context.Status(http.StatusOK).JSON(response)
}
