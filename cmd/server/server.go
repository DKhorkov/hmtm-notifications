package main

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"

	"github.com/DKhorkov/libs/db"
	"github.com/DKhorkov/libs/logging"
	customnats "github.com/DKhorkov/libs/nats"
	"github.com/DKhorkov/libs/tracing"

	"github.com/DKhorkov/hmtm-notifications/internal/app"
	ssogrpcclient "github.com/DKhorkov/hmtm-notifications/internal/clients/sso/grpc"
	"github.com/DKhorkov/hmtm-notifications/internal/config"
	"github.com/DKhorkov/hmtm-notifications/internal/contentbuilders"
	grpccontroller "github.com/DKhorkov/hmtm-notifications/internal/controllers/grpc"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
	"github.com/DKhorkov/hmtm-notifications/internal/repositories"
	"github.com/DKhorkov/hmtm-notifications/internal/senders"
	"github.com/DKhorkov/hmtm-notifications/internal/services"
	"github.com/DKhorkov/hmtm-notifications/internal/usecases"
	"github.com/DKhorkov/hmtm-notifications/internal/workers/handlers/builders"
)

func main() {
	settings := config.New()
	logger := logging.New(
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

	ssoClient, err := ssogrpcclient.New(
		settings.Clients.SSO.Host,
		settings.Clients.SSO.Port,
		settings.Clients.SSO.RetriesCount,
		settings.Clients.SSO.RetryTimeout,
		logger,
		traceProvider,
		settings.Tracing.Spans.Clients.SSO,
	)

	if err != nil {
		panic(err)
	}

	ssoRepository := repositories.NewSsoRepository(ssoClient)
	ssoService := services.NewSsoService(
		ssoRepository,
		logger,
	)

	emailsRepository := repositories.NewEmailsRepository(
		dbConnector,
		logger,
		traceProvider,
		settings.Tracing.Spans.Repositories.Emails,
	)

	emailsService := services.NewEmailsService(
		emailsRepository,
		logger,
	)

	contentBuilders := interfaces.ContentBuilders{
		VerifyEmail: contentbuilders.NewVerifyEmailContentBuilder(
			settings.Email.VerifyEmailURL,
		),
		ForgetPassword: contentbuilders.NewForgetPasswordContentBuilder(
			settings.Email.ForgetPasswordURL,
		),
	}

	communicationsSenders := interfaces.Senders{
		Email: senders.NewEmailSender(
			settings.Email.SMTP,
			traceProvider,
			settings.Tracing.Spans.Senders.Email,
		),
	}

	useCases := usecases.New(
		emailsService,
		ssoService,
		contentBuilders,
		communicationsSenders,
	)

	controller := grpccontroller.New(
		settings.HTTP.Host,
		settings.HTTP.Port,
		useCases,
		logger,
		traceProvider,
		settings.Tracing.Spans.Root,
	)

	verifyEmailWorker, err := customnats.NewWorker(
		settings.NATS.ClientURL,
		settings.NATS.Subjects.VerifyEmail,
		customnats.WithGoroutinesPoolSize(settings.NATS.GoroutinesPoolSize),
		customnats.WithMessageChannelBufferSize(settings.NATS.MessageChannelBufferSize),
		customnats.WithNatsOptions(nats.Name(settings.NATS.Workers.VerifyEmail.Name)),
		customnats.WithMessageHandler(
			builders.NewVerifyEmailBuilder(
				useCases,
				traceProvider,
				settings.Tracing.Spans.Handlers.VerifyEmail,
				logger,
			).MessageHandler(),
		),
	)

	if err != nil {
		panic(err)
	}

	if err = verifyEmailWorker.Run(); err != nil {
		panic(err)
	}

	defer func() {
		if err = verifyEmailWorker.Stop(); err != nil {
			logging.LogError(
				logger,
				fmt.Sprintf("Error shutting down \"%s\" worker", settings.NATS.Workers.VerifyEmail.Name),
				err,
			)
		}
	}()

	forgetPasswordWorker, err := customnats.NewWorker(
		settings.NATS.ClientURL,
		settings.NATS.Subjects.ForgetPassword,
		customnats.WithGoroutinesPoolSize(settings.NATS.GoroutinesPoolSize),
		customnats.WithMessageChannelBufferSize(settings.NATS.MessageChannelBufferSize),
		customnats.WithNatsOptions(nats.Name(settings.NATS.Workers.ForgetPassword.Name)),
		customnats.WithMessageHandler(
			builders.NewForgetPasswordBuilder(
				useCases,
				traceProvider,
				settings.Tracing.Spans.Handlers.ForgetPassword,
				logger,
			).MessageHandler(),
		),
	)

	if err != nil {
		panic(err)
	}

	if err = forgetPasswordWorker.Run(); err != nil {
		panic(err)
	}

	defer func() {
		if err = forgetPasswordWorker.Stop(); err != nil {
			logging.LogError(
				logger,
				fmt.Sprintf("Error shutting down \"%s\" worker", settings.NATS.Workers.ForgetPassword.Name),
				err,
			)
		}
	}()

	application := app.New(controller)
	application.Run()
}
