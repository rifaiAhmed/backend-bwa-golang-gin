package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetcampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to get campaign's transactions param's input", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("successfully to get data transactions by id", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("failed to get user's transactions", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("successfully to get data transactions by id", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		if err != nil {
			response := helper.APIResponse("failed to create transaction param's input", http.StatusUnprocessableEntity, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("failed to create transaction", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("successfully to cretae transaction", http.StatusOK, "success", transaction.FormaterTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		if err != nil {
			response := helper.APIResponse("failed to create transaction param's input", http.StatusUnprocessableEntity, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	err = h.service.ProcessPayment(input)
	if err != nil {
		if err != nil {
			response := helper.APIResponse("failed to process payment", http.StatusUnprocessableEntity, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	c.JSON(http.StatusOK, input)
}
