package service

import (
	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/dto/response"
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
	if err := ControlAccessToSubject(subjectCreateDTO.ProjectID, user_id); err != nil {
		return nil, err
	}
	res, err := repository.InsertSubject(subjectToCreate)
	if err != nil {
		return nil, utils.ReturnErrorResponse(422, "Subject could not be inserted", []string{err.Error()})
	}
	return InsertUserToSubject(res, user_id)

}

func InsertUserToSubjectIfAllowed(subjectID uint64, targetUserID uint64, userID uint64) (*model.Subject, error) {
	subject, err := repository.FindSubject(subjectID)
	if err != nil {
		return nil, utils.ReturnErrorResponse(404, "Subject not found", []string{err.Error()})
	}
	if err := ControlAccessToSubject(subject.ProjectID, userID); err != nil {
		return nil, err
	}
	return InsertUserToSubject(subject, targetUserID)
}

func GetSubjectsUsersOptions(subjectID uint64, userID uint64) ([]response.UserOptionResponse, error) {
	options, err := repository.GetSubjectsUsersOptions(subjectID)
	if err != nil {
		return nil, utils.ReturnErrorResponse(500, "Unexpected Error", []string{err.Error()})
	}
	if !(len(options) > 0) {
		return options, err
	}
	hasAccess := false
	for _, option := range options {
		if option.ID == userID {
			hasAccess = true
		}
	}
	if hasAccess {
		return options, err
	}
	return nil, utils.ReturnErrorResponse(403, "User has no access to subject", []string{})
}

func ControlAccessToSubject(projectID uint64, userID uint64) error {
	resBool, err := repository.ProjectExistsByIDAndLeaderID(projectID, userID)
	if err != nil {
		return utils.ReturnErrorResponse(500, "Unexpected Error", []string{err.Error()})
	}
	if !resBool {
		return utils.ReturnErrorResponse(403, "You must be project leader of target project for adding subject", []string{})
	}
	return nil
}

func InsertUserToSubject(subject *model.Subject, user_id uint64) (*model.Subject, error) {
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
