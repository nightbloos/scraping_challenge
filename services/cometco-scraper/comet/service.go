package comet

import (
	"context"

	"go.uber.org/zap"
	"scraping_challenge/services/cometco-scraper/config"
	"scraping_challenge/services/cometco-scraper/domain/model"
)

type Service struct {
	defaultCredentials    config.ProfileCredentialsConfig
	taskRepo              TaskRepo
	freelancerProfileRepo FreelancerProfileRepo
	taskCh                chan<- TaskWithCredentials
	logger                *zap.Logger
}

func NewService(
	defaultCredentials config.ProfileCredentialsConfig,
	taskRepo TaskRepo,
	freelancerProfileRepo FreelancerProfileRepo,
	taskCh chan<- TaskWithCredentials,
	logger *zap.Logger,
) *Service {
	return &Service{
		defaultCredentials:    defaultCredentials,
		taskRepo:              taskRepo,
		freelancerProfileRepo: freelancerProfileRepo,
		taskCh:                taskCh,
		logger:                logger.With(zap.String("internal-service", "comet-service")),
	}
}

func (s *Service) CreateTask(ctx context.Context, login string, pass string) (string, error) {
	task, err := s.taskRepo.Create(ctx, model.Task{Status: model.TaskStatusInProgress})
	if err != nil {
		s.logger.Error("failed to create task", zap.Error(err))
		return "", err
	}

	login, pass = s.getProfileCredentials(login, pass)

	s.taskCh <- TaskWithCredentials{
		Task:     task,
		Login:    login,
		Password: pass,
	}

	return task.UUID, nil
}

func (s *Service) GetTask(ctx context.Context, id string) (model.Task, error) {
	task, err := s.taskRepo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get task", zap.Error(err))
		return model.Task{}, err
	}

	return task, nil
}

func (s *Service) GetTasks(ctx context.Context, limit int64, offset int64) ([]model.Task, error) {
	tasks, err := s.taskRepo.Find(ctx, limit, offset)
	if err != nil {
		s.logger.Error("failed to get tasks", zap.Error(err))
		return nil, err
	}

	return tasks, nil
}

func (s *Service) GetUserProfile(ctx context.Context, id string) (model.FreelancerProfile, error) {
	profile, err := s.freelancerProfileRepo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get profile", zap.Error(err))
		return model.FreelancerProfile{}, err
	}

	return profile, nil
}

func (s *Service) getProfileCredentials(login string, pass string) (string, string) {
	if login == "" && pass == "" {
		return s.defaultCredentials.Email, s.defaultCredentials.Pass
	}

	return login, pass
}
