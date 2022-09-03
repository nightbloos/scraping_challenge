package factory

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"

	"scraping_challenge/scrapers/cometco/domain/model"
)

func NewFreelancerProfile(ctx context.Context) (model.FreelancerProfile, error) {
	details, err := GetFreelancerProfileDetails(ctx)
	if err != nil {
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to create freelancer profile details")
	}
	languages, err := GetFreelancerLanguages(ctx)
	if err != nil {
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to create freelancer profile languages")
	}
	skills, err := GetFreelancerSkills(ctx)
	if err != nil {
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to create freelancer profile skills")
	}
	return model.FreelancerProfile{
		Details:   details,
		Skills:    skills,
		Resume:    model.FreelancerProfileResume{},
		Languages: languages,
	}, nil
}

var avatarRE = regexp.MustCompile(`url\("(.*)"\)`)

func GetFreelancerProfileDetails(ctx context.Context) (model.FreelancerProfileDetails, error) {
	fullNameSel := `//div[contains(@class, "FreelancerDetails_fullName")]`
	subtitleSel := `//div[contains(@class, "FreelancerDetails_subtitle")]`

	var nodes []*cdp.Node
	var fullName, subtitle, avatarURL string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Text(fullNameSel, &fullName, chromedp.NodeVisible),
		chromedp.Text(subtitleSel, &subtitle, chromedp.NodeVisible),
		chromedp.Nodes(`div.v-avatar>div.v-image>div.v-image__image`, &nodes),
	})
	if err != nil {
		return model.FreelancerProfileDetails{}, err
	}

	for _, node := range nodes {
		if res := avatarRE.FindStringSubmatch(node.AttributeValue("style")); res != nil {
			avatarURL = res[1]
		}
	}

	return model.FreelancerProfileDetails{
		FullName:  fullName,
		Subtitle:  subtitle,
		AvatarURL: avatarURL,
	}, nil
}

func GetFreelancerLanguages(ctx context.Context) ([]model.FreelancerProfileLanguage, error) {
	languagesSel := `//div[contains(@class, "FreelancerLanguages_freelancerLanguages")]/div[contains(@class, "v-card__text")]`

	var nodes []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Nodes(languagesSel, &nodes),
		chromedp.ActionFunc(func(c context.Context) error {
			for _, node := range nodes {
				if err := dom.RequestChildNodes(node.NodeID).WithDepth(2).Do(c); err != nil {
					return err
				}
			}
			return nil
		}),
		chromedp.Sleep(time.Millisecond * 200), // let's wait for the nodes to be loaded
	})
	if err != nil {
		return nil, err
	}

	res := make([]model.FreelancerProfileLanguage, 0)
	for _, node := range nodes {
		lang := model.FreelancerProfileLanguage{}
		for _, child := range node.Children {
			if child.NodeName == "SPAN" {
				spanVal := child.Children[0].NodeValue
				if strings.Contains(child.AttributeValue("class"), "font-weight-bold") {
					lang.Name = spanVal
				} else {
					lang.Level = spanVal
				}
			}
		}
		if lang.Name != "" && lang.Level != "" {
			res = append(res, lang)
		}
	}

	return res, nil
}

var skillPeriodRE = regexp.MustCompile(`· (.*) \+`)

func GetFreelancerSkills(ctx context.Context) ([]model.FreelancerProfileSkill, error) {
	skillSel := `//div[contains(@class, "FreelancerProfileSkills_freelancerProfileSkills")]//span[contains(@class, "FreelancerSkillsField_chip")]`
	var nodes []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Nodes(skillSel, &nodes),
	})
	if err != nil {
		return nil, err
	}

	res := make([]model.FreelancerProfileSkill, 0)
	for _, node := range nodes {
		var name, period string
		err = chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Text(`span.font-weight-semi-bold`, &name, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(`span.ml-1`, &period, chromedp.ByQuery, chromedp.FromNode(node)),
		})
		if err != nil {
			return nil, err
		}
		if name != "" && period != "" {
			if trimmedPeriod := skillPeriodRE.FindStringSubmatch(period); trimmedPeriod != nil {
				res = append(res, model.FreelancerProfileSkill{
					Name:   name,
					Period: trimmedPeriod[1],
				})
			}
		}
	}

	return res, nil
}
