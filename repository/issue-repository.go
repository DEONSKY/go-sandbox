package repository

import (
	"database/sql"
	"strings"

	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/dto/response"
	"github.com/DEONSKY/go-sandbox/model"
	"gorm.io/gorm"
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
	if issueGetQuery.CreatorID != nil {
		queryParams = append(queryParams, "creator_id = @creator_id")
	}
	if issueGetQuery.SubjectID != nil {
		queryParams = append(queryParams, "subject_id = @subject_id")
	}
	if issueGetQuery.ProjectID != nil {
		queryParams = append(queryParams, "subject_id IN (@project_id)")
	}
	if issueGetQuery.UserID != nil {
		queryParams = append(queryParams, "subject_id IN (@user_id)")
	}
	if issueGetQuery.AssignieID != nil {
		queryParams = append(queryParams, "assignie_id = @assignie_id")
	}
	if issueGetQuery.Status != nil {
		queryParams = append(queryParams, "status = @status")
	}
	if issueGetQuery.ParentIssueID != nil {
		queryParams = append(queryParams, "parent_id = @parent_id")
	}
	res := strings.Join(queryParams, " AND ")

	if result := config.DB.Model(&model.Issue{}).Preload("ChildIssues", func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&model.Issue{}).Select("Title", "Description", "IssueForeignId", "TargetTime", "Status", "ParentIssueID", "CreatorID", "AssignieID", "CreatedAt", "UpdatedAt")
	}).Where(res,
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
		sql.Named("parent_id", issueGetQuery.ParentIssueID),
	).Find(&issues); result.Error != nil {
		return nil, result.Error
	}

	return issues, nil
}
