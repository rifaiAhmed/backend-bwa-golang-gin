package transaction

import "time"

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
