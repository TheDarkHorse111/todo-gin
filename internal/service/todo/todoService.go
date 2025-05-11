package todo

import (
	"context"
	"to-do-gin/internal/model"
	"to-do-gin/internal/repository/todo"
)

type Service interface {
	CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	GetTodo(ctx context.Context, todoName string) (*model.Todo, error)
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
	UpdateTodo(ctx context.Context, todo *model.Todo) error
	DeleteTodo(ctx context.Context, todoName string) error
}

type service struct {
	todoRepository todo.Repository
}

func NewTodoService(repository todo.Repository) Service {
	return &service{repository}
}

func (s *service) CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	createdTodo, err := s.todoRepository.CreateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}
	return createdTodo, nil
}

func (s *service) GetTodo(ctx context.Context, todoName string) (*model.Todo, error) {
	todoModel, err := s.todoRepository.GetTodo(ctx, todoName)
	if err != nil {
		return nil, err
	}
	return todoModel, nil
}

func (s *service) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := s.todoRepository.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *service) UpdateTodo(ctx context.Context, todo *model.Todo) error {
	err := s.todoRepository.UpdateTodo(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteTodo(ctx context.Context, todoName string) error {
	err := s.todoRepository.DeleteTodo(ctx, todoName)
	if err != nil {
		return err
	}
	return nil
}
