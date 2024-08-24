package main

import (
	"backend/controllers"
	"backend/initializers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}
