package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sinyavcev/todoGolang/repository/models"
	"net/http"
)

func (h *Handler) createTask(c *gin.Context) {

	var input models.Todo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Todo.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []models.Todo `json:"data"`
}

func (h *Handler) getAllTask(c *gin.Context) {
	todos, err := h.services.Todo.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: todos,
	})
}

func (h *Handler) getTaskById(c *gin.Context) {
}

func (h *Handler) updateTask(c *gin.Context) {
}

func (h *Handler) deleteTask(c *gin.Context) {

}
