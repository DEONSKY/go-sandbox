package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/dto/response"
	"github.com/DEONSKY/go-sandbox/model"
)

func InsertIssue(issue model.Issue) (*model.Issue, error) {
	if result := config.DB.Save(&issue); result.Error != nil {
		return nil, result.Error
	}
	//config.DB.Preload("User").Find(&issue)
	return &issue, nil
}

func GetIssues(issueGetQuery *request.IssueGetQuery) ([]response.IssueResponse, error) {

	var issues []response.IssueResponse
	var queryParams []string

	chain := config.DB.Model(&model.Issue{}).Preload("ChildIssues").Preload("DependentIssues")
	log.Println(issueGetQuery)
	var queryCount uint8
	if issueGetQuery.CreatorID != nil {
		queryParams = append(queryParams, "creator_id = @creator_id")
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
	if issueGetQuery.UserID != nil {
		queryParams = append(queryParams, "subject_id IN (@user_id)")
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
			sql.Named("creator_id", issueGetQuery.CreatorID),
			sql.Named("subject_id", issueGetQuery.SubjectID),
			sql.Named("project_id", config.DB.Table("subjects").
				Select("id").
				Where("project_id = ?", issueGetQuery.ProjectID)),
			sql.Named("user_id", config.DB.Table("subject_users").
				Select("subject_id").
				Where("user_id = ?", issueGetQuery.UserID)),
			sql.Named("assignie_id", issueGetQuery.AssignieID),
			sql.Named("status", issueGetQuery.Status),
			sql.Named("parent_issue_id", issueGetQuery.ParentIssueID),
		)
	}
	log.Println(len(queryParams))

	if result := chain.Find(&issues); result.Error != nil {
		return nil, result.Error
	}

	return issues, nil
}

func FindIssue(issue_id uint64) (*model.Issue, error) {
	var issue model.Issue
	if result := config.DB.First(&issue, issue_id); result.Error != nil {
		return nil, result.Error
	}
	return &issue, nil
}

func InsertDependentIssueAssociation(issue model.Issue, dependentIssue model.Issue) (*model.Issue, error) {
	if err := config.DB.Model(&issue).Association("DependentIssues").Append(&dependentIssue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func AssignieIssueToUser(issue model.Issue, user model.User) (*model.Issue, error) {
	if err := config.DB.Model(&issue).Association("Assignie").Append(&user); err != nil {
		return nil, err
	}
	return &issue, nil
}
