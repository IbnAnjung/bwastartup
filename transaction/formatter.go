package transaction

import (
	"time"
)

type CampaignTrasactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionsFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTrasactionsFormatter {
	var formattedCampaignTransaction CampaignTrasactionsFormatter

	formattedCampaignTransaction.ID = transaction.ID
	formattedCampaignTransaction.Name = transaction.User.Name
	formattedCampaignTransaction.Amount = transaction.Amount
	formattedCampaignTransaction.CreatedAt = transaction.CreatedAt

	return formattedCampaignTransaction

}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTrasactionsFormatter {
	var formattedCampaignTransactions []CampaignTrasactionsFormatter

	for _, transaction := range transactions {
		formattedCampaignTransaction := FormatCampaignTransaction(transaction)

		formattedCampaignTransactions = append(formattedCampaignTransactions, formattedCampaignTransaction)
	}

	return formattedCampaignTransactions
}

func FormatUserTransaction(transaction Transaction) UserTransactionsFormatter {
	formattedUserTransaction := UserTransactionsFormatter{}
	formattedCampaign := CampaignFormatter{}

	formattedUserTransaction.ID = transaction.ID
	formattedUserTransaction.Amount = transaction.Amount
	formattedUserTransaction.CreatedAt = transaction.CreatedAt

	formattedCampaign.Name = transaction.Campaign.Name
	formattedCampaign.ImageURL = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		formattedCampaign.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formattedUserTransaction.Campaign = formattedCampaign

	return formattedUserTransaction
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionsFormatter {
	formattedUserTransactions := []UserTransactionsFormatter{}

	for _, transaction := range transactions {
		formattedUserTransaction := FormatUserTransaction(transaction)
		formattedUserTransactions = append(formattedUserTransactions, formattedUserTransaction)
	}

	return formattedUserTransactions

}
