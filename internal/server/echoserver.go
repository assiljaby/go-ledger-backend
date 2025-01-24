package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/assiljaby/go-ledger-backend/environments"
	"github.com/assiljaby/go-ledger-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type EchoServer struct {
	Server *echo.Echo
}

func New() *EchoServer {
	return &EchoServer{
		Server: echo.New(),
	}
}

func (e *EchoServer) Start() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	e.Server.Logger.SetLevel(log.INFO)
	e.Server.HideBanner = true

	go func() {
		err := e.Server.Start(environments.GetServer().Port)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Server.Logger.Fatal(err)
		}
	}()

	<-ctx.Done()
	logger.Info("server shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	errChan := make(chan error, 1)

	go func() {
		if err := e.Server.Shutdown(shutdownCtx); err != nil {
			errChan <- fmt.Errorf("server couldn't shutdown: %w", err)

			return
		}
		errChan <- nil
	}()

	if err := <-errChan; err != nil {
		logger.Error("error stopping server", slog.String("err", err.Error()))
	}

	logger.Info("server gracefully stopped")
}
