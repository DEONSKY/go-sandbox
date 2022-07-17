package service

import (
	"github.com/DEONSKY/go-sandbox/constant"
	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/dto/response"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/repository"
	"github.com/DEONSKY/go-sandbox/utils"
	"github.com/mashingan/smapping"
)

type IssueService interface {
	CreateIssue(issueDto request.IssueCreateRequest) (*model.Issue, error)
	GetIssues(issueGetQuery *request.IssueGetQuery, userID uint64) ([]response.IssueResponse, error)
	GetIssuesKanban(issueGetQuery *request.IssueGetQuery, userID uint64) ([]response.IssueKanbanResponse, error)
	InsertDependentIssueAssociation(issueID uint64, dependentIssueID uint64, userID uint64) (*model.Issue, error)
	AssignieIssueToUser(issueID uint64, assignieID uint64, userID uint64) (*model.Issue, error)
}
type issueService struct {
	issueRepository repository.IssueRepository
}

func NewIssueService(issueRepo repository.IssueRepository) IssueService {
	return &issueService{
		issueRepository: issueRepo,
	}
}

func (service *issueService) CreateIssue(issueDto request.IssueCreateRequest) (*model.Issue, error) {
	issueToCreate := model.Issue{}
	if issueDto.Status == 0 {
		issueDto.Status = 1
	}
	err := smapping.FillStruct(&issueToCreate, smapping.MapFields(&issueDto))
	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "Request DTO Parse Problem", []string{err.Error()})
	}
	res, err := service.issueRepository.InsertIssue(issueToCreate)
	if err != nil {
		return nil, utils.ReturnErrorResponse(422, "Issue could not be inserted", []string{err.Error()})
	}
	return res, err
}

func (service *issueService) GetIssues(issueGetQuery *request.IssueGetQuery, userID uint64) ([]response.IssueResponse, error) {

	res, err := service.issueRepository.GetIssues(issueGetQuery, userID)

	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "Cannot get issues", []string{err.Error()})
	}

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

func (service *issueService) GetIssuesKanban(issueGetQuery *request.IssueGetQuery, userID uint64) ([]response.IssueKanbanResponse, error) {

	res, err := service.issueRepository.GetIssues(issueGetQuery, userID)

	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "Cannot get issues", []string{err.Error()})
	}

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

func (service *issueService) InsertDependentIssueAssociation(issueID uint64, dependentIssueID uint64, userID uint64) (*model.Issue, error) {
	issue, err := service.issueRepository.FindIssueByAccess(issueID, userID)
	if err != nil {
		return nil, utils.ReturnErrorResponse(404, "Issue not found", []string{err.Error()})
	}
	dependentIssue, err := service.issueRepository.FindIssueByAccess(dependentIssueID, userID)
	if err != nil {
		return nil, utils.ReturnErrorResponse(404, "Depent Issue not found", []string{err.Error()})
	}
	res, err := service.issueRepository.InsertDependentIssueAssociation(*issue, *dependentIssue)
	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "Depentent Issue insertion error", []string{err.Error()})
	}
	return res, err
}

func (service *issueService) AssignieIssueToUser(issueID uint64, assignieID uint64, userID uint64) (*model.Issue, error) {
	issue, err := service.issueRepository.FindIssueByAccess(issueID, userID)
	if err != nil {
		return nil, utils.ReturnErrorResponse(404, "Issue not found", []string{err.Error()})
	}
	user, err := repository.FindUser(assignieID)
	if err != nil {
		return nil, utils.ReturnErrorResponse(404, "User not found", []string{err.Error()})
	}
	res, err := service.issueRepository.AssignieIssueToUser(*issue, *user)
	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "User assignie associtoation insertion error", []string{err.Error()})
	}
	return res, err
}
