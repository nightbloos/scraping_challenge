package model

type TaskStatus string

const (
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

type CreateTaskRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CreateTaskResponse struct {
	ID string `json:"id"`
}

type GetTaskListQuery struct {
	Limit  int64 `form:"limit"`
	Offset int64 `form:"offset"`
}

type GetTasksResponse struct {
	Tasks []GetShortTaskResponse `json:"tasks"`
}

type GetShortTaskResponse struct {
	ID string `json:"id"`
}

type GetTaskResponse struct {
	ID         string     `json:"id"`
	Status     TaskStatus `json:"status"`
	FailReason string     `json:"fail_reason,omitempty"`
	ResultID   string     `json:"result_id,omitempty"`
}
