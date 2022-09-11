package model

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type FreelancerProfile struct {
	mgm.DefaultModel `bson:",inline"`
	UUID             string                      `bson:"uuid"`
	Details          FreelancerProfileDetails    `json:"details" bson:"details"`
	Skills           []FreelancerProfileSkill    `json:"skills" bson:"skills"`
	Resume           FreelancerProfileResume     `json:"resume" bson:"resume"`
	Languages        []FreelancerProfileLanguage `json:"languages" bson:"languages"`
}

type FreelancerProfileDetails struct {
	FullName  string `json:"full_name" bson:"full_name"`
	Subtitle  string `json:"subtitle" bson:"subtitle"`
	AvatarURL string `json:"avatar_url" bson:"avatar_url"`
}

type SkillSign string

const (
	LessSkillSign SkillSign = "less"
	MoreSkillSign SkillSign = "more"
)

type FreelancerProfileSkill struct {
	Name  string    `json:"name" bson:"name"`
	Years int64     `json:"years" bson:"years"`
	Sign  SkillSign `json:"sign" bson:"sign"`
}

type FreelancerProfileResume struct {
	Biography   string                        `json:"biography" bson:"biography"`
	Experiences []FreelancerProfileExperience `json:"experiences" bson:"experiences"`
}

type FreelancerProfileExperience struct {
	Location    string    `json:"location" bson:"location"`
	Company     string    `json:"company" bson:"company"`
	Type        string    `json:"type" bson:"type"`
	Description string    `json:"description" bson:"description"`
	StartAt     time.Time `json:"start_at" bson:"start_at"`
	EndAt       time.Time `json:"end_at" bson:"end_at"`
	Skills      []string  `json:"skills" bson:"skills"`
}

type FreelancerProfileLanguage struct {
	Name  string `json:"name" bson:"name"`
	Level string `json:"level" bson:"level"`
}

func (a *FreelancerProfile) CollectionName() string {
	return "freelancer_profiles"
}
