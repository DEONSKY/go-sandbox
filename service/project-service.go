package service

import (
	"log"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/dto/response"
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

func GetProjectsByUserId(userID uint64) ([]response.ProjectNavTreeResponse, error) {

	subjectRes, subjectErr := repository.GetSubjectsByUserId(userID)
	if subjectErr != nil {
		return nil, subjectErr
	}

	subjectNavTreeMap := make(map[uint64][]response.SubjectNavTreeResponse)
	var projectIdSlice []uint64
	for _, subject := range subjectRes {
		if subjectNavTreeMap[subject.ProjectID] == nil {
			projectIdSlice = append(projectIdSlice, subject.ProjectID)
		}
		subjectNavTreeMap[subject.ProjectID] = append(subjectNavTreeMap[subject.ProjectID], subject)
	}

	res, err := repository.GetProjectsBySubjectIds(projectIdSlice)
	if err != nil {
		return nil, err
	}

	for i := range res {
		res[i].Subjects = subjectNavTreeMap[res[i].ID]
	}
	return res, nil
}
