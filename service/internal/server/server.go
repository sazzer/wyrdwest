package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sirupsen/logrus"
)

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

// AddRoutes will add a new router onto the server
func (server *Server) AddRoutes(base string, router http.Handler) {
	server.router.Mount(base, router)
}

// Start will start the server listening
func (server *Server) Start(port int) error {
	logrus.WithField("port", port).Info("Starting server...")
	address := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(address, server.router)
}
