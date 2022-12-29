package main

import (
	"github.com/JulesAD96/go-jwt-auth/controllers"
	"github.com/JulesAD96/go-jwt-auth/database"
	"github.com/JulesAD96/go-jwt-auth/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("root:root@tcp(localhost:3306)/jwtauth?parseTime=true")
	database.Migrate()

	// Initialize router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/register", controllers.Register)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
