package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// InserIssue is a function to create new Issue
// @Summary Create new Issue
// @Description Creates new issue
// @Tags Issues
// @Accept json
// @Produce json
// @Param Issue body request.IssueCreateRequest true "createIssues"
// @Success 200 {object} helper.Response{data=model.Issue}
// @Failure 400 {object} helper.Response{data=helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /api/issue [post]
func InsertIssue(context *fiber.Ctx) error {
	user := context.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var IssueCreateDTO request.IssueCreateRequest
	log.Println("Here")
	errDTO := context.BodyParser(&IssueCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}

	id, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)
	if err == nil {
		IssueCreateDTO.CreatorID = id
	}
	result, err := service.CreateIssue(IssueCreateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	response := helper.BuildResponse(true, "Issue created succesfully", result)
	return context.Status(http.StatusCreated).JSON(response)

}

// GetIssues is a function to get all issues data from database with dynamic query parameters
// @Summary Get all issues with query parameters
// @Description GetIssues is a function to get all issues data from database with dynamic query parameters
// @Tags Issues
// @Accept json
// @Produce json
// @Param Issue query request.IssueGetQuery true "getIssues"
// @Success 200 {object} helper.Response{data=[]response.IssueResponse}
// @Failure 400 {object} helper.Response{data=[]helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /api/issue [get]
func GetIssues(context *fiber.Ctx) error {
	iq := new(request.IssueGetQuery)

	if err := context.QueryParser(iq); err != nil {
		res := helper.BuildCustomErrorResponse("Missed required parameters", "Query must be contain one of them of this parameters:"+
			" user_id subject_id, project_id, creator_id, assignie_id, parent_issue_id")
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	log.Println(iq.SubjectID)
	result, err := service.GetIssues(iq)
	if err != nil {
		res := helper.BuildErrorResponse("Repository Error", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	response := helper.BuildResponse(true, "OK", result)
	return context.Status(http.StatusOK).JSON(response)
}
