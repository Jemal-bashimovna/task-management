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

	id, err := h.services.Tasks.CreateTask(task)
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getTasks(ctx *gin.Context) {
	tasks, err := h.services.Tasks.GetTasks()
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}
func (h *Handler) updateTask(ctx *gin.Context) {

}
func (h *Handler) deleteTask(ctx *gin.Context) {

}
