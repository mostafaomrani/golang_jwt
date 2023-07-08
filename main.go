package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jwt/controllers"
	"jwt/initializers"
	"jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConncetToDb()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/validate", middleware.RequireAuth ,controllers.Validate)

	r.Run()
}
