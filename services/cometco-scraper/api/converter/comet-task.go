package converter

import (
	httpModel "scraping_challenge/services/cometco-scraper/api/model"
	"scraping_challenge/services/cometco-scraper/domain/model"
)

func ToGetTaskResponse(task model.Task) httpModel.GetTaskResponse {
	return httpModel.GetTaskResponse{
		ID:         task.UUID,
		Status:     ToTaskStatus(task.Status),
		FailReason: task.FailReason,
		ResultID:   task.ResultID,
	}
}

func ToGetTasksResponse(tasks []model.Task) httpModel.GetTasksResponse {
	resTasks := make([]httpModel.GetShortTaskResponse, len(tasks))
	for i, task := range tasks {
		resTasks[i] = ToGetShortTaskResponse(task)
	}

	return httpModel.GetTasksResponse{
		Tasks: resTasks,
	}
}

func ToGetShortTaskResponse(task model.Task) httpModel.GetShortTaskResponse {
	return httpModel.GetShortTaskResponse{
		ID: task.UUID,
	}
}

func ToTaskStatus(status model.TaskStatus) httpModel.TaskStatus {
	switch status {
	case model.TaskStatusInProgress:
		return httpModel.TaskStatusInProgress
	case model.TaskStatusCompleted:
		return httpModel.TaskStatusCompleted
	case model.TaskStatusFailed:
		return httpModel.TaskStatusFailed
	}

	return httpModel.TaskStatusInProgress
}
