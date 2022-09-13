package model

import (
	"time"
)

type CometFreelancerProfile struct {
	ID        string                           `json:"id"`
	Details   CometFreelancerProfileDetails    `json:"details"`
	Skills    []CometFreelancerProfileSkill    `json:"skills"`
	Resume    CometFreelancerProfileResume     `json:"resume"`
	Languages []CometFreelancerProfileLanguage `json:"languages"`
}

type CometFreelancerProfileDetails struct {
	FullName  string `json:"full_name"`
	Subtitle  string `json:"subtitle"`
	AvatarURL string `json:"avatar_url"`
}

type CometSkillSign string

const (
	LessSkillSign CometSkillSign = "less"
	MoreSkillSign CometSkillSign = "more"
)

type CometFreelancerProfileSkill struct {
	Name  string         `json:"name"`
	Years int64          `json:"years"`
	Sign  CometSkillSign `json:"sign"`
}

type CometFreelancerProfileResume struct {
	Biography   string                             `json:"biography"`
	Experiences []CometFreelancerProfileExperience `json:"experiences"`
}

type CometFreelancerProfileExperience struct {
	Location    string    `json:"location"`
	Company     string    `json:"company"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	Skills      []string  `json:"skills"`
}

type CometFreelancerProfileLanguage struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}
