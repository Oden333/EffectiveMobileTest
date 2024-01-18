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
	//Конфигурация страницы ответа
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")

	//Конфигурация рутов
	auth := router.Group("/people")
	{
		// Метод для добавления новых людей в формате
		auth.POST("/add", h.create_user)
		// Метод для получения данных с различными фильтрами и пагинацией
		auth.GET("/page/:page", h.get_all_users)
		// Метод для изменения сущности
		auth.POST("/edit/:id")
		// Метод для удаления сущности
		auth.DELETE("/delete/:id", h.delete_user_by_id)

	}
	router.Static("assets", "./assets")

	return router
}
