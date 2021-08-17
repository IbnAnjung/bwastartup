package payment

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
	campaignRepository campaign.Repository
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService(campaignRepository campaign.Repository) *service {
	return &service{campaignRepository}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midClient := midtrans.NewClient()

	midClient.ServerKey = ""
	midClient.ClientKey = ""
	midClient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midClient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
