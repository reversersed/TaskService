package mapper

import (
	"github.com/reversersed/taskservice/internal/domain/entities"
	"github.com/reversersed/taskservice/internal/interface/api/rest/dto/response"
)

func ToTaskResponse(task entities.Task) *response.TaskResponse {
	return &response.TaskResponse{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Due:         task.Due.Format("2006-01-02T15:04:05"),
		Created:     task.Created.Format("2006-01-02T15:04:05"),
		Updated:     task.Updated.Format("2006-01-02T15:04:05"),
	}
}
func ToTaskListResponse(task []entities.Task) []*response.TaskResponse {
	resp := make([]*response.TaskResponse, 0)
	for _, t := range task {
		resp = append(resp, &response.TaskResponse{
			Id:          t.Id,
			Title:       t.Title,
			Description: t.Description,
			Due:         t.Due.Format("2006-01-02T15:04:05"),
			Created:     t.Created.Format("2006-01-02T15:04:05"),
			Updated:     t.Updated.Format("2006-01-02T15:04:05"),
		})
	}
	return resp
}
