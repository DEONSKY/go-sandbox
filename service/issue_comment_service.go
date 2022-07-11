package service

import (
	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/repository"
	"github.com/DEONSKY/go-sandbox/utils"
	"github.com/mashingan/smapping"
)

func AddIssueComment(issueCommentDto request.IssueCommentCreateRequest) (*model.IssueComment, error) {
	issueCommentToCreate := model.IssueComment{}
	err := smapping.FillStruct(&issueCommentToCreate, smapping.MapFields(&issueCommentDto))
	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "Request DTO Parse Problem", []string{err.Error()})
	}
	res, err := repository.InsertIssueComment(issueCommentToCreate)
	if err != nil {
		return nil, utils.ReturnErrorResponse(422, "Issue comment could not be inserted", []string{err.Error()})
	}
	return res, err
}
