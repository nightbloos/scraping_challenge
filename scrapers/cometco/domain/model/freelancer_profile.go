package model

type FreelancerProfile struct {
	Details   FreelancerProfileDetails
	Skills    []FreelancerProfileSkill
	Resume    FreelancerProfileResume
	Languages []FreelancerProfileLanguage
}

type FreelancerProfileDetails struct {
	FullName  string
	Subtitle  string
	AvatarURL string
}

type FreelancerProfileSkill struct {
	Name   string
	Period string
}

type FreelancerProfileResume struct {
	Biography   string
	Experiences []FreelancerProfileExperience
}

type FreelancerProfileExperience struct {
	Location    string
	Company     string
	Type        string
	Description string
	StartAt     string
	EndAt       string
	Skills      []string
}

type FreelancerProfileLanguage struct {
	Name  string
	Level string
}
