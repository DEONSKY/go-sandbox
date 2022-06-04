package service

import (
	"log"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/repository"
	"github.com/mashingan/smapping"
)

func CreateProject(projectCreateDTO request.ProjectCreateRequest) model.Project {
	projectToCreate := model.Project{}
	err := smapping.FillStruct(&projectToCreate, smapping.MapFields(&projectCreateDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := repository.InsertProject(projectToCreate)
	return res
}
