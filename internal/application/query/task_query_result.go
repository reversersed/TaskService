package query

import "github.com/reversersed/taskservice/internal/domain/entities"

type TaskQueryResult struct {
	Result entities.Task
}
type TaskQueryResultList struct {
	Result []entities.Task
}
