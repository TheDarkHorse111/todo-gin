package todo

import (
	"context"
	"to-do-gin/internal/database"
	"to-do-gin/internal/mapper/todo"
	"to-do-gin/internal/model"
)

type Repository interface {
	CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	GetTodo(ctx context.Context, todoName string) (*model.Todo, error)
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
	UpdateTodo(ctx context.Context, todo *model.Todo) error
	DeleteTodo(ctx context.Context, todoName string) error
}

type repository struct {
	db     database.Service
	mapper todo.Mapper
}

func NewTodoRepository(db database.Service, mapper todo.Mapper) Repository {
	return &repository{db, mapper}
}

func (r *repository) CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	entity := r.mapper.ToEntity(todo)
	err := r.db.CreateTodo(ctx, entity)
	if err != nil {
		return nil, err
	}
	return r.mapper.ToModel(entity), nil
}

func (r *repository) GetTodo(ctx context.Context, todoName string) (*model.Todo, error) {
	todoEntity, err := r.db.GetTodo(ctx, todoName)
	if err != nil {
		return nil, err
	}

	return r.mapper.ToModel(todoEntity), nil
}

func (r *repository) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	todoEntities, err := r.db.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}
	var todos []*model.Todo
	for _, entity := range todoEntities {
		todos = append(todos, r.mapper.ToModel(&entity))
	}

	return todos, nil
}

func (r *repository) UpdateTodo(ctx context.Context, todo *model.Todo) error {
	err := r.db.UpdateTodo(ctx, r.mapper.ToEntity(todo))
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTodo(ctx context.Context, todoName string) error {
	err := r.db.DeleteTodo(ctx, todoName)
	if err != nil {
		return err
	}
	return nil
}
