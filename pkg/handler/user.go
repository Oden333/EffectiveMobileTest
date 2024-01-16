package handler

import (
	"EMtest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) create_user(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.UserService.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id})
}
