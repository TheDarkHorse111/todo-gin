package database

import (
	"context"
	"to-do-gin/internal/entity"
)

func (s *service) CreateTodo(ctx context.Context, todo *entity.Todo) error {
	query := `
				INSERT INTO todo (name, description)
				VALUES ($1, $2)
				RETURNING id
			`

	preparedStatement, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	err = preparedStatement.QueryRowContext(ctx, todo.Name, todo.Description).Scan(&todo.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetTodo(ctx context.Context, todoName string) (*entity.Todo, error) {
	query := `
			SELECT * FROM todo WHERE name = $1
			`
	row := s.db.QueryRowContext(ctx, query, todoName)
	var todo entity.Todo
	err := row.Scan(&todo.ID, &todo.Name, &todo.Description)
	if err != nil {
		return nil, err
	}
	return &todo, nil

}
func (s *service) GetAllTodos(ctx context.Context) ([]*entity.Todo, error) {
	query := `
			SELECT * FROM todo
			`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*entity.Todo
	for rows.Next() {
		var todo entity.Todo
		err = rows.Scan(&todo.ID, &todo.Name, &todo.Description)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (s *service) UpdateTodo(ctx context.Context, todo *entity.Todo) error {
	query := `
UPDATE todo SET name = $1, description = $2 WHERE id = $3
`
	_, err := s.db.ExecContext(ctx, query, todo.Name, todo.Description, todo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteTodo(ctx context.Context, todoName string) error {
	query := `
DELETE FROM todo WHERE name = $1
`
	_, err := s.db.ExecContext(ctx, query, todoName)
	if err != nil {
		return err
	}
	return nil
}
