package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/assiljaby/go-ledger-backend/environments"
	"github.com/assiljaby/go-ledger-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var (
	e    *echo.Echo
	once sync.Once
)

func GetInstance() *echo.Echo {
	once.Do(func() {
		e = echo.New()
	})

	return e
}

func Start(e *echo.Echo) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	e.Logger.SetLevel(log.INFO)
	e.HideBanner = true

	go func() {
		err := e.Start(environments.GetServer().Port)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal(err)
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
