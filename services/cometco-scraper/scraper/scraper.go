package scraper

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"scraping_challenge/services/cometco-scraper/config"
	"scraping_challenge/services/cometco-scraper/parser"
)

type Scraper struct {
	allocOpts []chromedp.ExecAllocatorOption
	ctxOpts   []chromedp.ContextOption
	logger    *zap.Logger
}

func NewScraper(cfg config.ChromeDPConfig, logger *zap.Logger) *Scraper {
	ctxOpts := make([]chromedp.ContextOption, 0)
	if cfg.Debug {
		ctxOpts = append(ctxOpts, chromedp.WithDebugf(logger.Sugar().Infof))
	}

	allocOpts := chromedp.DefaultExecAllocatorOptions[:]
	if cfg.Headless {
		allocOpts = append(allocOpts, chromedp.Flag("headless", false))
	}

	return &Scraper{
		allocOpts: allocOpts,
		ctxOpts:   ctxOpts,
		logger:    logger,
	}
}

func (s *Scraper) Run(ctx context.Context, login, pass string) error {
	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, s.allocOpts...)
	defer allocCancel()

	browserCtx, contextCancel := chromedp.NewContext(allocCtx, s.ctxOpts...)
	defer contextCancel()

	err := s.login(browserCtx, login, pass)
	if err != nil {
		s.logger.Error("failed to login", zap.Error(err))
		return errors.Wrap(err, "failed to login")
	}

	err = s.navigateToProfile(browserCtx)
	if err != nil {
		s.logger.Error("failed to navigate to profile", zap.Error(err))
		return errors.Wrap(err, "failed to navigate to profile")
	}

	freelancerProfile, err := parser.ParseFreelancerProfile(browserCtx)
	if err != nil {
		s.logger.Error("failed to parse freelancer profile", zap.Error(err))
		return errors.Wrap(err, "failed to parse freelancer profile")
	}

	s.logger.Info("freelancer profile", zap.Any("profile", freelancerProfile))

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
