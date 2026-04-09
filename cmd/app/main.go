package main

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

	"github.com/Rephobia/green-api-test-task/internal/assets"
	"github.com/Rephobia/green-api-test-task/internal/config"
	"github.com/Rephobia/green-api-test-task/internal/loggerx"
	"github.com/Rephobia/green-api-test-task/internal/router"
)

const (
	readTimeout     = 5 * time.Second
	writeTimeout    = 10 * time.Second
	idleTimeout     = 60 * time.Second
	shutdownTimeout = 5 * time.Second
)

func main() {
	config, errConfig := config.New()
	if errConfig != nil {
		fatal(slog.Default(), "failder to load config", errConfig)
	}

	logger := loggerx.New(config)

	frontFiles, errAssets := assets.GetFrontendFiles()
	if errAssets != nil {
		fatal(logger, "failed to load assets", errAssets)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      router.New(logger, *frontFiles),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	gracefulShutdown(logger, srv)
}

func gracefulShutdown(logger *slog.Logger, srv *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info("server started", loggerx.AddrField(srv.Addr))

		errServe := srv.ListenAndServe()
		if errServe != nil && !errors.Is(errServe, http.ErrServerClosed) {
			fatal(logger, "server failed", errServe)
		}
	}()

	<-stop
	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if errShutdown := srv.Shutdown(ctx); errShutdown != nil {
		logger.Error("shutdown error", loggerx.ErrorField(errShutdown))
	}

	logger.Info("server stopped")
}

func fatal(logger *slog.Logger, msg string, err error) {
	logger.Error(msg, "error", err)
	os.Exit(1)
}
