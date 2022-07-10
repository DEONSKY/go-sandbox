package service

import (
	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/repository"
	"github.com/DEONSKY/go-sandbox/utils"
	"github.com/mashingan/smapping"
)

func CreateSubject(subjectCreateDTO request.SubjectCreateRequest, user_id uint64) (*model.Subject, error) {
	subjectToCreate := model.Subject{}
	err := smapping.FillStruct(&subjectToCreate, smapping.MapFields(&subjectCreateDTO))
	if err != nil {
		return nil, utils.ReturnErrorResponse(400, "Request DTO parse problem", []string{err.Error()})
	}
	res, err := repository.InsertSubject(subjectToCreate)
	if err != nil {
		return nil, utils.ReturnErrorResponse(422, "Subject could not be inserted", []string{err.Error()})
	}
	return InsertUserToSubject(res.ID, user_id)

}

func InsertUserToSubject(subject_id uint64, user_id uint64) (*model.Subject, error) {
	subject, err := repository.FindSubject(subject_id)
	if err != nil {
		return nil, utils.ReturnErrorResponse(404, "Subject not found", []string{err.Error()})
	}
	user, err := repository.FindUser(user_id)
	if err != nil {
		return nil, utils.ReturnErrorResponse(404, "User not found", []string{err.Error()})
	}
	res, err := repository.InsertUserToSubject(*subject, *user)
	if err != nil {
		return nil, utils.ReturnErrorResponse(409, "Failed create association between subject and user", []string{err.Error()})
	}
	return res, err
}
