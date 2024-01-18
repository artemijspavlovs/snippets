package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/artemijspavlovs/snippets/demo/internal/api/rest/models/tasks"
	"github.com/artemijspavlovs/snippets/demo/internal/api/rest/models/users"
)

type Server struct {
	db     *sqlx.DB
	Router *echo.Echo
}

func New(db *sqlx.DB) *Server {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	server := &Server{db, router}
	server.SetupRoutes(db)

	return server
}

func (s *Server) SetupRoutes(db *sqlx.DB) {
	taskHandlers := tasks.NewHandlers(db)
	userHandlers := users.NewHandlers(db)

	// API docs
	s.Router.GET("/swagger/*", echoSwagger.WrapHandler)

	// TODO: create specific API endpoints for adding assignees or dependencies,
	// POST /tasks/:id/assignees and POST /tasks/:id/dependencies same for removing or updating these
	s.Router.GET("/tasks", taskHandlers.GetAllTasks)
	s.Router.POST("/tasks", taskHandlers.CreateTask)
	s.Router.GET("/tasks/:id", taskHandlers.GetTaskByID)
	s.Router.PUT("/tasks/:id", taskHandlers.UpdateTask)

	s.Router.POST("/users", userHandlers.CreateUser)
	s.Router.GET("/users/:id", userHandlers.GetUser)
}
