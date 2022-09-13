package api

import (
	"context"

	"scraping_challenge/services/cometco-scraper/domain/model"
)

type CometService interface {
	CreateTask(ctx context.Context, login string, pass string) (string, error)
	GetTask(ctx context.Context, id string) (model.Task, error)
	GetTasks(ctx context.Context, limit int64, offset int64) ([]model.Task, error)
	GetUserProfile(ctx context.Context, id string) (model.FreelancerProfile, error)
}
