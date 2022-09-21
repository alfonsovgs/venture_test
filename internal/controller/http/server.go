package http

import (
	"fmt"

	v1 "github.com/alfonsovgs/venture/internal/controller/http/v1"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Server struct {
	server       *echo.Echo
	dependencies *container
}

func NewServer(dependencies *container) *Server {
	e := echo.New()
	e.Validator = NewValidator(validator.New())

	return &Server{
		server:       e,
		dependencies: dependencies,
	}
}

func (s *Server) Start(port string) {
	s.server.Logger.Fatal(s.server.Start(fmt.Sprintf(":%s", port)))
}

func (s *Server) MapRoutes() {
	v1.NewRouter(s.server)
}
