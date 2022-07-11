package handler

import (
	"log"
	"strconv"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/DEONSKY/go-sandbox/utils"
	"github.com/gofiber/fiber/v2"
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
	userID := context.Locals("user_id").(uint64)
	var IssueCreateDTO request.IssueCreateRequest

	errDTO := context.BodyParser(&IssueCreateDTO)
	if errDTO != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request DTO Parse Problem", []string{errDTO.Error()})
	}

	errors := utils.ValidateStruct(IssueCreateDTO)
	if errors != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Validation error", errors)
	}

	IssueCreateDTO.ReporterID = userID

	result, err := service.CreateIssue(IssueCreateDTO)
	if err != nil {
		return err
	}
	response := helper.BuildResponse("Issue created succesfully", result)
	return context.Status(fiber.StatusCreated).JSON(response)

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
	userID := context.Locals("user_id").(uint64)

	if err := context.QueryParser(iq); err != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request Query parse problem", []string{err.Error()})
	}
	log.Println("Params", iq)
	if iq.GetOnlyOrphans != nil && iq.ParentIssueID != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request Error", []string{"An issue cannot be orphan and has parent at the same time"})
	}
	result, err := service.GetIssues(iq, userID)
	if err != nil {
		return err
	}
	response := helper.BuildShortResponse(result)
	return context.Status(fiber.StatusOK).JSON(response)
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

	userID := context.Locals("user_id").(uint64)

	if err := context.QueryParser(iq); err != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request Query parse problem", []string{err.Error()})
	}
	log.Println(iq)
	if iq.GetOnlyOrphans != nil && iq.ParentIssueID != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request Error", []string{"An issue cannot be orphan and has parent at the same time"})
	}
	result, err := service.GetIssuesKanban(iq, userID)
	if err != nil {
		return err
	}
	response := helper.BuildShortResponse(result)
	return context.Status(fiber.StatusOK).JSON(response)
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
	userID := context.Locals("user_id").(uint64)

	log.Println(issueID)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Wrong issue parameter", []string{err.Error()})
	}
	dependentIssueID, err := strconv.ParseUint(context.Params("dependent_issue_id"), 10, 64)
	log.Println(dependentIssueID)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Wrong dependent issue parameter", []string{err.Error()})
	}

	result, err := service.InsertDependentIssueAssociation(issueID, dependentIssueID, userID)
	if err != nil {
		return err
	}

	response := helper.BuildResponse("OK", result)
	return context.Status(fiber.StatusCreated).JSON(response)
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
	userID := context.Locals("user_id").(uint64)

	log.Println(issueID)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong Issue Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(fiber.StatusBadRequest).JSON(res)
	}
	assignieID, err := strconv.ParseUint(context.Params("user_id"), 10, 64)
	log.Println(assignieID)
	if err != nil {
		res := helper.BuildErrorResponse("Wrong User Parameter", err.Error(), helper.EmptyObj{})
		return context.Status(fiber.StatusBadRequest).JSON(res)
	}
	log.Println("here")
	result, err := service.AssignieIssueToUser(issueID, assignieID, userID)
	if err != nil {
		return err
	}

	response := helper.BuildResponse("OK", result)
	return context.Status(fiber.StatusCreated).JSON(response)
}
