package todo_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"to-do-gin/.gen/mocks/mock_todo"
	"to-do-gin/internal/model"
	"to-do-gin/internal/service/todo"
)

func TestService_CreateTodo(t *testing.T) {
	var mockRepo = mock_todo.NewMockRepository(t)
	var service = todo.NewTodoService(mockRepo)

	todoModel := model.Todo{
		Name:        "todo",
		Description: "todoDesc",
	}

	expected := model.Todo{
		ID:          1,
		Name:        "todo",
		Description: "todoDesc",
	}

	mockRepo.EXPECT().CreateTodo(nil, &todoModel).Return(&expected, nil)

	actual, err := service.CreateTodo(nil, &todoModel)
	if err != nil {
		return
	}

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description)
}

func TestService_CreateTodo_TodoNotFound(t *testing.T) {
	var mockRepo = mock_todo.NewMockRepository(t)
	var service = todo.NewTodoService(mockRepo)

	var todoModel *model.Todo

	expectedError := errors.New("could not find todo")
	mockRepo.EXPECT().CreateTodo(nil, todoModel).Return(nil, expectedError)

	actual, err := service.CreateTodo(nil, todoModel)

	assert.Nil(t, actual)
	if err != nil {
		assert.Equal(t, expectedError.Error(), err.Error())
	}
}
