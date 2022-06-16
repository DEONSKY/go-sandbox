package repository

import (
	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/dto/response"
	"github.com/DEONSKY/go-sandbox/model"
)

func InsertSubject(subject model.Subject) (*model.Subject, error) {
	if result := config.DB.Save(&subject); result.Error != nil {
		return nil, result.Error
	}

	return &subject, nil
}

func InsertUserToSubject(subject model.Subject, user model.User) (*model.Subject, error) {
	if err := config.DB.Model(&subject).Association("User").Append(&user); err != nil {
		return nil, err
	}
	return &subject, nil
}

func FindSubject(subjet_id uint64) (*model.Subject, error) {
	var subject model.Subject
	if result := config.DB.First(&subject, subjet_id); result.Error != nil {
		return nil, result.Error
	}
	return &subject, nil
}

func GetSubjectsByUserId(userID uint64) ([]response.SubjectNavTreeResponse, error) {
	var subjectNavTreeResponse []response.SubjectNavTreeResponse
	if result := config.DB.Model(&model.Subject{}).
		Joins("INNER JOIN subject_users su on su.subject_id = id").
		Where("su.user_id", userID).Order("ID").Find(&subjectNavTreeResponse); result.Error != nil {
		return nil, result.Error
	}
	return subjectNavTreeResponse, nil
}
