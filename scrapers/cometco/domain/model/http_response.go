package model

import "time"

//	{
//	  "data": {
//	    "experience": {
//	      "id": 377160,
//	      "isCometMission": false,
//	      "startDate": "2019-01-01T00:00:00.000Z",
//	      "endDate": "2019-12-31T00:00:00.000Z",
//	      "companyName": "Freelance",
//	      "description": "Software Engineer\n",
//	      "location": null,
//	      "type": "freelancing",
//	      "skills": [
//	        {
//	          "id": 9,
//	          "name": "Django",
//	          "primary": false,
//	          "freelanceExperienceSkillId": 3670455,
//	          "__typename": "FreelanceExperienceSkill"
//	        }
//	      ],
//	      "__typename": "Experience"
//	    }
//	  }
//	}
type Experience struct {
	Data struct {
		Experience struct {
			Id             int         `json:"id"`
			IsCometMission bool        `json:"isCometMission"`
			StartDate      time.Time   `json:"startDate"`
			EndDate        time.Time   `json:"endDate"`
			CompanyName    string      `json:"companyName"`
			Description    string      `json:"description"`
			Location       interface{} `json:"location"`
			Type           string      `json:"type"`
			Skills         []struct {
				Id                         int    `json:"id"`
				Name                       string `json:"name"`
				Primary                    bool   `json:"primary"`
				FreelanceExperienceSkillId int    `json:"freelanceExperienceSkillId"`
				Typename                   string `json:"__typename"`
			} `json:"skills"`
			Typename string `json:"__typename"`
		} `json:"experience"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "recommendedSkills": [
//	     {
//	       "id": 14,
//	       "name": "Flask",
//	       "aliases": [],
//	       "popularity": 1094,
//	       "__typename": "Skill"
//	     },
//	     {
//	       "id": 31,
//	       "name": "Oracle",
//	       "aliases": [],
//	       "popularity": 6690,
//	       "__typename": "Skill"
//	     },
//	     {
//	       "id": 36,
//	       "name": "Python",
//	       "aliases": [],
//	       "popularity": 10698,
//	       "__typename": "Skill"
//	     }
//	   ]
//	 }
//	}
type T4 struct {
	Data struct {
		RecommendedSkills []struct {
			Id         int           `json:"id"`
			Name       string        `json:"name"`
			Aliases    []interface{} `json:"aliases"`
			Popularity int           `json:"popularity"`
			Typename   string        `json:"__typename"`
		} `json:"recommendedSkills"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "experiences": [
//	       {
//	         "id": 377175,
//	         "companyName": "Freelance",
//	         "__typename": "Experience"
//	       },
//	       {
//	         "id": 377160,
//	         "companyName": "Freelance",
//	         "__typename": "Experience"
//	       }
//	     ],
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type T5 struct {
	Data struct {
		Freelance struct {
			Id          int `json:"id"`
			Experiences []struct {
				Id          int    `json:"id"`
				CompanyName string `json:"companyName"`
				Typename    string `json:"__typename"`
			} `json:"experiences"`
			Typename string `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "linkedInUrl": "https://www.linkedin.com/in/test-rollee-908472243",
//	     "websiteUrl": null,
//	     "twitterUrl": null,
//	     "gitHubUrl": null,
//	     "kaggleUrl": null,
//	     "gitlabUrl": null,
//	     "stackExchangeUrl": null,
//	     "bitbucketUrl": null,
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type T6 struct {
	Data struct {
		Freelance struct {
			Id               int         `json:"id"`
			LinkedInUrl      string      `json:"linkedInUrl"`
			WebsiteUrl       interface{} `json:"websiteUrl"`
			TwitterUrl       interface{} `json:"twitterUrl"`
			GitHubUrl        interface{} `json:"gitHubUrl"`
			KaggleUrl        interface{} `json:"kaggleUrl"`
			GitlabUrl        interface{} `json:"gitlabUrl"`
			StackExchangeUrl interface{} `json:"stackExchangeUrl"`
			BitbucketUrl     interface{} `json:"bitbucketUrl"`
			Typename         string      `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "prefTime": "none",
//	     "prefEnvironment": "none",
//	     "prefMobility": "country",
//	     "prefContract": "fullTime",
//	     "prefWorkplace": "none",
//	     "retribution": 1000,
//	     "skills": [
//	       {
//	         "id": 9,
//	         "freelanceSkillId": 5582909,
//	         "name": "Django",
//	         "duration": 1,
//	         "wishedInMissions": false,
//	         "__typename": "FreelanceSkill"
//	       },
//	       {
//	         "id": 45,
//	         "freelanceSkillId": 5583241,
//	         "name": "SQL",
//	         "duration": 1,
//	         "wishedInMissions": false,
//	         "__typename": "FreelanceSkill"
//	       }
//	     ],
//	     "user": {
//	       "id": 77301,
//	       "address": {
//	         "id": 39001,
//	         "city": "Paris",
//	         "__typename": "Address"
//	       },
//	       "__typename": "User"
//	     },
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type SkillsResponse struct {
	Data struct {
		Freelance struct {
			Id              int    `json:"id"`
			PrefTime        string `json:"prefTime"`
			PrefEnvironment string `json:"prefEnvironment"`
			PrefMobility    string `json:"prefMobility"`
			PrefContract    string `json:"prefContract"`
			PrefWorkplace   string `json:"prefWorkplace"`
			Retribution     int    `json:"retribution"`
			Skills          []struct {
				Id               int    `json:"id"`
				FreelanceSkillId int    `json:"freelanceSkillId"`
				Name             string `json:"name"`
				Duration         int    `json:"duration"`
				WishedInMissions bool   `json:"wishedInMissions"`
				Typename         string `json:"__typename"`
			} `json:"skills"`
			User struct {
				Id      int `json:"id"`
				Address struct {
					Id       int    `json:"id"`
					City     string `json:"city"`
					Typename string `json:"__typename"`
				} `json:"address"`
				Typename string `json:"__typename"`
			} `json:"user"`
			Typename string `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "profileScore": 45,
//	     "profileStrength": 0,
//	     "profileCompletionTasks": [
//	       {
//	         "id": "notifications:onboarding:personal-profile-picture",
//	         "title": "Uploade ta plus belle photo de profil",
//	         "done": true,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:personal-profile-biography",
//	         "title": "Ajoute une biographie",
//	         "done": true,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:experiences",
//	         "title": "Ajoute ta première expérience",
//	         "done": true,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:three-experiences",
//	         "title": "Ajoute au moins 3 expériences",
//	         "done": true,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:three-years-experiences",
//	         "title": "Ajoute au moins 3 ans de compétences",
//	         "done": true,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:update-availability-last-three-months",
//	         "title": "Mets à jour ta disponibilité",
//	         "done": false,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:update-preference-last-three-months",
//	         "title": "Mets à jour ta recherche de job",
//	         "done": true,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:three-skill-on-one-experiences",
//	         "title": "Ajoute 3 compétences sur une de tes expériences",
//	         "done": false,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:three-skill-on-all-experiences",
//	         "title": "Ajoute 3 compétences sur chaque expériences",
//	         "done": false,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:long-description-on-every-experiences",
//	         "title": "Complète tes expériences avec plus de 15 mots",
//	         "done": false,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:has-social-link",
//	         "title": "Ajoute des liens vers tes autres profils",
//	         "done": false,
//	         "__typename": "ProfileTaskType"
//	       },
//	       {
//	         "id": "notifications:onboarding:has-education",
//	         "title": "Ajoute une formation",
//	         "done": false,
//	         "__typename": "ProfileTaskType"
//	       }
//	     ],
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type T8 struct {
	Data struct {
		Freelance struct {
			Id                     int `json:"id"`
			ProfileScore           int `json:"profileScore"`
			ProfileStrength        int `json:"profileStrength"`
			ProfileCompletionTasks []struct {
				Id       string `json:"id"`
				Title    string `json:"title"`
				Done     bool   `json:"done"`
				Typename string `json:"__typename"`
			} `json:"profileCompletionTasks"`
			Typename string `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "autoPilotEnabled": false,
//	     "languages": [
//	       {
//	         "id": 1,
//	         "name": "Français",
//	         "level": "native",
//	         "__typename": "FreelanceLanguage"
//	       }
//	     ],
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type Language struct {
	Data struct {
		Freelance struct {
			Id               int  `json:"id"`
			AutoPilotEnabled bool `json:"autoPilotEnabled"`
			Languages        []struct {
				Id       int    `json:"id"`
				Name     string `json:"name"`
				Level    string `json:"level"`
				Typename string `json:"__typename"`
			} `json:"languages"`
			Typename string `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "autoPilotEnabled": false,
//	     "referents": [],
//	     "experiences": [
//	       {
//	         "id": 377175,
//	         "__typename": "Experience"
//	       },
//	       {
//	         "id": 377160,
//	         "__typename": "Experience"
//	       }
//	     ],
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type T10 struct {
	Data struct {
		Freelance struct {
			Id               int           `json:"id"`
			AutoPilotEnabled bool          `json:"autoPilotEnabled"`
			Referents        []interface{} `json:"referents"`
			Experiences      []struct {
				Id       int    `json:"id"`
				Typename string `json:"__typename"`
			} `json:"experiences"`
			Typename string `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "experienceInYears": 2,
//	     "autoPilotEnabled": false,
//	     "experiences": [
//	       {
//	         "id": 377175,
//	         "isCometMission": false,
//	         "startDate": "2020-01-01T00:00:00.000Z",
//	         "endDate": "2020-12-31T00:00:00.000Z",
//	         "companyName": "Freelance",
//	         "description": "Worked on Sql",
//	         "location": null,
//	         "type": "freelancing",
//	         "skills": [
//	           {
//	             "id": 45,
//	             "name": "SQL",
//	             "primary": false,
//	             "freelanceExperienceSkillId": 3670458,
//	             "__typename": "FreelanceExperienceSkill"
//	           }
//	         ],
//	         "__typename": "Experience"
//	       },
//	       {
//	         "id": 377160,
//	         "isCometMission": false,
//	         "startDate": "2019-01-01T00:00:00.000Z",
//	         "endDate": "2019-12-31T00:00:00.000Z",
//	         "companyName": "Freelance",
//	         "description": "Software Engineer\n",
//	         "location": null,
//	         "type": "freelancing",
//	         "skills": [
//	           {
//	             "id": 9,
//	             "name": "Django",
//	             "primary": false,
//	             "freelanceExperienceSkillId": 3670455,
//	             "__typename": "FreelanceExperienceSkill"
//	           }
//	         ],
//	         "__typename": "Experience"
//	       }
//	     ],
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type T11 struct {
	Data struct {
		Freelance struct {
			Id                int  `json:"id"`
			ExperienceInYears int  `json:"experienceInYears"`
			AutoPilotEnabled  bool `json:"autoPilotEnabled"`
			Experiences       []struct {
				Id             int         `json:"id"`
				IsCometMission bool        `json:"isCometMission"`
				StartDate      time.Time   `json:"startDate"`
				EndDate        time.Time   `json:"endDate"`
				CompanyName    string      `json:"companyName"`
				Description    string      `json:"description"`
				Location       interface{} `json:"location"`
				Type           string      `json:"type"`
				Skills         []struct {
					Id                         int    `json:"id"`
					Name                       string `json:"name"`
					Primary                    bool   `json:"primary"`
					FreelanceExperienceSkillId int    `json:"freelanceExperienceSkillId"`
					Typename                   string `json:"__typename"`
				} `json:"skills"`
				Typename string `json:"__typename"`
			} `json:"experiences"`
			Typename string `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "biography": "Web Developer for 10 years, I'm currently working on creating the ultimate website.",
//	     "autoPilotEnabled": false,
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type T12 struct {
	Data struct {
		Freelance struct {
			Id               int    `json:"id"`
			Biography        string `json:"biography"`
			AutoPilotEnabled bool   `json:"autoPilotEnabled"`
			Typename         string `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "skills": [
//	       {
//	         "id": 9,
//	         "freelanceSkillId": 5582909,
//	         "name": "Django",
//	         "duration": 1,
//	         "wishedInMissions": false,
//	         "__typename": "FreelanceSkill"
//	       },
//	       {
//	         "id": 45,
//	         "freelanceSkillId": 5583241,
//	         "name": "SQL",
//	         "duration": 1,
//	         "wishedInMissions": false,
//	         "__typename": "FreelanceSkill"
//	       }
//	     ],
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type SkillsResponse1 struct {
	Data struct {
		Freelance struct {
			Id     int `json:"id"`
			Skills []struct {
				Id               int    `json:"id"`
				FreelanceSkillId int    `json:"freelanceSkillId"`
				Name             string `json:"name"`
				Duration         int    `json:"duration"`
				WishedInMissions bool   `json:"wishedInMissions"`
				Typename         string `json:"__typename"`
			} `json:"skills"`
			Typename string `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

// ProfileResponse
//
//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "prefMobility": "country",
//	     "retribution": 1000,
//	     "experienceInYears": 2,
//	     "averageRating": null,
//	     "linkedInUrl": "https://www.linkedin.com/in/test-rollee-908472243",
//	     "user": {
//	       "id": 77301,
//	       "jobTitle": "Software Engineer at Freelance (Self employed)",
//	       "fullName": "Rollee Test",
//	       "lastConnectionDate": "2022-09-05T16:00:12.489Z",
//	       "createdAt": "2022-06-27T12:58:45.570Z",
//	       "__typename": "User",
//	       "profilePictureUrl": "https://profile-pic-comet.s3.eu-central-1.amazonaws.com/FREELANCE/45891/243hv4f1l4wrogsq.jpeg"
//	     },
//	     "availabilityDate": null,
//	     "isAvailable": true,
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type ProfileResponse struct {
	Data struct {
		Freelance struct {
			Id                int         `json:"id"`
			PrefMobility      string      `json:"prefMobility"`
			Retribution       int         `json:"retribution"`
			ExperienceInYears int         `json:"experienceInYears"`
			AverageRating     interface{} `json:"averageRating"`
			LinkedInUrl       string      `json:"linkedInUrl"`
			User              struct {
				Id                 int       `json:"id"`
				JobTitle           string    `json:"jobTitle"`
				FullName           string    `json:"fullName"`
				LastConnectionDate time.Time `json:"lastConnectionDate"`
				CreatedAt          time.Time `json:"createdAt"`
				Typename           string    `json:"__typename"`
				ProfilePictureUrl  string    `json:"profilePictureUrl"`
			} `json:"user"`
			AvailabilityDate interface{} `json:"availabilityDate"`
			IsAvailable      bool        `json:"isAvailable"`
			Typename         string      `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "availabilityDate": null,
//	     "isAvailable": true,
//	     "daysBeforeSendingPropositions": 0,
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type T15 struct {
	Data struct {
		Freelance struct {
			Id                            int         `json:"id"`
			AvailabilityDate              interface{} `json:"availabilityDate"`
			IsAvailable                   bool        `json:"isAvailable"`
			DaysBeforeSendingPropositions int         `json:"daysBeforeSendingPropositions"`
			Typename                      string      `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "freelance": {
//	     "id": 45891,
//	     "retribution": 1000,
//	     "prefTime": "none",
//	     "prefWorkplace": "none",
//	     "prefEnvironment": "none",
//	     "prefContract": "fullTime",
//	     "prefMobility": "country",
//	     "skills": [
//	       {
//	         "id": 9,
//	         "freelanceSkillId": 5582909,
//	         "name": "Django",
//	         "duration": 1,
//	         "wishedInMissions": false,
//	         "__typename": "FreelanceSkill"
//	       },
//	       {
//	         "id": 45,
//	         "freelanceSkillId": 5583241,
//	         "name": "SQL",
//	         "duration": 1,
//	         "wishedInMissions": false,
//	         "__typename": "FreelanceSkill"
//	       }
//	     ],
//	     "user": {
//	       "id": 77301,
//	       "address": {
//	         "id": 39001,
//	         "street": null,
//	         "zipCode": null,
//	         "city": "Paris",
//	         "state": "IDF",
//	         "country": {
//	           "id": "FRA",
//	           "iso2": "FR",
//	           "iso3": "FRA",
//	           "name": "France",
//	           "__typename": "Country"
//	         },
//	         "point": {
//	           "latitude": 48.856614,
//	           "longitude": 2.352222,
//	           "__typename": "GraphQLCoordinates"
//	         },
//	         "__typename": "Address"
//	       },
//	       "__typename": "User"
//	     },
//	     "__typename": "Freelance"
//	   }
//	 }
//	}
type T16 struct {
	Data struct {
		Freelance struct {
			Id              int    `json:"id"`
			Retribution     int    `json:"retribution"`
			PrefTime        string `json:"prefTime"`
			PrefWorkplace   string `json:"prefWorkplace"`
			PrefEnvironment string `json:"prefEnvironment"`
			PrefContract    string `json:"prefContract"`
			PrefMobility    string `json:"prefMobility"`
			Skills          []struct {
				Id               int    `json:"id"`
				FreelanceSkillId int    `json:"freelanceSkillId"`
				Name             string `json:"name"`
				Duration         int    `json:"duration"`
				WishedInMissions bool   `json:"wishedInMissions"`
				Typename         string `json:"__typename"`
			} `json:"skills"`
			User struct {
				Id      int `json:"id"`
				Address struct {
					Id      int         `json:"id"`
					Street  interface{} `json:"street"`
					ZipCode interface{} `json:"zipCode"`
					City    string      `json:"city"`
					State   string      `json:"state"`
					Country struct {
						Id       string `json:"id"`
						Iso2     string `json:"iso2"`
						Iso3     string `json:"iso3"`
						Name     string `json:"name"`
						Typename string `json:"__typename"`
					} `json:"country"`
					Point struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
						Typename  string  `json:"__typename"`
					} `json:"point"`
					Typename string `json:"__typename"`
				} `json:"address"`
				Typename string `json:"__typename"`
			} `json:"user"`
			Typename string `json:"__typename"`
		} `json:"freelance"`
	} `json:"data"`
}

//	{
//	 "data": {
//	   "me": {
//	     "id": 77301,
//	     "freelance": {
//	       "id": 45891,
//	       "isInstructor": false,
//	       "status": "validated",
//	       "publicId": "pnelgwjdKB",
//	       "instructorSkills": [],
//	       "instructorSlots": [],
//	       "__typename": "Freelance",
//	       "flaggedForQualif": null,
//	       "candidateInterviews": [],
//	       "fetchingLinkedIn": false,
//	       "lastLinkedInImport": {
//	         "id": "3833300",
//	         "status": "done",
//	         "lastError": null,
//	         "importedAt": "2022-06-27T12:59:33.299Z",
//	         "__typename": "LinkedInImportType"
//	       },
//	       "biography": "Web Developer for 10 years, I'm currently working on creating the ultimate website.",
//	       "user": {
//	         "id": 77301,
//	         "profilePictureUrl": "https://profile-pic-comet.s3.eu-central-1.amazonaws.com/FREELANCE/45891/243hv4f1l4wrogsq.jpeg",
//	         "jobTitle": "Software Engineer at Freelance (Self employed)",
//	         "__typename": "User"
//	       }
//	     },
//	     "__typename": "User"
//	   }
//	 }
//	}
type T17 struct {
	Data struct {
		Me struct {
			Id        int `json:"id"`
			Freelance struct {
				Id                  int           `json:"id"`
				IsInstructor        bool          `json:"isInstructor"`
				Status              string        `json:"status"`
				PublicId            string        `json:"publicId"`
				InstructorSkills    []interface{} `json:"instructorSkills"`
				InstructorSlots     []interface{} `json:"instructorSlots"`
				Typename            string        `json:"__typename"`
				FlaggedForQualif    interface{}   `json:"flaggedForQualif"`
				CandidateInterviews []interface{} `json:"candidateInterviews"`
				FetchingLinkedIn    bool          `json:"fetchingLinkedIn"`
				LastLinkedInImport  struct {
					Id         string      `json:"id"`
					Status     string      `json:"status"`
					LastError  interface{} `json:"lastError"`
					ImportedAt time.Time   `json:"importedAt"`
					Typename   string      `json:"__typename"`
				} `json:"lastLinkedInImport"`
				Biography string `json:"biography"`
				User      struct {
					Id                int    `json:"id"`
					ProfilePictureUrl string `json:"profilePictureUrl"`
					JobTitle          string `json:"jobTitle"`
					Typename          string `json:"__typename"`
				} `json:"user"`
			} `json:"freelance"`
			Typename string `json:"__typename"`
		} `json:"me"`
	} `json:"data"`
}
