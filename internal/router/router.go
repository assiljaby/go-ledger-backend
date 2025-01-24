package router

import (
	"log/slog"
	"net/http"

	"github.com/assiljaby/go-ledger-backend/pkg/logger"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	RegisterMainRoutes(e)
}

func RegisterMainRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		logger.Info("Hello", slog.String("gitgud", "scrublord"))

		return c.String(http.StatusOK, "Hello, Echo!")
	})
}
