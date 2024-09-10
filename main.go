package main

import (
	"os"

	"example.com/gin-api/configs"
	"example.com/gin-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.Run(":" + os.Getenv("GO_PORT"))
}

func SetupRouter() *gin.Engine {
	// connect db
	configs.Connection()

	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	routes.InitHomeRoutes(apiV1)
	routes.InitUserRoutes(apiV1)

	return router
}
