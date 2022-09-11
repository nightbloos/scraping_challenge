package app

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const dbName = "scraper"

func (a *Application) initDB(ctx context.Context) error {
	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(a.config.DB.MongodbURI))
	if err != nil {
		a.logger.Error("failed to create mongodb client", zap.Error(err))
		return err
	}

	_, client, database, _ := mgm.DefaultConfigs()

	closeFn := func() {
		if err := client.Disconnect(ctx); err != nil {
			a.logger.Error("DB close error", zap.Error(err))
		}
	}

	a.db = database
	a.dbClosFn = closeFn
	return nil
}
