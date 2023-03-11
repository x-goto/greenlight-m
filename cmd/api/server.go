package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) Shutdown(shutdownError chan error) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit

	app.logger.LogInfo("shutting down server", map[string]string{
		"signal": s.String(),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	shutdownError <- app.server.Shutdown(ctx)
}

func (app *application) serve() error {
	app.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go app.Shutdown(shutdownError)

	app.logger.LogInfo("server starting", map[string]string{
		"addr": app.server.Addr,
		"env":  app.config.env,
	})

	err := app.server.ListenAndServe()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	if err = <-shutdownError; err != nil {
		return err
	}

	app.logger.LogInfo("stopped server", map[string]string{
		"addr": app.server.Addr,
	})

	return nil
}
