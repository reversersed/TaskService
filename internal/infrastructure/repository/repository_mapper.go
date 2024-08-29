package repository

import "github.com/reversersed/taskservice/internal/domain/entities"

func fromDBToEntity(task Task) entities.Task {
	return entities.Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Due:         task.Due,
		Created:     task.Created,
		Updated:     task.Updated,
	}
}
func fromEntityToDB(task entities.Task) Task {
	return Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Due:         task.Due,
		Created:     task.Created,
		Updated:     task.Updated,
	}
}
