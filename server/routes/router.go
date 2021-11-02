package routes

import (
	"github.com/daniel0liver/books-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.List)
			books.GET("/:id", controllers.FindById)
			books.POST("/", controllers.Create)
			books.PUT("/", controllers.Update)
			books.DELETE("/", controllers.Delete)
		}
	}

	return router
}
