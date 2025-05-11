package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	todoMapper "to-do-gin/internal/mapper/todo"
	todoRepository "to-do-gin/internal/repository/todo"
	"to-do-gin/internal/service/todo"

	_ "github.com/joho/godotenv/autoload"

	"to-do-gin/internal/database"
)

type Server struct {
	port int

	db          database.Service
	todoService todo.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbService := database.New()
	err := dbService.InitializeDb()
	if err != nil {
		log.Fatal(err)
	}

	mapper := todoMapper.NewMapper()

	repository := todoRepository.NewTodoRepository(dbService, mapper)

	service := todo.NewTodoService(repository)

	NewServer := &Server{
		port:        port,
		db:          dbService,
		todoService: service,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
