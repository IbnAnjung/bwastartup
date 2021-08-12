package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserID           int    `json:"user_id"`
	Slug             string `json:"slug"`
}

type CampaignDetailFormatter struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	ImageURL         string                    `json:"image_url"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"curret_amount"`
	UserID           int                       `json:"user_id"`
	Slug             string                    `json:"slug"`
	Perks            []string                  `json:"perks"`
	User             CampaignUserFormatter     `json:"user"`
	Images           []CampaignImagesFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}

	campaignFormatter.ID = campaign.ID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	formattedCampaigns := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		formattedCampaigns = append(formattedCampaigns, campaignFormatter)
	}

	return formattedCampaigns
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	formattedCampaign := CampaignDetailFormatter{}

	formattedCampaign.ID = campaign.ID
	formattedCampaign.Name = campaign.Name
	formattedCampaign.ShortDescription = campaign.ShortDescription
	formattedCampaign.Description = campaign.Description
	formattedCampaign.GoalAmount = campaign.GoalAmount
	formattedCampaign.CurrentAmount = campaign.CurrentAmount
	formattedCampaign.UserID = campaign.UserID
	formattedCampaign.Slug = campaign.Slug

	formattedCampaign.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		formattedCampaign.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	formattedCampaign.Perks = perks

	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.ID = campaign.User.ID
	campaignUserFormatter.ImageURL = campaign.User.AvatarFileName

	formattedCampaign.User = campaignUserFormatter

	campaignImagesFormatter := []CampaignImagesFormatter{}

	for _, campaignImage := range campaign.CampaignImages {
		formattedCampaignImage := CampaignImagesFormatter{}
		formattedCampaignImage.ImageURL = campaignImage.FileName
		formattedCampaignImage.IsPrimary = false
		if campaignImage.IsPrimary == 1 {
			formattedCampaignImage.IsPrimary = true
		}
		campaignImagesFormatter = append(campaignImagesFormatter, formattedCampaignImage)
	}

	formattedCampaign.Images = campaignImagesFormatter

	return formattedCampaign
}
