package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// HandlerRegistrationFunc represents a function that can be used to register handlers with the web server
type HandlerRegistrationFunc = func(*echo.Echo)

// Server represents the actual web server to run
type Server struct {
	server *echo.Echo
}

// New creates a new Web Server
func New() Server {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	return Server{
		server: e,
	}
}

// Register allows us to register handlers with the server
func (server *Server) Register(handlers HandlerRegistrationFunc) {
	handlers(server.server)
}

// Start will start the server listening
func (server *Server) Start(port int) error {
	address := fmt.Sprintf(":%d", port)
	return server.server.Start(address)
}
