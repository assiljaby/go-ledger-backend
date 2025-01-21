package main

import (
	"log/slog"
	"net/http"

	"github.com/assiljaby/go-ledger-backend/pkg/logger"
	"github.com/labstack/echo/v4"
)

func main() {
	logger.StartLogger()
	e := echo.New()
	e.HidePort = true

	e.GET("/", func(c echo.Context) error {
		logger.Info("Hello", slog.String("gitgud", "scrublord"))

		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
