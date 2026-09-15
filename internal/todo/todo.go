package todo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/subratamondal1029/goTodo/internal/database"
)

type TodoService struct {
	queries database.Queries
	ctx     context.Context
}

type ITodoService interface {
	ReadAll() ([]database.Todo, error)
	ReadOne(id int) (*database.Todo, error)
	Create(title string) (*database.Todo, error)
	ToggleDone(id int, done bool) (*database.Todo, error)
	Delete(id int) error
}

func NewTodoService(queries database.Queries) ITodoService {
	return &TodoService{
		queries: queries,
		ctx:     context.Background(),
	}
}

func (t *TodoService) Create(title string) (*database.Todo, error) {
	todo, err := t.queries.Create(t.ctx, title)

	if err != nil {
		return nil, fmt.Errorf("Error Creating new todo, %w", err)
	}

	return &todo, nil
}

func (t *TodoService) ReadAll() ([]database.Todo, error) {
	todos, err := t.queries.ReadAll(t.ctx)
	if err != nil {
		return nil, fmt.Errorf("Error reading all todos, %w", err)
	}
	return todos, nil
}
func (t *TodoService) ReadOne(id int) (*database.Todo, error) {
	todo, err := t.queries.ReadOne(t.ctx, int64(id))

	if err != nil {
		return nil, fmt.Errorf("Error reading todo, %w", err)
	}

	return &todo, nil
}

func (t *TodoService) ToggleDone(id int, done bool) (*database.Todo, error) {
	todo, err := t.queries.ChangeStatus(t.ctx, database.ChangeStatusParams{
		ID: int64(id),
		Completed: sql.NullBool{
			Valid: true,
			Bool:  done,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("Error toggling todo, %w", err)
	}

	return &todo, nil
}

func (t *TodoService) Delete(id int) error {
	err := t.queries.Delete(t.ctx, int64(id))

	if err != nil {
		return fmt.Errorf("Error deleting todo, %w", err)
	}

	return nil
}
