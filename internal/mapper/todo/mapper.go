package todo

import (
	"to-do-gin/internal/entity"
	"to-do-gin/internal/model"
)

type Mapper interface {
	ToModel(todoEntity *entity.Todo) *model.Todo
	ToEntity(todo *model.Todo) *entity.Todo
}

type mapper struct{}

func NewMapper() Mapper {
	return &mapper{}
}

func (m *mapper) ToModel(todoEntity *entity.Todo) *model.Todo {
	todo := model.Todo{}

	todo.ID = todoEntity.ID
	todo.Name = todoEntity.Name
	todo.Description = todoEntity.Description

	return &todo
}

func (m *mapper) ToEntity(todo *model.Todo) *entity.Todo {
	todoEntity := entity.Todo{}
	todoEntity.ID = todo.ID
	todoEntity.Name = todo.Name
	todoEntity.Description = todo.Description

	return &todoEntity
}
