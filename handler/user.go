package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	NewUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "succes", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(NewUser, "token token")
	response := helper.APIResponse("Account Has Been Created", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("login failed first", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	LoggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(LoggedInUser, "tokentokentoken")
	response := helper.APIResponse("Login Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	isEmailAvailable, err := h.userService.IsEmailAvalability(input)
	if err != nil {
		errorMessage := gin.H{"errors": "server error"}
		response := helper.APIResponse("Email is Available", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{
		"is_available": isEmailAvailable,
	}
	var metMessage string
	if isEmailAvailable {
		metMessage = "Email Is Available"
	} else {
		metMessage = "Email Has Been Registered"
	}
	response := helper.APIResponse(metMessage, http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)
}
