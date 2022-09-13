package scraper

import (
	"context"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"scraping_challenge/services/cometco-scraper/config"
	"scraping_challenge/services/cometco-scraper/domain/model"
	"scraping_challenge/services/cometco-scraper/scraper/converter"
	"scraping_challenge/services/cometco-scraper/scraper/parser"
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
	if !cfg.Headless {
		allocOpts = append(allocOpts, chromedp.Flag("headless", false))
	}

	return &Scraper{
		allocOpts: allocOpts,
		ctxOpts:   ctxOpts,
		logger:    logger,
	}
}

func (s *Scraper) Run(ctx context.Context, login, pass string) (model.FreelancerProfile, error) {
	logger := s.logger.With(zap.String("login", login))

	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, s.allocOpts...)
	defer allocCancel()

	browserCtx, contextCancel := chromedp.NewContext(allocCtx, s.ctxOpts...)
	defer contextCancel()

	err := s.login(browserCtx, login, pass)
	if err != nil {
		logger.Error("failed to login", zap.Error(err))
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to login")
	}

	err = s.navigateToProfile(browserCtx)
	if err != nil {
		logger.Error("failed to navigate to profile", zap.Error(err))
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to navigate to profile")
	}

	cometFreelancerProfile, err := parser.ParseFreelancerProfile(browserCtx)
	if err != nil {
		logger.Error("failed to parse freelancer profile", zap.Error(err))
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to parse freelancer profile")
	}

	freelancerProfile, err := converter.FromCometFreelancerProfile(cometFreelancerProfile)
	if err != nil {
		logger.Error("failed to convert comet freelancer profile", zap.Error(err))
		return model.FreelancerProfile{}, errors.Wrap(err, "failed to convert comet freelancer profile")
	}

	return freelancerProfile, nil
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

	successLogin, err := s.checkSuccessLogin(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to check success login")
	}

	if !successLogin {
		return errors.New("failed to login")
	}

	return nil
}

func (s *Scraper) checkSuccessLogin(ctx context.Context) (bool, error) {
	errSel := `//div[contains(@class, "Signin_signin")]//li[contains(text(), "Mot de passe incorrect")]`
	welcomeSel := `//div[contains(@class, "DashboardView_header")]//h4[contains(text(),'Bienvenue')]`

	timeoutTicket := time.NewTicker(time.Second * 15)
	checkTicker := time.NewTicker(time.Millisecond * 100)
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()

		case <-timeoutTicket.C:
			return false, errors.New("For too long there was no success login check")

		case <-checkTicker.C:
			var errNode, welcomeNode []*cdp.Node
			err := chromedp.Run(ctx, chromedp.Tasks{
				chromedp.Nodes(errSel, &errNode, chromedp.AtLeast(0)),
				chromedp.Nodes(welcomeSel, &welcomeNode, chromedp.AtLeast(0)),
			})
			if err != nil {
				return false, errors.Wrap(err, "failed to check success login")
			}
			if len(welcomeNode) > 0 {
				return true, nil
			} else if len(errNode) > 0 {
				return false, nil
			}
			// else ... let's wait for next check :)
		}
	}
}

func (s *Scraper) navigateToProfile(ctx context.Context) error {
	err := chromedp.Run(ctx, chromedp.Navigate("https://app.comet.co/freelancer/profile"))
	if err != nil {
		return errors.Wrap(err, "failed to navigate to comet.co")
	}

	return nil
}
