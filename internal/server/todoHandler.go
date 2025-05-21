package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"to-do-gin/internal/model"
)

func (s *Server) GetTodo(ctx *gin.Context) {
	value := ctx.Param("todoName")
	todo, err := s.todoService.GetTodo(ctx.Request.Context(), value)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"todo": todo})
}

func (s *Server) CreateTodo(ctx *gin.Context) {
	var todo model.Todo
	err := ctx.ShouldBindBodyWithJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		log.Println(err)
		return
	}
	createTodo, err := s.todoService.CreateTodo(ctx.Request.Context(), &todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create todo"})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"todo": createTodo})
}

func (s *Server) GetAllTodos(ctx *gin.Context) {
	todos, err := s.todoService.GetAllTodos(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not get todos"})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"todos": todos})

}

func (s *Server) UpdateTodo(ctx *gin.Context) {
	var todo model.Todo
	err := ctx.ShouldBindBodyWithJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		log.Println(err)
		return
	}
	err = s.todoService.UpdateTodo(ctx.Request.Context(), &todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not update todo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "todo updated", "todo": todo})
}

func (s *Server) DeleteTodo(ctx *gin.Context) {
	todoName := ctx.Param("todoName")
	err := s.todoService.DeleteTodo(ctx.Request.Context(), todoName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete todo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("deleted todo %s", todoName)})
}
