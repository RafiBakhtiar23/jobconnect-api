package routes

import (
	"jobconnect-api/controllers"

	"github.com/gin-gonic/gin"

	"jobconnect-api/middleware"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/register", controllers.Register)

	router.POST("/login", controllers.Login)

	protected := router.Group("/")

	protected.Use(middleware.AuthMiddleware())

	{
		protected.GET("/profile", func(c *gin.Context) {

			c.JSON(200, gin.H{
				"message": "Selamat datang, kamu sudah login",
			})

		})
		protected.GET("/jobs", controllers.GetJobs)

		protected.GET("/jobs/:id", controllers.GetJobByID)

		protected.POST(
			"/jobs",
			middleware.AdminMiddleware(),
			controllers.CreateJob,
		)

		protected.PUT(
			"/jobs/:id",
			middleware.AdminMiddleware(),
			controllers.UpdateJob,
		)

		protected.DELETE(
			"/jobs/:id",
			middleware.AdminMiddleware(),
			controllers.DeleteJob,
		)
	}
}
