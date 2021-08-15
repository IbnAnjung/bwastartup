package transaction

import "time"

type CampaignTrasactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
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
