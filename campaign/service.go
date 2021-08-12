package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s service) FindCampaigns(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	var err error

	if userID == 0 {
		campaigns, err = s.repository.FindAll()
	} else {
		campaigns, err = s.repository.FindByUserID(userID)
	}

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
