package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server represents the actual web server to run
type Server struct {
	server *echo.Echo
}

// New creates a new Web Server
func New() Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return Server{
		server: e,
	}
}

// Start will start the server listening
func (server *Server) Start(port uint16) error {
	address := fmt.Sprintf(":%d", port)
	return server.server.Start(address)
}
