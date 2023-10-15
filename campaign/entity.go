package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	ID               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	BackerCount      int
	CurrentAmount    int
	Perks            string
	Slug             string
	CreatedDate      time.Time
	UpdatedDate      time.Time
	CampaignImages   []CampaignImages
	User             user.User
}

type CampaignImages struct {
	ID          int
	CampaignId  int
	FileName    string
	IsPrimary   int
	CreatedDate time.Time
	UpdatedDate time.Time
}
