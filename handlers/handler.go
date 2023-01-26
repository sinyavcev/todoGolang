package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sinyavcev/todoGolang/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	todos := router.Group("/todos")
	{
		todos.POST("/", h.createTask)
		todos.GET("/", h.getAllTask)
		todos.GET("/:id", h.getTaskById)
		todos.PUT("/:id", h.updateTask)
		todos.DELETE("/:id", h.deleteTask)
	}

	return router
}
