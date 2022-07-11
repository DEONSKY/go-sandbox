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

func ProjectExistsByIDAndLeaderID(projectId uint64, projectLeaderId uint64) (bool, error) {
	exists := false
	if result := config.DB.Model(model.Project{}).Select("count(*) > 0").
		Where("id= ? AND project_leader_id = ?", projectId, projectLeaderId).
		Find(&exists); result.Error != nil {
		return false, result.Error
	}
	return exists, nil
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
