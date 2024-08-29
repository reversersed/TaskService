package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/reversersed/taskservice/internal/domain/entities"
	"github.com/reversersed/taskservice/pkg/middleware"
)

func (r *repository) Create(ctx context.Context, title, description string, due time.Time) (entities.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return entities.Task{}, middleware.InternalError("can't begin transaction: %v", err)
	}

	row, err := tx.Query(ctx, "INSERT INTO tasks (Title,Description,Due) VALUES ($1,$2,$3)", title, description, due.Format("2006-01-02 15:04:05.0000"))
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.Task{}, middleware.InternalError("can't query: %v", err)
	} else if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return entities.Task{}, middleware.NotFoundError("no tasks found: %v", err)
	}

	created_task, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Task])
	if err != nil {
		return entities.Task{}, middleware.NotFoundError("no tasks found: %v", err)
	}
	tx.Commit(ctx)
	return fromDBToEntity(created_task), nil
}
func (r *repository) Update(ctx context.Context, entity entities.Task) (entities.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return entities.Task{}, middleware.InternalError("can't begin transaction: %v", err)
	}
	task := fromEntityToDB(entity)
	row, err := tx.Query(ctx, "UPDATE tasks SET Title=$1,Description=$2,Due=$3 WHERE Id=$4 RETURNING *", task.Title, task.Description, task.Due.Format("2006-01-02 15:04:05.0000"), task.Id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.Task{}, middleware.InternalError("can't query: %v", err)
	} else if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return entities.Task{}, middleware.NotFoundError("no tasks found: %v", err)
	}

	updated_task, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Task])
	if err != nil {
		return entities.Task{}, middleware.NotFoundError("no tasks found: %v", err)
	}
	tx.Commit(ctx)
	return fromDBToEntity(updated_task), nil
}
func (r *repository) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return middleware.InternalError("can't begin transaction: %v", err)
	}
	cmd, err := tx.Exec(ctx, "delete from tasks where Id = $1 limit 1", id)
	if err != nil {
		return middleware.InternalError("can't exec: %v", err)
	}
	if cmd.RowsAffected() == 0 {
		return middleware.NotFoundError("no rows deleted: %v", err)
	}
	tx.Commit(ctx)
	return nil
}
func (r *repository) GetAll(ctx context.Context) ([]entities.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, middleware.InternalError("can't begin transaction: %v", err)
	}
	rows, err := tx.Query(ctx, "select * from tasks")
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, middleware.InternalError("can't query: %v", err)
	} else if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, middleware.NotFoundError("no tasks found: %v", err)
	}
	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[Task])
	if err != nil {
		return nil, middleware.NotFoundError("no tasks found: %v", err)
	}
	entites := make([]entities.Task, len(tasks))
	for i, task := range tasks {
		entites[i] = fromDBToEntity(task)
	}
	tx.Commit(ctx)
	return entites, nil
}
func (r *repository) Get(ctx context.Context, id int) (entities.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return entities.Task{}, middleware.InternalError("can't begin transaction: %v", err)
	}
	row, err := tx.Query(ctx, "select * from tasks where Id = $1 limit 1", id)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.Task{}, middleware.InternalError("can't query: %v", err)
	} else if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return entities.Task{}, middleware.NotFoundError("no tasks found: %v", err)
	}
	task, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Task])
	if err != nil {
		return entities.Task{}, middleware.NotFoundError("no tasks found: %v", err)
	}
	tx.Commit(ctx)
	return fromDBToEntity(task), nil
}
