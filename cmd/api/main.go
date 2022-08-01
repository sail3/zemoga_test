package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/sail3/zemoga_test/internal/config"
	"github.com/sail3/zemoga_test/internal/db/postgres"
	"github.com/sail3/zemoga_test/internal/transport"
	"github.com/sail3/zemoga_test/pkg/health"
	"github.com/sail3/zemoga_test/pkg/portfolio"
	"github.com/sail3/zemoga_test/pkg/twitter"

	"github.com/sail3/zemoga_test/internal/logger"
)

const migrationsRootFolder = "file://migrations"

func main() {
	conf := config.New()

	l := logger.New("zemoga_test service", conf.Debug)

	err := doMigrate(conf.DbPostgresUrl)
	if err != nil {
		l.With("transport", "http").Error(err)
		os.Exit(1)
	}

	db := postgres.NewPostgresClient(conf.DbPostgresUrl)
	err = db.Ping()
	if err != nil {
		l.Error(err)
		os.Exit(1)
	}
	hSvc := health.NewService(
		l.With("scope", "health service"),
	)
	pRepo := portfolio.NewRepository(db.DB, l)
	tSvc := twitter.NewService(conf)

	pSvc := portfolio.NewService(pRepo, tSvc, l)

	pHand := portfolio.NewHandler(pSvc, l)

	httpTransportRouter := transport.NewHTTPRouter(hSvc, pHand, l)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", conf.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      httpTransportRouter,
	}
	l.
		With("transport", "http").
		With("port", conf.Port).
		Info("Transport Start")

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		l := l.With("transport", "http")
		if err := srv.ListenAndServe(); err != nil {
			l.Error(err)
		}
		l.Info("Transport Stopped")
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	l.Info("Service gracefully shut down")
	os.Exit(0)
}

func doMigrate(databaseURL string) error {
	m, err := migrate.New(
		migrationsRootFolder,
		databaseURL,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
