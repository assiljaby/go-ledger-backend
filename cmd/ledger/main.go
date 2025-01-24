package main

import (
	"log/slog"
	"net/http"

	"github.com/assiljaby/go-ledger-backend/internal/server"
	"github.com/assiljaby/go-ledger-backend/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	logger.StartLogger()

	_ = godotenv.Load()

	e := server.GetInstance()

	e.GET("/", func(c echo.Context) error {
		logger.Info("Hello", slog.String("gitgud", "scrublord"))

		return c.String(http.StatusOK, "Hello, Echo!")
	})

	server.Start(e)
}
