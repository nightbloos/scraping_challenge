package app

import (
	"context"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"scraping_challenge/services/cometco-scraper/config"
)

type Application struct {
	config config.Config
	logger *zap.Logger
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Run(ctx context.Context) error {
	rand.Seed(time.Now().Unix())

	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	a.config = cfg

	if err := a.initLogger(); err != nil {
		return err
	}

	errGrp, ctx := errgroup.WithContext(ctx)

	a.initServices(ctx, errGrp)

	return errGrp.Wait()
}
