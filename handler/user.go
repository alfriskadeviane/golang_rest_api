package handler

import (
	"net/http"
	"strconv"

	"fmt"
	"mini-project/user"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"

	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var usersResponse []user.UserResponse

	for _, u := range users {
		userResponse := convertToUserResponse(u)
		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": usersResponse,
	})
}

func (h *userHandler) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	// id, _ := uuid.Parse(idString)
	u, err := h.userService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToUserResponse(u)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var userRequest user.UsersRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {

		// log.Fatal(err)
		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		// fmt.Println(err)
		return

	}
	user, err := h.userService.Create(userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToUserResponse(user),
		// "title": bookInput.Title,
		// "price": bookInput.Price,
		// "sub_title": bookInput.SubTitle,
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var userRequest user.UsersRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {

		// log.Fatal(err)
		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		// fmt.Println(err)
		return

	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	// id, _ := uuid.Parse(idString)
	user, err := h.userService.Update(id, userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToUserResponse(user),
		// "title": bookInput.Title,
		// "price": bookInput.Price,
		// "sub_title": bookInput.SubTitle,
	})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	// id, _ := uuid.Parse(idString)

	u, err := h.userService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	userResponse := convertToUserResponse(u)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func convertToUserResponse(u user.User) user.UserResponse {
	return user.UserResponse{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		TempatLahir: u.TempatLahir,
		TglLahir:    u.TglLahir,
	}
}
