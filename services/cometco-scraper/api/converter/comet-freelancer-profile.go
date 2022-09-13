package converter

import (
	httpModel "scraping_challenge/services/cometco-scraper/api/model"
	"scraping_challenge/services/cometco-scraper/domain/model"
)

func ToCometFreelancerProfile(profile model.FreelancerProfile) httpModel.CometFreelancerProfile {
	return httpModel.CometFreelancerProfile{
		ID:        profile.UUID,
		Details:   ToCometFreelancerProfileDetails(profile.Details),
		Skills:    ToCometFreelancerProfileSkills(profile.Skills),
		Resume:    ToCometFreelancerProfileResume(profile.Resume),
		Languages: ToCometFreelancerProfileLanguages(profile.Languages),
	}
}

func ToCometFreelancerProfileDetails(details model.FreelancerProfileDetails) httpModel.CometFreelancerProfileDetails {
	return httpModel.CometFreelancerProfileDetails{
		FullName:  details.FullName,
		Subtitle:  details.Subtitle,
		AvatarURL: details.AvatarURL,
	}
}

func ToCometFreelancerProfileSkills(skills []model.FreelancerProfileSkill) []httpModel.CometFreelancerProfileSkill {
	resSkills := make([]httpModel.CometFreelancerProfileSkill, len(skills))
	for i, skill := range skills {
		resSkills[i] = ToCometFreelancerProfileSkill(skill)
	}

	return resSkills
}

func ToCometFreelancerProfileSkill(skill model.FreelancerProfileSkill) httpModel.CometFreelancerProfileSkill {
	return httpModel.CometFreelancerProfileSkill{
		Name:  skill.Name,
		Years: skill.Years,
		Sign:  ToCometSkillSign(skill.Sign),
	}
}

func ToCometSkillSign(sign model.SkillSign) httpModel.CometSkillSign {
	switch sign {
	case model.LessSkillSign:
		return httpModel.LessSkillSign
	case model.MoreSkillSign:
		return httpModel.MoreSkillSign
	}

	return httpModel.LessSkillSign
}

func ToCometFreelancerProfileResume(resume model.FreelancerProfileResume) httpModel.CometFreelancerProfileResume {
	return httpModel.CometFreelancerProfileResume{
		Biography:   resume.Biography,
		Experiences: ToCometFreelancerProfileExperiences(resume.Experiences),
	}
}

func ToCometFreelancerProfileExperiences(experiences []model.FreelancerProfileExperience) []httpModel.CometFreelancerProfileExperience {
	resExperiences := make([]httpModel.CometFreelancerProfileExperience, len(experiences))
	for i, experience := range experiences {
		resExperiences[i] = ToCometFreelancerProfileExperience(experience)
	}

	return resExperiences
}

func ToCometFreelancerProfileExperience(experience model.FreelancerProfileExperience) httpModel.CometFreelancerProfileExperience {
	return httpModel.CometFreelancerProfileExperience{
		Location:    experience.Location,
		Company:     experience.Company,
		Type:        experience.Type,
		Description: experience.Description,
		StartAt:     experience.StartAt,
		EndAt:       experience.EndAt,
		Skills:      experience.Skills,
	}
}

func ToCometFreelancerProfileLanguages(languages []model.FreelancerProfileLanguage) []httpModel.CometFreelancerProfileLanguage {
	resLanguages := make([]httpModel.CometFreelancerProfileLanguage, len(languages))
	for i, language := range languages {
		resLanguages[i] = ToCometFreelancerProfileLanguage(language)
	}

	return resLanguages
}

func ToCometFreelancerProfileLanguage(language model.FreelancerProfileLanguage) httpModel.CometFreelancerProfileLanguage {
	return httpModel.CometFreelancerProfileLanguage{
		Name:  language.Name,
		Level: language.Level,
	}
}
