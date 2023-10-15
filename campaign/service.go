package campaign

import (
	"errors"
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(InputID GetCampaignDetailInput, InputData CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserId(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserId = input.User.ID
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)
	campaign.CreatedDate = time.Now()
	campaign.UpdatedDate = time.Now()

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return campaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(InputID GetCampaignDetailInput, InputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(InputID.ID)
	if err != nil {
		return campaign, err
	}
	campaign.Name = InputData.Name
	campaign.Description = InputData.Description
	campaign.ShortDescription = InputData.ShortDescription
	campaign.Perks = InputData.Perks
	campaign.GoalAmount = InputData.GoalAmount
	campaign.UpdatedDate = time.Now()

	if campaign.UserId != InputData.User.ID {
		return campaign, errors.New("not an owner that campaign")
	}

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return campaign, err
	}

	return updateCampaign, nil
}
