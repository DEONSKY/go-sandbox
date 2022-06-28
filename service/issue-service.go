package service

import (
	"log"

	"github.com/DEONSKY/go-sandbox/constant"
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

	for i, issue := range res {
		res[i].Status = response.StatusResponse(constant.PredefinedStatusMap[issue.StatusID])
		for j, childIssues := range res[i].ChildIssues {
			res[i].ChildIssues[j].Status = response.StatusResponse(constant.PredefinedStatusMap[childIssues.StatusID])
		}
		for k, dependentIssues := range res[i].DependentIssues {
			res[i].DependentIssues[k].Status = response.StatusResponse(constant.PredefinedStatusMap[dependentIssues.StatusID])
		}
	}
	return res, err
}

func GetIssuesKanban(issueGetQuery *request.IssueGetQuery) ([]response.IssueKanbanResponse, error) {

	res, err := repository.GetIssues(issueGetQuery)

	issueResponseMap := make(map[uint32][]response.IssueResponse)

	for i, issue := range res {
		issue.Status = response.StatusResponse(constant.PredefinedStatusMap[issue.StatusID])
		for j, childIssues := range res[i].ChildIssues {
			issue.ChildIssues[j].Status = response.StatusResponse(constant.PredefinedStatusMap[childIssues.StatusID])
		}
		for k, dependentIssues := range res[i].DependentIssues {
			issue.DependentIssues[k].Status = response.StatusResponse(constant.PredefinedStatusMap[dependentIssues.StatusID])
		}
		issueResponseMap[issue.StatusID] = append(issueResponseMap[issue.StatusID], issue)
	}

	issueKanbanSlice := make([]response.IssueKanbanResponse, 0, len(issueResponseMap))

	for i := range issueResponseMap {
		var issueKanban response.IssueKanbanResponse
		issueKanban.Issues = issueResponseMap[i]
		issueKanban.Status = response.StatusResponse(constant.PredefinedStatusMap[i])
		issueKanbanSlice = append(issueKanbanSlice, issueKanban)
	}
	return issueKanbanSlice, err
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
