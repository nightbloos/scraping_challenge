package app

import (
	"context"

	"golang.org/x/sync/errgroup"
	"scraping_challenge/services/cometco-scraper/scraper"
)

func (a *Application) initServices(
	grpCtx context.Context,
	errGrp *errgroup.Group,
) {
	scrp := a.initScraper()
	errGrp.Go(func() error {
		return scrp.Run(grpCtx, a.config.ProfileCredentials.Email, a.config.ProfileCredentials.Pass)
	})

}

func (a *Application) initScraper() *scraper.Scraper {
	return scraper.NewScraper(a.config.ChromeDP, a.logger)
}
