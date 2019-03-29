package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"

	"github.com/labstack/echo/v4"
)

// HandlerRegistrationFunc represents a function that can be used to register handlers with the web server
type HandlerRegistrationFunc = func(*echo.Echo)

// Server represents the actual web server to run
type Server struct {
	router *chi.Mux
}

// New creates a new Web Server
func New() Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	return Server{
		router: r,
	}
}

// Register allows us to register handlers with the server
func (server *Server) Register(handlers HandlerRegistrationFunc) {
}

// AddRoutes will add a new router onto the server
func (server *Server) AddRoutes(router http.Handler) {
	server.router.Mount("/", router)
}

// Start will start the server listening
func (server *Server) Start(port int) error {
	logrus.WithField("port", port).Info("Starting server...")
	address := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(address, server.router)
}
