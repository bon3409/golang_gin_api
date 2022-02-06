package repositories

import (
	"app/config"
	"app/models"
	"fmt"
	"log"
)

func GetAllUsers() []*models.User {
	fmt.Println("get all users")
	var users []*models.User
	config.DB.Find(&users)
	return users
}

func GetUserById(id string) []*models.User {
	var user []*models.User
	config.DB.First(&user, id)
	return user
}

func CreateUser(user *models.User) bool {
	if result := config.DB.Create(&user); result.Error != nil {
		log.Fatalf("Create user error: %v", result.Error)
		return false
	}

	return true
}
