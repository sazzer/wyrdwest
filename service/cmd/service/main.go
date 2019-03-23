package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sazzer/wyrdwest/service/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	server := server.New()

	server.Register(func(e *echo.Echo) {
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
	})

	if err := server.Start(3000); err != nil {
		logrus.WithError(err).Error("Failed to start server")
	}
}
