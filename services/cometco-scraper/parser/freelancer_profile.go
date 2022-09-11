package parser

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"

	"scraping_challenge/services/cometco-scraper/domain/model"
)

func ParseFreelancerProfile(ctx context.Context) (model.FreelancerProfile, error) {
	details, err := parseFreelancerProfileDetails(ctx)
	if err != nil {
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to create freelancer profile details")
	}
	languages, err := parseFreelancerLanguages(ctx)
	if err != nil {
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to create freelancer profile languages")
	}
	skills, err := parseFreelancerSkills(ctx)
	if err != nil {
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to create freelancer profile skills")
	}

	resume, err := parseFreelancerResume(ctx)
	if err != nil {
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to parse freelance profile resume")
	}

	return model.FreelancerProfile{
		Details:   details,
		Skills:    skills,
		Resume:    resume,
		Languages: languages,
	}, nil
}

var avatarRE = regexp.MustCompile(`url\("(.*)"\)`)

func parseFreelancerProfileDetails(ctx context.Context) (model.FreelancerProfileDetails, error) {
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

func parseFreelancerLanguages(ctx context.Context) ([]model.FreelancerProfileLanguage, error) {
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
		chromedp.Sleep(time.Millisecond * 100), // let's wait for the nodes to be loaded
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

var skillPeriodRE = regexp.MustCompile(`· (.*)`)

func parseFreelancerSkills(ctx context.Context) ([]model.FreelancerProfileSkill, error) {
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

var workPeriodRE = regexp.MustCompile(`(.*) - (.*) ·`)

func parseFreelancerResume(ctx context.Context) (model.FreelancerProfileResume, error) {
	resumeSel := `//div[contains(@class, "FreelancerFullResume_freelancerExperiences")]`
	bioSel := resumeSel + `//div[@data-cy-freelancer-biography]`
	experienceSel := `//div[@id="freelancerProfileExperiences"]/div`

	var nodes []*cdp.Node
	var bio string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Text(bioSel, &bio),
		chromedp.Nodes(experienceSel, &nodes),
	})
	if err != nil {
		return model.FreelancerProfileResume{}, err
	}

	res := model.FreelancerProfileResume{
		Biography: bio,
	}

	skillsSel := `//div[contains(@class, "FreelancerExperiences_skills")]//div[contains(@class, "Tag_tag")]//span[contains(@class, "Tag_name")]`
	descriptionSel := `//div[contains(@class, "FreelancerExperiences_description")]`
	freelanceFreelancerExperiencesHeaderSel := `//div[contains(@class, "FreelancerExperiences_headerMain")]`
	companyDetailsSel := freelanceFreelancerExperiencesHeaderSel + `/div[1]/div`
	periodAndLocationSel := freelanceFreelancerExperiencesHeaderSel + `//div//i`
	for _, node := range nodes {
		if !strings.Contains(node.AttributeValue("id"), "experience") {
			continue
		}
		var description string
		var companyNodes, skillNodes, periodAndLocationNodes []*cdp.Node

		nodeXPath := node.FullXPath()
		err = chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Text(nodeXPath+descriptionSel, &description),
			chromedp.Nodes(nodeXPath+skillsSel, &skillNodes),
			chromedp.Nodes(nodeXPath+companyDetailsSel, &companyNodes),
			chromedp.Nodes(nodeXPath+periodAndLocationSel, &periodAndLocationNodes),
			chromedp.ActionFunc(func(c context.Context) error {
				for _, n := range periodAndLocationNodes {
					if err = dom.RequestChildNodes(n.NodeID).WithDepth(1).Do(c); err != nil {
						return err
					}
				}
				for _, n := range companyNodes {
					if err = dom.RequestChildNodes(n.NodeID).WithDepth(1).Do(c); err != nil {
						return err
					}
				}
				return nil
			}),
			chromedp.Sleep(time.Millisecond * 100), // let's wait for the nodes to be loaded
		})
		if err != nil {
			return model.FreelancerProfileResume{}, err
		}

		skills := make([]string, 0)
		for _, sn := range skillNodes {
			skills = append(skills, sn.Children[0].NodeValue)
		}

		var companyName, employmentType string
		if len(companyNodes) == 2 {
			employmentType = companyNodes[0].Children[0].Children[0].NodeValue
			companyName = companyNodes[1].Children[0].Children[0].NodeValue
		}

		var workPeriod, location string
		for _, periodNode := range periodAndLocationNodes {
			switch {
			case strings.Contains(periodNode.AttributeValue("class"), "mdi-calendar-outline"):
				workPeriod = periodNode.Parent.Children[1].Children[0].NodeValue
			case strings.Contains(periodNode.AttributeValue("class"), "mdi-map-marker-outline"):
				location = periodNode.Parent.Children[1].Children[0].NodeValue
			}
		}

		var startAt, endAt string
		if reSubMatch := workPeriodRE.FindStringSubmatch(workPeriod); reSubMatch != nil {
			startAt = reSubMatch[1]
			endAt = reSubMatch[2]
		}

		experience := model.FreelancerProfileExperience{
			Location:    location,
			Company:     companyName,
			Type:        employmentType,
			Description: description,
			StartAt:     startAt,
			EndAt:       endAt,
			Skills:      skills,
		}

		res.Experiences = append(res.Experiences, experience)
	}

	return res, nil
}
