package service

import (
	"log"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/repository"
	"github.com/mashingan/smapping"
)

func CreateSubject(subjectCreateDTO request.SubjectCreateRequest) (*model.Subject, error) {
	subjectToCreate := model.Subject{}
	log.Println(subjectCreateDTO)
	err := smapping.FillStruct(&subjectToCreate, smapping.MapFields(&subjectCreateDTO))
	if err != nil {
		return nil, err
	}
	log.Println(subjectToCreate)
	res, err := repository.InsertSubject(subjectToCreate)
	return res, err
}

func InsertUserToSubject(subject_id uint64, user_id uint64) (*model.Subject, error) {
	subject, err := repository.FindSubject(subject_id)
	if err != nil {
		return nil, err
	}
	user, err := repository.FindUser(user_id)
	if err != nil {
		return nil, err
	}
	log.Println(user)
	res, err := repository.InsertUserToSubject(*subject, *user)
	return res, err
}
