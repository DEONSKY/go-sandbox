package repository

import (
	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/model"
)

func InsertSubject(subject model.Subject) (*model.Subject, error) {
	if result := config.DB.Save(&subject); result.Error != nil {
		return nil, result.Error
	}

	return &subject, nil
}

func InsertUserToSubject(subject model.Subject, user model.User) model.Subject {
	config.DB.Model(&subject).Association("User").Append(&user)
	return subject
}

func FindSubject(subjet_id uint64) model.Subject {
	var subject model.Subject
	config.DB.First(&subject, subjet_id)
	return subject
}
