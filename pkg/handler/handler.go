package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	tasks := router.Group("tasks")
	{
		tasks.POST("/")
		tasks.GET("/")
		tasks.PUT("/:id")
		tasks.DELETE("/:id")
	}
	return router
}
