package mapper

import (
	"github.com/reversersed/taskservice/internal/application/query"
	"github.com/reversersed/taskservice/internal/domain/entities"
)

func FromEntityToResult(task entities.Task) *query.TaskQueryResult {
	return &query.TaskQueryResult{Result: task}
}
func FromListEntityToResult(task []entities.Task) *query.TaskQueryResultList {
	return &query.TaskQueryResultList{Result: task}
}
