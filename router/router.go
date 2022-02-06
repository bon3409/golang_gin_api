package router

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!!##",
		})
	})

	api := router.Group("api")
	{
		api.GET("/users", controllers.GetAllUsers)
		api.GET("/user/:id", controllers.GetUserById)
		api.POST("/user/create", controllers.CreateUser)
	}

	return router
}
