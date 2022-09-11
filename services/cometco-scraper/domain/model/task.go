package model

import (
	"github.com/kamva/mgm/v3"
)

type TaskStatus string

const (
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

type Task struct {
	mgm.DefaultModel `bson:",inline"`
	UUID             string     `bson:"uuid"`
	Status           TaskStatus `json:"status" bson:"status"`
	FailReason       string     `json:"fail_reason" bson:"fail_reason"`
	ResultID         string     `json:"result_id" bson:"result_id"`
}

func (a *Task) CollectionName() string {
	return "tasks"
}
