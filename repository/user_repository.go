package repository

import (
	"log"

	"github.com/DEONSKY/go-sandbox/config"
	"github.com/DEONSKY/go-sandbox/dto/response"
	"github.com/DEONSKY/go-sandbox/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InsertUser(user model.User) model.User {
	user.Password = hashAndSalt([]byte(user.Password))
	config.DB.Save(&user)
	return user
}

func UpdateUser(user model.User) model.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser model.User
		config.DB.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	config.DB.Save(&user)
	return user
}

func VerifyCredential(email string, password string) interface{} {
	var user model.User
	res := config.DB.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user model.User
	return config.DB.Where("email = ?", email).Take(&user)
}

func FindByEmail(email string) model.User {
	var user model.User
	config.DB.Where("email = ?", email).Take(&user)
	return user
}

func ProfileUser(userID string) model.User {
	var user model.User
	config.DB.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func FindUser(user_id uint64) (*model.User, error) {
	var user model.User
	if result := config.DB.First(&user, user_id); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetSubjectsUsersOptions(subjectID uint64) ([]response.UserOptionResponse, error) {
	var userOptions []response.UserOptionResponse
	if result := config.DB.Model(model.User{}).
		Joins("INNER JOIN subject_users su on su.user_id = id").
		Where("subject_id", subjectID).Find(&userOptions); result.Error != nil {
		return nil, result.Error
	}
	return userOptions, nil
}
