package scraper

import (
	"context"

	"scraping_challenge/services/cometco-scraper/domain/model"
)

type TaskRepo interface {
	Create(ctx context.Context, task model.Task) (model.Task, error)
	Update(ctx context.Context, task model.Task) (model.Task, error)
}

type FreelancerProfileRepo interface {
	Create(ctx context.Context, profile model.FreelancerProfile) (model.FreelancerProfile, error)
}
