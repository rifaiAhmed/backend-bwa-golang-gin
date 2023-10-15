package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *CampaignHandler {
	return &CampaignHandler{service}
}

func (h *CampaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)

	if err != nil {
		response := helper.APIResponse("Error to get Campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("successfully to get Campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}
	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("detail campaign", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)

	if err != nil {
		response := helper.APIResponse("failed to create campaign", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("success to create campaign", http.StatusUnprocessableEntity, "error", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("update campaign", http.StatusOK, "success", campaign.FormatCampaignDetail(updatedCampaign))
	c.JSON(http.StatusOK, response)
}
