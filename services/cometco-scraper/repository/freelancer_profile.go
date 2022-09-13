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

	p := model.FreelancerProfile{}
	if err := res.Decode(&p); err != nil {
		return model.FreelancerProfile{}, domain.WrapWithInternalError(err, "failed to decode profile")
	}

	return p, nil
}

func (r *FreelancerProfileRepository) Create(_ context.Context, p model.FreelancerProfile) (model.FreelancerProfile, error) {
	err := p.Creating()
	if err != nil {
		return model.FreelancerProfile{}, domain.WrapWithInternalError(err, "failed to trigger create hooks for profile")
	}

	p.UUID = uuid.New().String()
	err = r.freelancerProfileColl.Create(&p)
	if err != nil {
		return model.FreelancerProfile{}, domain.WrapWithInternalError(err, "failed to create new profile")
	}

	return p, nil
}
