package service

import (
	"log"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/dto/response"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/repository"
	"github.com/mashingan/smapping"
)

func CreateIssue(issueDto request.IssueCreateRequest) (*model.Issue, error) {
	issueToCreate := model.Issue{}
	err := smapping.FillStruct(&issueToCreate, smapping.MapFields(&issueDto))

	log.Println("Issue Create Dto", issueToCreate)
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res, err := repository.InsertIssue(issueToCreate)
	return res, err
}

func GetIssues(issueGetQuery *request.IssueGetQuery) ([]response.IssueResponse, error) {

	res, err := repository.GetIssues(issueGetQuery)
	return res, err
}

func InsertDependentIssueAssociation(issueID uint64, dependentIssueID uint64) (*model.Issue, error) {
	issue, err := repository.FindIssue(issueID)
	if err != nil {
		return nil, err
	}
	dependentIssue, err := repository.FindIssue(dependentIssueID)
	if err != nil {
		return nil, err
	}
	res, err := repository.InsertDependentIssueAssociation(*issue, *dependentIssue)
	return res, err
}

func AssignieIssueToUser(issueID uint64, userID uint64) (*model.Issue, error) {
	issue, err := repository.FindIssue(issueID)
	if err != nil {
		return nil, err
	}
	user, err := repository.FindUser(userID)
	if err != nil {
		return nil, err
	}
	res, err := repository.AssignieIssueToUser(*issue, *user)
	return res, err
}
