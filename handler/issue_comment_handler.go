package handler

import (
	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/DEONSKY/go-sandbox/utils"
	"github.com/gofiber/fiber/v2"
)

// Insert Issue Comment
// @Summary Create new IssueComment
// @Description Creates new IssueComment
// @Tags IssueComments
// @Accept json
// @Produce json
// @Param Issue body request.IssueCommentCreateRequest true "createIssues"
// @Success 200 {object} helper.Response{data=model.IssueComment}
// @Failure 400 {object} helper.Response{data=helper.EmptyObj}
// @Security ApiKeyAuth
// @Router /api/issue-comment [post]
func AddIssueComment(context *fiber.Ctx) error {
	userID := context.Locals("user_id").(uint64)
	var issueCommentCreateDTO request.IssueCommentCreateRequest

	errDTO := context.BodyParser(&issueCommentCreateDTO)
	if errDTO != nil {
		return utils.ReturnErrorResponse(fiber.StatusInternalServerError, "Request DTO Parse Problem", []string{errDTO.Error()})
	}

	errors := utils.ValidateStruct(issueCommentCreateDTO)
	if errors != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Validation error", errors)
	}

	issueCommentCreateDTO.CreatorID = userID

	result, err := service.AddIssueComment(issueCommentCreateDTO)
	if err != nil {
		return err
	}
	response := helper.BuildResponse("Issue comment created succesfully", result)
	return context.Status(fiber.StatusCreated).JSON(response)

}
