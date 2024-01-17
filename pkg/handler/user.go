package handler

import (
	"EMtest/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) create_user(c *gin.Context) {

	var input models.User
	if err := c.BindJSON(&input); err != nil {
		logrus.Debug("error while unmarshalling request", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Name == "" {
		logrus.Debug("error: empty name in request")
	}
	if input.Patronymic == "" {
		input.Patronymic = "-"
	}
	logrus.Info("Got request", "User", input)

	id, err := h.services.UserService.CreateUser(input)
	if err != nil {
		logrus.Debug("error while creating user", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Info("User created, got id", "Id", id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id})
}
