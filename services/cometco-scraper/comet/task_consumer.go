package comet

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"scraping_challenge/services/cometco-scraper/domain/model"
)

type Consumer struct {
	taskRepo              TaskRepo
	freelancerProfileRepo FreelancerProfileRepo
	scraper               Scraper
	logger                *zap.Logger
}

type TaskWithCredentials struct {
	Task     model.Task
	Login    string
	Password string
}

func NewConsumer(
	taskRepo TaskRepo,
	freelancerProfileRepo FreelancerProfileRepo,
	scraper Scraper,
	logger *zap.Logger,
) *Consumer {
	return &Consumer{
		taskRepo:              taskRepo,
		freelancerProfileRepo: freelancerProfileRepo,
		scraper:               scraper,
		logger:                logger.With(zap.String("internal-service", "comet-consumer")),
	}
}

func (c *Consumer) ConsumeTasks(ctx context.Context, twcCh <-chan TaskWithCredentials) error {
	for {
		select {
		case twc := <-twcCh:
			err := c.handleIncomingTask(context.Background(), twc)
			if err != nil {
				c.logger.Error("failed to handle incoming task", zap.Error(err), zap.String("taskID", twc.Task.UUID))
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}

}

func (c *Consumer) handleIncomingTask(ctx context.Context, twc TaskWithCredentials) error {
	task := twc.Task

	logger := c.logger.With(
		zap.String("taskID", task.UUID),
		zap.String("login", twc.Login),
	)

	profile, err := c.scraper.Run(ctx, twc.Login, twc.Password)
	if err != nil {
		c.FinishTask(ctx, task, err)
		logger.Error("failed to scrape profile", zap.Error(err))

		return err
	}

	profile, err = c.freelancerProfileRepo.Create(ctx, profile)
	if err != nil {
		logger.Error("failed to create freelancer profile", zap.Error(err))
		return errors.Wrap(err, "failed to create freelancer profile")
	}

	task.ResultID = profile.UUID
	c.FinishTask(ctx, task, nil)

	return nil
}

func (c *Consumer) FinishTask(ctx context.Context, task model.Task, err error) {
	logger := c.logger.With(zap.String("task_id", task.UUID))
	if err != nil {
		task.Status = model.TaskStatusFailed
		task.FailReason = err.Error()
	} else {
		task.Status = model.TaskStatusCompleted
	}

	_, err = c.taskRepo.Update(ctx, task)
	if err != nil {
		logger.Error("failed to update task", zap.Error(err), zap.String("status", string(task.Status)), zap.String("fail_reason", task.FailReason))
	}
}
