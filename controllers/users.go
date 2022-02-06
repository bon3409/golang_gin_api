package controllers

import (
	"app/models"
	"app/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.User{}

	c.BindJSON(&user)

	result := repositories.CreateUser(&user)

	if result {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "failed",
		})
	}
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user := repositories.GetUserById(id)
	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	users := repositories.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
