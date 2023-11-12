package transaction

import (
	"time"
)

type CampaignTransactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_ata"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionsFormatter {
	transactionFormater := CampaignTransactionsFormatter{}
	transactionFormater.ID = transaction.ID
	transactionFormater.Name = transaction.User.Name
	transactionFormater.Amount = transaction.Amount
	transactionFormater.CreatedAt = transaction.CreatedAt

	return transactionFormater
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionsFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionsFormatter{}
	}
	var transactionsFormatter []CampaignTransactionsFormatter
	for _, transaction := range transactions {
		formatterTransaction := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatterTransaction)
	}
	return transactionsFormatter
}

type FormatTransactionByUserID struct {
	ID        int                          `json:"id"`
	Amount    int                          `json:"amount"`
	Status    string                       `json:"status"`
	CreatedAt time.Time                    `json:"created_at"`
	Campaign  FormatCampaignForTransaction `json:"campaign"`
}

type FormatCampaignForTransaction struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormaterUserTransactionByUserID(transaction Transaction) FormatTransactionByUserID {
	formatter := FormatTransactionByUserID{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status

	formatCampaign := FormatCampaignForTransaction{}
	formatCampaign.Name = transaction.Campaign.Name
	formatCampaign.ImageUrl = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		formatCampaign.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}
	formatter.Campaign = formatCampaign

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []FormatTransactionByUserID {
	if len(transactions) == 0 {
		return []FormatTransactionByUserID{}
	}
	var transactionsFormatter []FormatTransactionByUserID
	for _, transaction := range transactions {
		formatterTransaction := FormaterUserTransactionByUserID(transaction)
		transactionsFormatter = append(transactionsFormatter, formatterTransaction)
	}
	return transactionsFormatter
}
