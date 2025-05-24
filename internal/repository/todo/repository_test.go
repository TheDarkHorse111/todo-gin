package todo_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
	"to-do-gin/internal/entity"
	todo2 "to-do-gin/internal/mapper/todo"
	"to-do-gin/internal/model"
	"to-do-gin/internal/repository/todo"
)

var todos []*entity.Todo
var repository = todo.NewTodoRepository(dbMock{}, todo2.NewMapper())

type dbMock struct{}

func (d dbMock) InitializeDb() error {
	panic("implement me")
}

func (d dbMock) Health() map[string]string {
	panic("implement me")
}

func (d dbMock) Close() error {
	panic("implement me")
}

func (d dbMock) CreateTodo(ctx context.Context, todo *entity.Todo) error {
	todos = append(todos, todo)
	return nil
}

func (d dbMock) GetTodo(ctx context.Context, todoName string) (*entity.Todo, error) {
	index := getTodoIdxByName(todoName)

	if index == -1 {
		return nil, errors.New("could not find todo with name: " + todoName + "")
	}

	return todos[index], nil
}

func (d dbMock) GetAllTodos(ctx context.Context) ([]*entity.Todo, error) {
	if todos == nil {
		return nil, errors.New("could not find todos")
	}
	return todos, nil
}

func (d dbMock) UpdateTodo(ctx context.Context, todo *entity.Todo) error {
	index := slices.IndexFunc(todos, func(oldTodo *entity.Todo) bool {
		return oldTodo.ID == todo.ID
	})

	if index == -1 {
		return errors.New("could not find todo with name: " + todo.Name + "")
	}

	todos[index] = todo
	return nil
}

func (d dbMock) DeleteTodo(ctx context.Context, todoName string) error {
	todoIdx := getTodoIdxByName(todoName)

	if todoIdx == -1 {
		return errors.New("could not find todo with name: " + todoName)
	}

	todos = slices.Delete(todos, todoIdx, todoIdx+1)
	return nil
}

func getTodoIdxByName(todoName string) int {
	return slices.IndexFunc(todos, func(todo *entity.Todo) bool {
		return todo.Name == todoName
	})
}

func TestRepository_CreateTodo(t *testing.T) {
	todoModel := model.Todo{
		Name:        "todo",
		Description: "todoDesc",
	}

	createdTodo := createTodo(t, todoModel)
	if createdTodo == nil {
		return
	}

	assert.NotEmpty(t, todos)
	assert.Equal(t, todoModel.ID, createdTodo.ID)
	assert.Equal(t, todoModel.Name, createdTodo.Name)
	assert.Equal(t, todoModel.Description, createdTodo.Description)
}

func createTodo(t *testing.T, todoModel model.Todo) *model.Todo {
	createdTodo, err := repository.CreateTodo(context.Background(), &todoModel)
	if err != nil {
		assert.Equal(t, err.Error(), "could not find todoModel with name: "+todoModel.Name)
		return nil
	}
	return createdTodo
}

func TestRepository_GetTodo(t *testing.T) {

	todoModel := model.Todo{
		Name:        "todo",
		Description: "todoDesc",
	}
	expected := createTodo(t, todoModel)
	actual, _ := repository.GetTodo(context.Background(), todoModel.Name)

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description)
}

func TestRepository_GetTodo_ReturnError(t *testing.T) {

	todoModel := model.Todo{
		Name:        "todo",
		Description: "todoDesc",
	}

	_, err := repository.GetTodo(context.Background(), todoModel.Name)
	if err != nil {
		assert.Equal(t, err.Error(), "could not find todo with name: "+todoModel.Name)
		return
	}
}

func TestRepository_GetAllTodos(t *testing.T) {
	todos = make([]*entity.Todo, 0)
	todoModel1 := model.Todo{
		Name:        "todo1",
		Description: "todoDesc1",
	}

	todoModel2 := model.Todo{
		Name:        "todo2",
		Description: "todoDesc2",
	}

	createTodo(t, todoModel1)
	createTodo(t, todoModel2)

	allTodos, _ := repository.GetAllTodos(context.Background())

	assert.Equal(t, len(todos), len(allTodos))
	assert.Equal(t, len(allTodos), 2)
	assert.Equal(t, todoModel1, *allTodos[0])
	assert.Equal(t, todoModel2, *allTodos[1])
}

func TestRepository_GetAllTodos_ReturnError(t *testing.T) {
	_, err := repository.GetAllTodos(context.Background())
	if err != nil {
		assert.Equal(t, err.Error(), "could not find todos")
	}
}

func TestRepository_UpdateTodo(t *testing.T) {
	todoModel := model.Todo{
		ID:          0,
		Name:        "todo",
		Description: "todoDesc",
	}

	updatedTodoModel := model.Todo{
		ID:          0,
		Name:        "updatedTodo",
		Description: "updatedTodoDesc",
	}

	createTodo(t, todoModel)

	_ = repository.UpdateTodo(context.Background(), &updatedTodoModel)

	updatedTodo, _ := repository.GetTodo(context.Background(), updatedTodoModel.Name)

	assert.Equal(t, updatedTodo.ID, updatedTodoModel.ID)
	assert.Equal(t, updatedTodo.Name, updatedTodoModel.Name)
	assert.Equal(t, updatedTodo.Description, updatedTodoModel.Description)
}

func TestRepository_DeleteTodo(t *testing.T) {
	todos = make([]*entity.Todo, 0)
	todoModel := model.Todo{
		ID:          0,
		Name:        "todo",
		Description: "todoDesc",
	}

	createTodo(t, todoModel)

	_ = repository.DeleteTodo(context.Background(), todoModel.Name)

	assert.Empty(t, todos)
}
