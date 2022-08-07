package repository

import (
	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/dto/response"
	"github.com/DEONSKY/go-sandbox/model"
	"gorm.io/gorm"
)

type IssueRepository interface {
	InsertIssue(issue model.Issue) (*model.Issue, error)
	GetIssues(issueGetQuery *request.IssueGetQuery, userID uint64) ([]response.IssueResponse, error)
	FindIssue(issue_id uint64) (*model.Issue, error)
	FindIssueByAccess(issue_id uint64, user_id uint64) (*model.Issue, error)
	InsertDependentIssueAssociation(issue model.Issue, dependentIssue model.Issue) (*model.Issue, error)
	AssignieIssueToUser(issue model.Issue, user model.User) (*model.Issue, error)
}

type issueConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewIssueRepository(dbConn *gorm.DB) IssueRepository {
	return &issueConnection{
		connection: dbConn,
	}
}

func (db *issueConnection) InsertIssue(issue model.Issue) (*model.Issue, error) {
	if result := config.DB.Save(&issue); result.Error != nil {
		return nil, result.Error
	}
	//config.DB.Preload("User").Find(&issue)
	return &issue, nil
}

func (db *issueConnection) GetIssues(issueGetQuery *request.IssueGetQuery, userID uint64) ([]response.IssueResponse, error) {

	var issues []response.IssueResponse

	chain := config.DB.Model(&model.Issue{}).
		Preload("ChildIssues").
		Preload("DependentIssues").
		Preload("Assignie").
		Preload("Reporter").
		Joins("INNER JOIN subjects s on subject_id = s.id").
		Joins("INNER JOIN subject_users su on su.subject_id = s.id").
		Where("user_id", userID)

	if issueGetQuery.ReporterID != nil {
		chain = chain.Where("reporter_id", issueGetQuery.ReporterID)
	}
	if issueGetQuery.SubjectID != nil {
		chain = chain.Where("s.id", issueGetQuery.SubjectID)
	}
	if issueGetQuery.ProjectID != nil {
		chain = chain.Where("project_id", issueGetQuery.ProjectID)
	}
	if issueGetQuery.AssignieID != nil {
		chain = chain.Where("assignie_id", issueGetQuery.AssignieID)
	}
	if issueGetQuery.Status != nil {
		chain = chain.Where("status_id", issueGetQuery.Status)
	}
	if issueGetQuery.ParentIssueID != nil {
		chain = chain.Where("parent_issue_id", issueGetQuery.ParentIssueID)
	}
	if issueGetQuery.GetOnlyOrphans != nil {
		chain = chain.Where("parent_issue_id IS NULL")
	}

	if result := chain.Find(&issues); result.Error != nil {
		return nil, result.Error
	}

	return issues, nil
}

func (db *issueConnection) FindIssue(issue_id uint64) (*model.Issue, error) {
	var issue model.Issue
	if result := config.DB.First(&issue, issue_id); result.Error != nil {
		return nil, result.Error
	}
	return &issue, nil
}

func (db *issueConnection) FindIssueByAccess(issue_id uint64, user_id uint64) (*model.Issue, error) {
	var issue model.Issue
	if result := config.DB.
		Joins("INNER JOIN subjects s on subject_id = s.id").
		Joins("INNER JOIN subject_users su on su.subject_id = s.id").
		Where("user_id", user_id).First(&issue, issue_id); result.Error != nil {
		return nil, result.Error
	}
	return &issue, nil
}

func (db *issueConnection) InsertDependentIssueAssociation(issue model.Issue, dependentIssue model.Issue) (*model.Issue, error) {
	if err := config.DB.Model(&issue).Omit("DependentIssues.*").Association("DependentIssues").Append(&dependentIssue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (db *issueConnection) AssignieIssueToUser(issue model.Issue, user model.User) (*model.Issue, error) {
	if err := config.DB.Model(&issue).Association("Assignie").Append(&user); err != nil {
		return nil, err
	}
	return &issue, nil
}
