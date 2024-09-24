package handler

import (
	"taskmanager/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	tasks := router.Group("/tasks")
	{
		tasks.POST("", h.createTask)
		tasks.GET("", h.getTasks)
		tasks.PUT("/:id", h.updateTask)
		tasks.DELETE("/:id", h.deleteTask)
	}
	return router
}
