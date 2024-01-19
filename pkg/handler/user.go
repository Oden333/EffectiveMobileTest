package handler

import (
	"EMtest/models"
	helpers "EMtest/pkg/handler/helper"
	"fmt"
	"math"
	"net/http"
	"strconv"

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
	logrus.Info("Got request creating", "User", input)

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

func (h *Handler) get_all_users(c *gin.Context) {

	//Узнаём номер страницы + данные для пагинации
	page := 1
	pageNum := c.Param("page")
	if pageNum != "" {
		page, _ = strconv.Atoi(pageNum)
	}
	//Узнаём параметры для страницы

	limit := 5
	offset := (page - 1) * limit

	// Получаем параметры фильтрации из запроса
	queryParamsData := []string{"name", "surname", "patronymic", "age", "gender", "country"}
	filters := map[string]string{}
	for _, param := range queryParamsData {
		filter := c.Query(param)
		if filter != "" {
			filters[param] = filter
		}
	}
	fmt.Println(page, pageNum, offset)

	var count int
	var people []models.User
	var err error
	if len(filters) == 0 {

		//Достаём людей, в случае отсутствия фильтров
		count, people, err = h.services.GetAllUsers(limit, offset)
		if err != nil {
			logrus.Debug("Error while requesting all users from DB", err)
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {

		logrus.Info("Got filtered request")

		//Достаём необходимых людей по фильтру
		count, people, err = h.services.GetCertainUsers(limit, offset, filters)
		if err != nil {
			logrus.Debug("Error while requesting all users from DB", err)
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// Параметр для корректного отображения крайней страницы
	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	/*

		// если есть остаток
		   	if count%limit > 0 {
		   		totalPages++
		   	}

	*/
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people": people,
		"pagination": helpers.PaginationData{
			NextPage:   page + 1,
			PrevPage:   page - 1,
			CurrPage:   page,
			TotalPages: int(totalPages),
		},
	})
}

func (h *Handler) delete_user_by_id(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.DeleteUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error while deleting user")
	}
}

func (h *Handler) update_user_by_id(c *gin.Context) {
	//ID изменяемого пользвателя из урла
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	// Json Данные запроса для изменения
	var input helpers.UserData
	if err := c.BindJSON(&input); err != nil {
		logrus.Debug("error while unmarshalling request", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Передаём в сервис
	err = h.services.UserService.UpdateUser(id, input)
	if err != nil {
		logrus.Debug("error while editing user", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}
