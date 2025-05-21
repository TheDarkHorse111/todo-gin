package todo_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"to-do-gin/internal/entity"
	"to-do-gin/internal/mapper/todo"
	"to-do-gin/internal/model"
)

var mapper = todo.NewMapper()

func TestNewMapper(t *testing.T) {
	testMapper := todo.NewMapper()
	assert.NotNil(t, testMapper)
}

func TestMapper_ToModel(t *testing.T) {
	e := entity.Todo{
		ID:          0,
		Name:        "todo",
		Description: "todoDesc",
	}

	m := mapper.ToModel(&e)
	assert.Equal(t, m.ID, e.ID)
	assert.Equal(t, m.Name, e.Name)
	assert.Equal(t, m.Description, e.Description)
}

func TestMapper_ToEntity(t *testing.T) {
	m := model.Todo{
		ID:          0,
		Name:        "todo",
		Description: "todoDesc",
	}

	e := mapper.ToEntity(&m)
	assert.Equal(t, e.ID, m.ID)
	assert.Equal(t, e.Name, m.Name)
	assert.Equal(t, e.Description, m.Description)
}
