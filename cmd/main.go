package main

import (
	"jobconnect-api/config"
	"jobconnect-api/routes"


	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.Migration()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
	
}