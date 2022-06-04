package repository

import (
	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/model"
)

func InsertProject(project model.Project) model.Project {
	config.DB.Save(&project)
	//config.DB.Preload("User").Find(&issue)
	return project
}
