package parser_test

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/stretchr/testify/assert"

	"scraping_challenge/services/cometco-scraper/scraper/model"
	"scraping_challenge/services/cometco-scraper/scraper/parser"
)

//go:embed freelancer_profile.html
var freelancerProfileHTML string

func Test_ParseFreelancerProfile(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, freelancerProfileHTML, r.URL.Path)
	}))
	defer ts.Close()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var title string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.Title(&title),
	); err != nil {
		log.Fatal(err)
	}

	expectedData := model.FreelancerProfile{
		Details: model.FreelancerProfileDetails{
			FullName:  "Rollee Test",
			Subtitle:  "Software Engineer at Freelance (Self employed)",
			AvatarURL: "https://profile-pic-comet.s3.eu-central-1.amazonaws.com/FREELANCE/45891/243hv4f1l4wrogsq.jpeg",
		},
		Skills: []model.FreelancerProfileSkill{
			{Name: "Django", Period: "1 an +"},
			{Name: "SQL", Period: "1 an +"},
			{Name: "Flask", Period: "< 1 an"},
			{Name: "Scikit-learn", Period: "< 1 an"},
			{Name: "TensorFlow", Period: "< 1 an"},
		},
		Resume: model.FreelancerProfileResume{
			Biography: "Web Developer for 10 years, I'm currently working on creating the ultimate website.",
			Experiences: []model.FreelancerProfileExperience{
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
		Languages: []model.FreelancerProfileLanguage{
			{
				Name:  "Français",
				Level: "Bilingue / natif",
			},
		},
	}
	data, err := parser.ParseFreelancerProfile(ctx)
	assert.NoError(t, err)
	assert.Equal(t, data, expectedData)
}
