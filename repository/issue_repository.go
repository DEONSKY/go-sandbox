package repository

import (
	"database/sql"
	"log"
	"strings"

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
	var queryParams []string
	queryParams = append(queryParams, "subject_id IN (@user_id)")
	chain := config.DB.Model(&model.Issue{}).Preload("ChildIssues").Preload("DependentIssues") //.Preload("Comments")

	var queryCount uint8
	if issueGetQuery.ReporterID != nil {
		queryParams = append(queryParams, "reporter_id = @reporter_id")
		queryCount++
	}
	if issueGetQuery.SubjectID != nil {
		queryParams = append(queryParams, "subject_id = @subject_id")
		queryCount++
	}
	if issueGetQuery.ProjectID != nil {
		queryParams = append(queryParams, "subject_id IN (@project_id)")
		queryCount++
	}
	if issueGetQuery.AssignieID != nil {
		queryParams = append(queryParams, "assignie_id = @assignie_id")
		queryCount++
	}
	if issueGetQuery.Status != nil {
		queryParams = append(queryParams, "status = @status")
		queryCount++
	}
	if issueGetQuery.ParentIssueID != nil {
		queryParams = append(queryParams, "parent_issue_id = @parent_issue_id")
		queryCount++
	}
	if issueGetQuery.GetOnlyOrphans != nil {
		chain = chain.Where("parent_issue_id IS NULL")
	}
	res := strings.Join(queryParams, " AND ")

	if queryCount == 0 {

	} else {
		chain = chain.Where(res,
			sql.Named("reporter_id", issueGetQuery.ReporterID),
			sql.Named("subject_id", issueGetQuery.SubjectID),
			sql.Named("project_id", config.DB.Table("subjects").
				Select("id").
				Where("project_id = ?", issueGetQuery.ProjectID)),
			sql.Named("user_id", config.DB.Table("subject_users").
				Select("subject_id").
				Where("user_id = ?", userID)),
			sql.Named("assignie_id", issueGetQuery.AssignieID),
			sql.Named("status", issueGetQuery.Status),
			sql.Named("parent_issue_id", issueGetQuery.ParentIssueID),
		)
	}
	log.Println("Param count: ", len(queryParams))

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
	if err := config.DB.Model(&issue).Omit("Assignie").Association("Assignie").Append(&user); err != nil {
		return nil, err
	}
	return &issue, nil
}
