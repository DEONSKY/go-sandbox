package repository

import (
	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/model"
)

func InsertIssueComment(issueComment model.IssueComment) (*model.IssueComment, error) {
	if result := config.DB.Save(&issueComment); result.Error != nil {
		return nil, result.Error
	}
	//config.DB.Preload("User").Find(&issue)
	return &issueComment, nil
}
