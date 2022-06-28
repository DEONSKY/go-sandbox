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
	log.Println(iq)
	if iq.GetOnlyOrphans != nil && iq.ParentIssueID != nil {
		res := helper.BuildErrorResponse("Request Error", "An issue cannot be orphan and has parent at the same time", helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	result, err := service.GetIssues(iq)
	if err != nil {
		res := helper.BuildErrorResponse("Repository Error", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	response := helper.BuildResponse(true, "OK", result)
	return context.Status(http.StatusOK).JSON(response)
}

// GetIssuesKanban is a function to get all issues data from database with dynamic query parameters as Kanban format
// @Summary Get all issues as Kanban Format with query parameters
// @Description Get all issues as Kanban Format with query parameters
// @Tags Issues
// @Accept json
// @Produce json
// @Param Issue query request.IssueGetQuery true "getIssues"
// @Success 200 {object} helper.Response{data=[]response.IssueKanbanResponse}
// @Failure 400 {object} helper.Response{data=[]helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /api/issue/kanban [get]
func GetIssuesKanban(context *fiber.Ctx) error {
	iq := new(request.IssueGetQuery)

	if err := context.QueryParser(iq); err != nil {
		res := helper.BuildCustomErrorResponse("Missed required parameters", "Query must be contain one of them of this parameters:"+
			" user_id subject_id, project_id, creator_id, assignie_id, parent_issue_id")
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	log.Println(iq)
	if iq.GetOnlyOrphans != nil && iq.ParentIssueID != nil {
		res := helper.BuildErrorResponse("Request Error", "An issue cannot be orphan and has parent at the same time", helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	result, err := service.GetIssuesKanban(iq)
	if err != nil {
		res := helper.BuildErrorResponse("Repository Error", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	response := helper.BuildResponse(true, "OK", result)
	return context.Status(http.StatusOK).JSON(response)
}

// InsertDependentIssueAssociation adds assocation between issue and dependent issue
// @Summary Adds assocation with issue and dependent issue
// @Description Adds assocation with issue and dependent issue
// @Tags Issues
// @Accept json
// @Produce json
// @Param issue_id path string true "Issue ID"
// @Param dependent_issue_id path string true "Dependent Issue ID"
// @Success 200 {object} helper.Response{data=model.Issue}
// @Failure 400 {object} helper.Response{data=[]helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /add-issue-dependency/{issue_id}/{dependent_issue_id} [put]
func InsertDependentIssueAssociation(context *fiber.Ctx) error {
	issueID, err := strconv.ParseUint(context.Params("issue_id"), 10, 64)
	log.Println(issueID)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong Issue Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	dependentIssueID, err := strconv.ParseUint(context.Params("dependent_issue_id"), 10, 64)
	log.Println(dependentIssueID)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong Dependent Issue Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	log.Println("here")
	result, err := service.InsertDependentIssueAssociation(
		issueID,
		dependentIssueID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := helper.BuildResponse(true, "OK", result)
	return context.Status(http.StatusCreated).JSON(response)
}

// AssignieIssueToUser adds assocation between issue and assigned user
// @Summary Adds assocation between issue and assigned user
// @Description Adds assocation with issue and dependent issue
// @Tags Issues
// @Accept json
// @Produce json
// @Param issue_id path string true "Issue ID"
// @Param user_id path string true "Assignie User ID"
// @Success 200 {object} helper.Response{data=model.Issue}
// @Failure 400 {object} helper.Response{data=[]helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /assignie-user/{issue_id}/{user_id} [put]
func AssignieIssueToUser(context *fiber.Ctx) error {
	issueID, err := strconv.ParseUint(context.Params("issue_id"), 10, 64)
	log.Println(issueID)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong Issue Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	userID, err := strconv.ParseUint(context.Params("user_id"), 10, 64)
	log.Println(userID)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong User Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}
	log.Println("here")
	result, err := service.AssignieIssueToUser(
		issueID,
		userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := helper.BuildResponse(true, "OK", result)
	return context.Status(http.StatusCreated).JSON(response)
}
