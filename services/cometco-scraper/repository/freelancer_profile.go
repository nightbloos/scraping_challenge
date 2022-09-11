package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"scraping_challenge/common/service/domain"
	"scraping_challenge/services/cometco-scraper/domain/model"
)

type FreelancerProfileRepository struct {
	db                    *mongo.Database
	freelancerProfileColl *mgm.Collection
}

func NewFreelancerProfileRepository(db *mongo.Database) *FreelancerProfileRepository {
	p := model.FreelancerProfile{}
	return &FreelancerProfileRepository{
		db:                    db,
		freelancerProfileColl: mgm.NewCollection(db, p.CollectionName()),
	}
}

func (r *FreelancerProfileRepository) FindByID(ctx context.Context, id string) (model.FreelancerProfile, error) {
	res := r.freelancerProfileColl.FindOne(ctx, bson.M{"uuid": id})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return model.FreelancerProfile{}, domain.WrapWithNotFoundError(res.Err(), "profile not found")
		}
		return model.FreelancerProfile{}, domain.WrapWithInternalError(res.Err(), "failed to find profile")
	}

	t := model.FreelancerProfile{}
	if err := res.Decode(&t); err != nil {
		return model.FreelancerProfile{}, domain.WrapWithInternalError(err, "failed to decode profile")
	}

	return t, nil
}

func (r *FreelancerProfileRepository) Create(_ context.Context, t model.FreelancerProfile) (model.FreelancerProfile, error) {
	err := t.Creating()
	if err != nil {
		return model.FreelancerProfile{}, domain.WrapWithInternalError(err, "failed to trigger create hooks for profile")
	}

	t.UUID = uuid.New().String()
	err = r.freelancerProfileColl.Create(&t)
	if err != nil {
		return model.FreelancerProfile{}, domain.WrapWithInternalError(err, "failed to create new profile")
	}

	return t, nil
}
