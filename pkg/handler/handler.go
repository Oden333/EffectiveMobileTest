package handler

import (
	"EMtest/pkg/service"

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

	auth := router.Group("/auth")
	{
		// Метод для добавления новых людей в формате
		auth.POST("/add", h.create_user)
	}

	return router
}
