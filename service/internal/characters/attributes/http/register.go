package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sazzer/wyrdwest/service/internal/characters/attributes"
	"github.com/sazzer/wyrdwest/service/internal/server"
)

// RegisterAttributes will return the means to register the Attributes Handler with the HTTP Server
func RegisterAttributes(dao attributes.Retriever) server.HandlerRegistrationFunc {
	return func(e *echo.Echo) {
		e.GET("/attributes/:id", func(c echo.Context) error { return getByID(c, dao) })
	}
}
