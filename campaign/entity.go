package campaign

import "time"

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
	Updateddate      time.Time
	CampaignImages   []CampaignImages
}

type CampaignImages struct {
	ID          int
	CampaignId  int
	FileName    string
	IsPrimary   int
	CreatedDate time.Time
	Updateddate time.Time
}
