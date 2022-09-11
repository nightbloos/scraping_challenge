package scraper

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"scraping_challenge/services/cometco-scraper/config"
	"scraping_challenge/services/cometco-scraper/domain/model"
	"scraping_challenge/services/cometco-scraper/scraper/converter"
	"scraping_challenge/services/cometco-scraper/scraper/converter/parser"
)

type Scraper struct {
	taskRepo      TaskRepo
	flProfileRepo FreelancerProfileRepo
	allocOpts     []chromedp.ExecAllocatorOption
	ctxOpts       []chromedp.ContextOption
	logger        *zap.Logger
}

func NewScraper(taskRepo TaskRepo, flProfileRepo FreelancerProfileRepo, cfg config.ChromeDPConfig, logger *zap.Logger) *Scraper {
	ctxOpts := make([]chromedp.ContextOption, 0)
	if cfg.Debug {
		ctxOpts = append(ctxOpts, chromedp.WithDebugf(logger.Sugar().Infof))
	}

	allocOpts := chromedp.DefaultExecAllocatorOptions[:]
	if !cfg.Headless {
		allocOpts = append(allocOpts, chromedp.Flag("headless", false))
	}

	return &Scraper{
		taskRepo:      taskRepo,
		flProfileRepo: flProfileRepo,
		allocOpts:     allocOpts,
		ctxOpts:       ctxOpts,
		logger:        logger,
	}
}

func (s *Scraper) Run(ctx context.Context, login, pass string) error {
	task, err := s.taskRepo.Create(ctx, model.Task{Status: model.TaskStatusInProgress})
	if err != nil {
		return errors.Wrap(err, "failed to create task")
	}

	logger := s.logger.With(zap.String("task_id", task.UUID))

	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, s.allocOpts...)
	defer allocCancel()

	browserCtx, contextCancel := chromedp.NewContext(allocCtx, s.ctxOpts...)
	defer contextCancel()

	err = s.login(browserCtx, login, pass)
	if err != nil {
		s.FinishTask(ctx, task, errors.Wrap(err, "failed to login"))
		logger.Error("failed to login", zap.Error(err))

		return errors.Wrap(err, "failed to login")
	}

	err = s.navigateToProfile(browserCtx)
	if err != nil {
		s.FinishTask(ctx, task, errors.Wrap(err, "failed to navigate to profile"))
		logger.Error("failed to navigate to profile", zap.Error(err))

		return errors.Wrap(err, "failed to navigate to profile")
	}

	cometFreelancerProfile, err := parser.ParseFreelancerProfile(browserCtx)
	if err != nil {
		s.FinishTask(ctx, task, errors.Wrap(err, "failed to parse freelancer profile"))
		logger.Error("failed to parse freelancer profile", zap.Error(err))

		return errors.Wrap(err, "failed to parse freelancer profile")
	}

	freelancerProfile, err := converter.FromCometFreelancerProfile(cometFreelancerProfile)
	if err != nil {
		s.FinishTask(ctx, task, errors.Wrap(err, "failed to convert comet freelancer profile"))
		logger.Error("failed to convert comet freelancer profile", zap.Error(err))

		return errors.Wrap(err, "failed to convert comet freelancer profile")
	}

	freelancerProfile, err = s.flProfileRepo.Create(ctx, freelancerProfile)
	if err != nil {
		s.FinishTask(ctx, task, errors.Wrap(err, "failed to create freelancer profile"))
		logger.Error("failed to create freelancer profile", zap.Error(err))

		return errors.Wrap(err, "failed to create freelancer profile")
	}

	task.ResultID = freelancerProfile.UUID
	s.FinishTask(ctx, task, nil)

	return nil
}

func (s *Scraper) login(ctx context.Context, login, pass string) error {
	err := chromedp.Run(ctx, chromedp.Navigate("https://app.comet.co/freelancer"))
	if err != nil {
		return errors.Wrap(err, "failed to navigate to comet.co")
	}

	emailInputSel := `//input[@name="email"]`
	passInputSel := `//input[@name="password"]`
	submitButtonSel := `//button[@type="submit"]`
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.WaitVisible(emailInputSel),
		chromedp.WaitVisible(passInputSel),
		chromedp.SendKeys(emailInputSel, login),
		chromedp.SendKeys(passInputSel, pass),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(submitButtonSel),
	})
	if err != nil {
		return errors.Wrap(err, "failed to fill login form")
	}

	err = chromedp.Run(ctx, chromedp.Tasks{})
	if err != nil {
		return errors.Wrap(err, "failed to submit login form")
	}

	// TODO: handle error if is present `Mot de passe incorrect.` after this step
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.WaitNotPresent(emailInputSel),
		chromedp.WaitNotPresent(passInputSel),
	})
	if err != nil {
		return errors.Wrap(err, "failed to wait for login form to disappear")
	}

	return nil
}

func (s *Scraper) navigateToProfile(ctx context.Context) error {
	err := chromedp.Run(ctx, chromedp.Navigate("https://app.comet.co/freelancer/profile"))
	if err != nil {
		return errors.Wrap(err, "failed to navigate to comet.co")
	}

	return nil
}

func (s *Scraper) FinishTask(ctx context.Context, task model.Task, err error) {
	logger := s.logger.With(zap.String("task_id", task.UUID))
	if err != nil {
		task.Status = model.TaskStatusFailed
		task.FailReason = err.Error()
	} else {
		task.Status = model.TaskStatusCompleted
	}

	_, err = s.taskRepo.Update(ctx, task)
	if err != nil {
		logger.Error("failed to update task", zap.Error(err), zap.String("status", string(task.Status)), zap.String("fail_reason", task.FailReason))
	}
}
