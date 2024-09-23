package handler

import (
	"net/http"
	taskmanagement "taskmanager"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(ctx *gin.Context) {
	var task taskmanagement.Tasks

	if err := ctx.BindJSON(&task); err != nil {
		NewError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	h.services.CreateTask(task)
}
func (h *Handler) getTasks(ctx *gin.Context) {

}
func (h *Handler) updateTask(ctx *gin.Context) {

}
func (h *Handler) deleteTask(ctx *gin.Context) {

}
