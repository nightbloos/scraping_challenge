package app

import (
	"context"

	"golang.org/x/sync/errgroup"
	"scraping_challenge/common/http"
	"scraping_challenge/services/cometco-scraper/api"
	"scraping_challenge/services/cometco-scraper/comet"
	"scraping_challenge/services/cometco-scraper/repository"
	"scraping_challenge/services/cometco-scraper/scraper"
)

func (a *Application) initServices(
	grpCtx context.Context,
	errGrp *errgroup.Group,
) {
	articleCh := make(chan comet.TaskWithCredentials)

	taskRepo := repository.NewTaskRepository(a.db)
	freelancerProfileRepo := repository.NewFreelancerProfileRepository(a.db)
	cometService := comet.NewService(
		a.config.ProfileCredentials,
		taskRepo,
		freelancerProfileRepo,
		articleCh,
		a.logger,
	)

	ginRouter := a.initGinRouter()
	cometApi := api.NewCometServer(cometService, a.logger)
	cometApi.Register(ginRouter)

	scrp := scraper.NewScraper(a.config.ChromeDP, a.logger)
	cometConsumer := comet.NewConsumer(taskRepo, freelancerProfileRepo, scrp, a.logger)
	errGrp.Go(func() error {
		return cometConsumer.ConsumeTasks(grpCtx, articleCh)
	})

	errGrp.Go(func() error {
		return http.NewServer(a.config.HTTP.Port, a.logger).Run(grpCtx, ginRouter)
	})
}
