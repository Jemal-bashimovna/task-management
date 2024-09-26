package handler

import (
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewError(ctx, http.StatusBadRequest, "Task not found")
		return
	}

	var task taskmanagement.UpdateTaskInput
	if err := ctx.BindJSON(&task); err != nil {
		NewError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateTask(id, task)
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "task updated",
	})

}
func (h *Handler) deleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewError(ctx, http.StatusBadRequest, "Task not found")
		return
	}

	err = h.services.Tasks.DeleteTask(id)
	if err != nil {
		NewError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "task deleted",
	})
}
