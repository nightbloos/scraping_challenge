package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"scraping_challenge/common/service/domain"
	"scraping_challenge/services/cometco-scraper/domain/model"
)

type TaskRepository struct {
	db       *mongo.Database
	taskColl *mgm.Collection
}

func NewTaskRepository(db *mongo.Database) *TaskRepository {
	t := model.Task{}
	return &TaskRepository{
		db:       db,
		taskColl: mgm.NewCollection(db, t.CollectionName()),
	}
}

func (r *TaskRepository) FindByID(ctx context.Context, id string) (model.Task, error) {
	res := r.taskColl.FindOne(ctx, bson.M{"uuid": id})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return model.Task{}, domain.WrapWithNotFoundError(res.Err(), "task not found")
		}
		return model.Task{}, domain.WrapWithInternalError(res.Err(), "failed to find task")
	}

	t := model.Task{}
	if err := res.Decode(&t); err != nil {
		return model.Task{}, domain.WrapWithInternalError(err, "failed to decode task")
	}

	return t, nil
}

func (r *TaskRepository) Find(_ context.Context, limit int64, offset int64) ([]model.Task, error) {
	t := []model.Task{}
	err := r.taskColl.SimpleFind(&t, bson.M{}, options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []model.Task{}, nil
		}
		return nil, domain.WrapWithInternalError(err, "failed to find task")
	}

	return t, nil
}

func (r *TaskRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.taskColl.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, domain.WrapWithInternalError(err, "failed to count tasks")
	}

	return count, nil
}

func (r *TaskRepository) Create(_ context.Context, t model.Task) (model.Task, error) {
	err := t.Creating()
	if err != nil {
		return model.Task{}, domain.WrapWithInternalError(err, "failed to trigger create hooks for task")
	}

	t.UUID = uuid.New().String()
	err = r.taskColl.Create(&t)
	if err != nil {
		return model.Task{}, domain.WrapWithInternalError(err, "failed to create new task")
	}

	return t, nil
}

func (r *TaskRepository) Update(_ context.Context, t model.Task) (model.Task, error) {
	err := t.Saving()
	if err != nil {
		return model.Task{}, domain.WrapWithInternalError(err, "failed to trigger save hooks for task")
	}

	err = r.taskColl.Update(&t)
	if err != nil {
		return model.Task{}, domain.WrapWithInternalError(err, "failed to create new task")
	}

	return t, nil
}
