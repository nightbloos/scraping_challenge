package converter

import (
	"regexp"
	"strconv"
	"time"

	"github.com/goodsign/monday"
	"github.com/pkg/errors"

	"scraping_challenge/services/cometco-scraper/domain/model"
	cometModel "scraping_challenge/services/cometco-scraper/scraper/model"
)

func FromCometFreelancerProfile(profile cometModel.FreelancerProfile) (model.FreelancerProfile, error) {
	skills, err := FromCometFreelancerProfileSkills(profile.Skills)
	if err != nil {
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to convert skills")
	}

	resume, err := FromCometFreelancerProfileResume(profile.Resume)
	if err != nil {
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to convert resume")
	}

	return model.FreelancerProfile{
		Details:   FromCometFreelancerProfileDetails(profile.Details),
		Skills:    skills,
		Resume:    resume,
		Languages: FromCometFreelancerProfileLanguages(profile.Languages),
	}, nil
}

func FromCometFreelancerProfileDetails(details cometModel.FreelancerProfileDetails) model.FreelancerProfileDetails {
	return model.FreelancerProfileDetails{
		FullName:  details.FullName,
		Subtitle:  details.Subtitle,
		AvatarURL: details.AvatarURL,
	}
}

func FromCometFreelancerProfileSkills(skills []cometModel.FreelancerProfileSkill) ([]model.FreelancerProfileSkill, error) {
	res := make([]model.FreelancerProfileSkill, len(skills))
	for i, s := range skills {
		skill, err := FromCometFreelancerProfileSkill(s)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert skill")
		}
		res[i] = skill
	}

	return res, nil
}

var skillsYearsRE = regexp.MustCompile(`^(<?)\s*(\d*)\s*ans?\s*(\+?)$`)

func FromCometFreelancerProfileSkill(skill cometModel.FreelancerProfileSkill) (model.FreelancerProfileSkill, error) {
	years := int64(1)
	sign := model.LessSkillSign
	yearsMatch := skillsYearsRE.FindStringSubmatch(skill.Period)
	if len(yearsMatch) > 0 {
		if yearsMatch[2] != "" {
			var err error
			years, err = strconv.ParseInt(yearsMatch[2], 10, 64)
			if err != nil {
				return model.FreelancerProfileSkill{}, errors.Wrap(err, "failed to parse years")
			}
		}
		if yearsMatch[1] == "<" {
			sign = model.LessSkillSign
		} else if yearsMatch[3] == "+" {
			sign = model.MoreSkillSign
		}
	}

	return model.FreelancerProfileSkill{
		Name:  skill.Name,
		Years: years,
		Sign:  sign,
	}, nil
}

func FromCometFreelancerProfileResume(resume cometModel.FreelancerProfileResume) (model.FreelancerProfileResume, error) {
	ex, err := FromCometFreelancerProfileExperiences(resume.Experiences)
	if err != nil {
		return model.FreelancerProfileResume{}, errors.Wrap(err, "failed to convert experiences")
	}
	return model.FreelancerProfileResume{
		Biography:   resume.Biography,
		Experiences: ex,
	}, nil
}

func FromCometFreelancerProfileExperiences(experience []cometModel.FreelancerProfileExperience) ([]model.FreelancerProfileExperience, error) {
	res := make([]model.FreelancerProfileExperience, len(experience))
	for i, e := range experience {
		ex, err := FromCometFreelancerProfileExperience(e)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert experience")

		}
		res[i] = ex
	}

	return res, nil
}

func FromCometFreelancerProfileExperience(experience cometModel.FreelancerProfileExperience) (model.FreelancerProfileExperience, error) {
	startAt, err := parseDate(experience.StartAt)
	if err != nil {
		return model.FreelancerProfileExperience{}, errors.Wrap(err, "failed to parse start date")
	}
	endAt := time.Time{}
	if experience.EndAt != "En cours" {
		if parsedTime, err := parseDate(experience.EndAt); err == nil {
			endAt = parsedTime
		}
	}
	return model.FreelancerProfileExperience{
		Location:    experience.Location,
		Company:     experience.Company,
		Type:        experience.Type,
		Description: experience.Description,
		StartAt:     startAt,
		EndAt:       endAt,
		Skills:      experience.Skills,
	}, nil
}

func FromCometFreelancerProfileLanguages(languages []cometModel.FreelancerProfileLanguage) []model.FreelancerProfileLanguage {
	res := make([]model.FreelancerProfileLanguage, len(languages))
	for i, l := range languages {
		res[i] = FromCometFreelancerProfileLanguage(l)
	}

	return res
}

func FromCometFreelancerProfileLanguage(language cometModel.FreelancerProfileLanguage) model.FreelancerProfileLanguage {
	return model.FreelancerProfileLanguage{
		Name:  language.Name,
		Level: language.Level,
	}
}

func parseDate(date string) (time.Time, error) {
	parsedTime, err := monday.Parse("Jan 2006", date, monday.LocaleFrFR)
	if err == nil {
		return parsedTime, nil
	}
	parsedTime, err = monday.Parse("Jan. 2006", date, monday.LocaleFrFR)
	if err == nil {
		return parsedTime, nil
	}
	return time.Time{}, errors.Wrap(err, "failed to parse date")
}
