package repository

import (
	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/dto/response"
	"github.com/DEONSKY/go-sandbox/model"
)

func InsertProject(project model.Project) (*model.Project, error) {
	if result := config.DB.Save(&project); result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func GetProjectsBySubjectIds(userIDs []uint64) ([]response.ProjectNavTreeResponse, error) {
	var projectNavTreeResponse []response.ProjectNavTreeResponse
	if result := config.DB.Model(&model.Project{}).
		Where("id IN (?)", userIDs).
		Order("Title").
		Find(&projectNavTreeResponse); result.Error != nil {
		return nil, result.Error
	}
	return projectNavTreeResponse, nil
}
