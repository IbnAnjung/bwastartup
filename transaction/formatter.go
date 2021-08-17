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

type CampaignTransactionFormat struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
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

func FormatTransaction(transaction Transaction) CampaignTransactionFormat {
	formattedCampaignTransaction := CampaignTransactionFormat{}

	formattedCampaignTransaction.ID = transaction.ID
	formattedCampaignTransaction.CampaignID = transaction.CampaignID
	formattedCampaignTransaction.UserID = transaction.UserID
	formattedCampaignTransaction.Amount = transaction.Amount
	formattedCampaignTransaction.Status = transaction.Status
	formattedCampaignTransaction.Code = transaction.Code
	formattedCampaignTransaction.PaymentURL = transaction.PaymentURL

	return formattedCampaignTransaction

}
