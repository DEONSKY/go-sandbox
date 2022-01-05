package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var err error
var dsn string

type User struct {
	gorm.Model
	Name  string
	Email string
}

type UserDTO struct {
	Name  string
	Email string
}

func AllUsers(c *gin.Context) {

	var users []User

	db.Find(&users)
	c.JSON(200, users)

}

func NewUser(c *gin.Context) {

	fmt.Println("New User")
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("An error Occured")
		return
	}

	fmt.Println("name" + req.Name)
	fmt.Println("email" + req.Email)
	db.Create(&User{Name: req.Name, Email: req.Email})

	c.JSON(201, gin.H{
		"message": "User Created",
	})
}

func DeleteUser(c *gin.Context) {

	name := c.Param("name")

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	c.JSON(200, gin.H{
		"message": "User Deleted",
	})
}

func UpdateUser(c *gin.Context) {

	name := c.Param("name")

	var user User
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "User not Found",
		})
		return
	}
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("An error Occured")
		return
	}
	user.Email = req.Email

	db.Save(&user)

	c.JSON(200, gin.H{
		"message": "User Updated",
	})

}
