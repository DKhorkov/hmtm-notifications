package main

import (
	"context"

	"github.com/DKhorkov/hmtm-notifications/internal/app"
	"github.com/DKhorkov/hmtm-notifications/internal/config"
	grpccontroller "github.com/DKhorkov/hmtm-notifications/internal/controllers/grpc"
	"github.com/DKhorkov/hmtm-notifications/internal/repositories"
	"github.com/DKhorkov/hmtm-notifications/internal/services"
	"github.com/DKhorkov/hmtm-notifications/internal/usecases"
	"github.com/DKhorkov/libs/db"
	"github.com/DKhorkov/libs/logging"
	"github.com/DKhorkov/libs/tracing"
)

func main() {
	settings := config.New()
	logger := logging.GetInstance(
		settings.Logging.Level,
		settings.Logging.LogFilePath,
	)

	dbConnector, err := db.New(
		db.BuildDsn(settings.Database),
		settings.Database.Driver,
		logger,
		db.WithMaxOpenConnections(settings.Database.Pool.MaxOpenConnections),
		db.WithMaxIdleConnections(settings.Database.Pool.MaxIdleConnections),
		db.WithMaxConnectionLifetime(settings.Database.Pool.MaxConnectionLifetime),
		db.WithMaxConnectionIdleTime(settings.Database.Pool.MaxConnectionIdleTime),
	)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = dbConnector.Close(); err != nil {
			logging.LogError(logger, "Failed to close db connections pool", err)
		}
	}()

	traceProvider, err := tracing.New(settings.Tracing.Server)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = traceProvider.Shutdown(context.Background()); err != nil {
			logging.LogError(logger, "Error shutting down tracer", err)
		}
	}()

	emailsRepository := repositories.NewCommonEmailsRepository(
		dbConnector,
		logger,
		traceProvider,
		settings.Tracing.Spans.Repositories.Notifications,
	)

	emailsService := services.NewCommonEmailsService(
		emailsRepository,
		logger,
	)

	useCases := usecases.NewCommonUseCases(emailsService)
	controller := grpccontroller.New(
		settings.HTTP.Host,
		settings.HTTP.Port,
		useCases,
		logger,
		traceProvider,
		settings.Tracing.Spans.Root,
	)

	application := app.New(controller)
	application.Run()
}
