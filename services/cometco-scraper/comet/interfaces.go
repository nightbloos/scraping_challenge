package comet

import (
	"context"

	"scraping_challenge/services/cometco-scraper/domain/model"
)

type TaskRepo interface {
	Find(ctx context.Context, limit int64, offset int64) ([]model.Task, error)
	Create(ctx context.Context, t model.Task) (model.Task, error)
	FindByID(ctx context.Context, id string) (model.Task, error)
	Update(ctx context.Context, t model.Task) (model.Task, error)
}

type FreelancerProfileRepo interface {
	FindByID(ctx context.Context, id string) (model.FreelancerProfile, error)
	Create(ctx context.Context, p model.FreelancerProfile) (model.FreelancerProfile, error)
}

type Scraper interface {
	Run(ctx context.Context, login string, pass string) (model.FreelancerProfile, error)
}
