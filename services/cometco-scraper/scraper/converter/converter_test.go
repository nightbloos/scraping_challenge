package converter_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"scraping_challenge/services/cometco-scraper/domain/model"
	"scraping_challenge/services/cometco-scraper/scraper/converter"
	cometModel "scraping_challenge/services/cometco-scraper/scraper/model"
)

func Test_FromCometFreelancerProfileSkill(t *testing.T) {
	tests := []struct {
		name                           string
		cometFreelancerProfileSkill    cometModel.FreelancerProfileSkill
		expectedFreelancerProfileSkill model.FreelancerProfileSkill
	}{
		{
			name: "LESS_THAN_1_YEAR",
			cometFreelancerProfileSkill: cometModel.FreelancerProfileSkill{
				Name:   "Flask",
				Period: "< 1 an",
			},
			expectedFreelancerProfileSkill: model.FreelancerProfileSkill{
				Name:  "Flask",
				Years: 1,
				Sign:  model.LessSkillSign,
			},
		}, {
			name: "MORE_THAN_1_YEAR",
			cometFreelancerProfileSkill: cometModel.FreelancerProfileSkill{
				Name:   "SQL",
				Period: "1 an +",
			},
			expectedFreelancerProfileSkill: model.FreelancerProfileSkill{
				Name:  "SQL",
				Years: 1,
				Sign:  model.MoreSkillSign,
			},
		}, {
			name: "MORE_THAN_2_YEARS",
			cometFreelancerProfileSkill: cometModel.FreelancerProfileSkill{
				Name:   "SQL",
				Period: "2 ans +",
			},
			expectedFreelancerProfileSkill: model.FreelancerProfileSkill{
				Name:  "SQL",
				Years: 2,
				Sign:  model.MoreSkillSign,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualFreelancerProfileSkill, err := converter.FromCometFreelancerProfileSkill(tt.cometFreelancerProfileSkill)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedFreelancerProfileSkill, actualFreelancerProfileSkill)
		})
	}
}

func Test_FromCometFreelancerProfileLanguage(t *testing.T) {
	tests := []struct {
		name                              string
		cometFreelancerProfileLanguage    cometModel.FreelancerProfileLanguage
		expectedFreelancerProfileLanguage model.FreelancerProfileLanguage
	}{
		{
			name: "FRENCH_NATIVE",
			cometFreelancerProfileLanguage: cometModel.FreelancerProfileLanguage{
				Name:  "Français",
				Level: "Bilingue / natif",
			},
			expectedFreelancerProfileLanguage: model.FreelancerProfileLanguage{
				Name:  "Français",
				Level: "Bilingue / natif",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualFreelancerProfileLanguage := converter.FromCometFreelancerProfileLanguage(tt.cometFreelancerProfileLanguage)
			assert.Equal(t, tt.expectedFreelancerProfileLanguage, actualFreelancerProfileLanguage)
		})
	}
}

func Test_FromCometFreelancerProfileDetails(t *testing.T) {
	tests := []struct {
		name                             string
		cometFreelancerProfileDetails    cometModel.FreelancerProfileDetails
		expectedFreelancerProfileDetails model.FreelancerProfileDetails
	}{
		{
			name: "FULL_DETAILS",
			cometFreelancerProfileDetails: cometModel.FreelancerProfileDetails{
				FullName:  "Rollee Test",
				Subtitle:  "Software Engineer at Freelance (Self employed)",
				AvatarURL: "https://profile-pic-comet.s3.eu-central-1.amazonaws.com/FREELANCE/45891/243hv4f1l4wrogsq.jpeg",
			},
			expectedFreelancerProfileDetails: model.FreelancerProfileDetails{
				FullName:  "Rollee Test",
				Subtitle:  "Software Engineer at Freelance (Self employed)",
				AvatarURL: "https://profile-pic-comet.s3.eu-central-1.amazonaws.com/FREELANCE/45891/243hv4f1l4wrogsq.jpeg",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualFreelancerProfileDetails := converter.FromCometFreelancerProfileDetails(tt.cometFreelancerProfileDetails)
			assert.Equal(t, tt.expectedFreelancerProfileDetails, actualFreelancerProfileDetails)
		})
	}
}

func Test_FromCometFreelancerProfileExperience(t *testing.T) {
	tests := []struct {
		name                                string
		cometFreelancerProfileExperience    cometModel.FreelancerProfileExperience
		expectedFreelancerProfileExperience model.FreelancerProfileExperience
	}{
		{
			name: "WITH_END_DATE_IN_THE_PAST",
			cometFreelancerProfileExperience: cometModel.FreelancerProfileExperience{
				Location:    "",
				Company:     "Freelance",
				Type:        "Freelancing",
				Description: "Worked on Sql",
				StartAt:     "janv. 2020",
				EndAt:       "déc. 2020",
				Skills: []string{
					"SQL",
				},
			},
			expectedFreelancerProfileExperience: model.FreelancerProfileExperience{
				Location:    "",
				Company:     "Freelance",
				Type:        "Freelancing",
				Description: "Worked on Sql",
				StartAt:     time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndAt:       time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC),
				Skills: []string{
					"SQL",
				},
			},
		},
		{
			name: "WITH_END_DATE_ONGOING",
			cometFreelancerProfileExperience: cometModel.FreelancerProfileExperience{
				Location:    "Here",
				Company:     "My Company",
				Type:        "Prestation ESN / SSII",
				Description: "Worked hard",
				StartAt:     "août 2020",
				EndAt:       "En cours",
				Skills: []string{
					"Flask",
					"Scikit-learn",
					"TensorFlow",
				},
			},
			expectedFreelancerProfileExperience: model.FreelancerProfileExperience{
				Location:    "Here",
				Company:     "My Company",
				Type:        "Prestation ESN / SSII",
				Description: "Worked hard",
				StartAt:     time.Date(2020, time.August, 1, 0, 0, 0, 0, time.UTC),
				EndAt:       time.Time{},
				Skills: []string{
					"Flask",
					"Scikit-learn",
					"TensorFlow",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualFreelancerProfileExperience, err := converter.FromCometFreelancerProfileExperience(tt.cometFreelancerProfileExperience)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedFreelancerProfileExperience, actualFreelancerProfileExperience)
		})
	}
}

func Test_FromCometFreelancerProfile(t *testing.T) {
	cometProfile := cometModel.FreelancerProfile{
		Details: cometModel.FreelancerProfileDetails{
			FullName:  "Rollee Test",
			Subtitle:  "Software Engineer at Freelance (Self employed)",
			AvatarURL: "https://profile-pic-comet.s3.eu-central-1.amazonaws.com/FREELANCE/45891/243hv4f1l4wrogsq.jpeg",
		},
		Skills: []cometModel.FreelancerProfileSkill{
			{Name: "Django", Period: "1 an +"},
			{Name: "SQL", Period: "1 an +"},
			{Name: "Flask", Period: "< 1 an"},
			{Name: "Scikit-learn", Period: "< 1 an"},
			{Name: "TensorFlow", Period: "< 1 an"},
		},
		Resume: cometModel.FreelancerProfileResume{
			Biography: "Web Developer for 10 years, I'm currently working on creating the ultimate website.",
			Experiences: []cometModel.FreelancerProfileExperience{
				{
					Location:    "Here",
					Company:     "My Company",
					Type:        "Prestation ESN / SSII",
					Description: "Worked hard",
					StartAt:     "sept. 2022",
					EndAt:       "En cours",
					Skills: []string{
						"Flask",
						"Scikit-learn",
						"TensorFlow",
					},
				},
				{
					Location:    "",
					Company:     "Freelance",
					Type:        "Freelancing",
					Description: "Worked on Sql",
					StartAt:     "janv. 2020",
					EndAt:       "déc. 2020",
					Skills: []string{
						"SQL",
					},
				},
				{
					Location:    "",
					Company:     "Freelance",
					Type:        "Freelancing",
					Description: "Software Engineer",
					StartAt:     "janv. 2019",
					EndAt:       "déc. 2019",
					Skills:      []string{"Django"},
				},
			},
		},
		Languages: []cometModel.FreelancerProfileLanguage{
			{
				Name:  "Français",
				Level: "Bilingue / natif",
			},
		},
	}

	expectedProfile := model.FreelancerProfile{
		Details: model.FreelancerProfileDetails{
			FullName:  "Rollee Test",
			Subtitle:  "Software Engineer at Freelance (Self employed)",
			AvatarURL: "https://profile-pic-comet.s3.eu-central-1.amazonaws.com/FREELANCE/45891/243hv4f1l4wrogsq.jpeg",
		},
		Skills: []model.FreelancerProfileSkill{
			{
				Name:  "Django",
				Years: 1,
				Sign:  model.MoreSkillSign,
			},
			{
				Name:  "SQL",
				Years: 1,
				Sign:  model.MoreSkillSign,
			},
			{
				Name:  "Flask",
				Years: 1,
				Sign:  model.LessSkillSign,
			},
			{
				Name:  "Scikit-learn",
				Years: 1,
				Sign:  model.LessSkillSign,
			},
			{
				Name:  "TensorFlow",
				Years: 1,
				Sign:  model.LessSkillSign,
			},
		},
		Resume: model.FreelancerProfileResume{
			Biography: "Web Developer for 10 years, I'm currently working on creating the ultimate website.",
			Experiences: []model.FreelancerProfileExperience{
				{
					Location:    "Here",
					Company:     "My Company",
					Type:        "Prestation ESN / SSII",
					Description: "Worked hard",
					StartAt:     time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC),
					EndAt:       time.Time{},
					Skills:      []string{"Flask", "Scikit-learn", "TensorFlow"},
				},
				{
					Location:    "",
					Company:     "Freelance",
					Type:        "Freelancing",
					Description: "Worked on Sql",
					StartAt:     time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:       time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC),
					Skills:      []string{"SQL"},
				},
				{
					Location:    "",
					Company:     "Freelance",
					Type:        "Freelancing",
					Description: "Software Engineer",
					StartAt:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:       time.Date(2019, time.December, 1, 0, 0, 0, 0, time.UTC),
					Skills:      []string{"Django"},
				},
			},
		},
		Languages: []model.FreelancerProfileLanguage{
			{
				Name:  "Français",
				Level: "Bilingue / natif",
			},
		},
	}

	actualProfile, err := converter.FromCometFreelancerProfile(cometProfile)
	assert.NoError(t, err)
	assert.Equal(t, expectedProfile, actualProfile)
}
