package repositories

import (
	"context"
	"log/slog"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/libs/logging"

	"github.com/DKhorkov/libs/db"
	"github.com/DKhorkov/libs/tracing"
)

func NewCommonEmailsRepository(
	dbConnector db.Connector,
	logger *slog.Logger,
	traceProvider tracing.TraceProvider,
	spanConfig tracing.SpanConfig,
) *CommonEmailsRepository {
	return &CommonEmailsRepository{
		dbConnector:   dbConnector,
		logger:        logger,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
	}
}

type CommonEmailsRepository struct {
	dbConnector   db.Connector
	logger        *slog.Logger
	traceProvider tracing.TraceProvider
	spanConfig    tracing.SpanConfig
}

func (repo *CommonEmailsRepository) GetUserEmailCommunications(
	ctx context.Context,
	userID uint64,
) ([]entities.Email, error) {
	ctx, span := repo.traceProvider.Span(ctx, tracing.CallerName(tracing.DefaultSkipLevel))
	defer span.End()

	span.AddEvent(repo.spanConfig.Events.Start.Name, repo.spanConfig.Events.Start.Opts...)

	connection, err := repo.dbConnector.Connection(ctx)
	if err != nil {
		return nil, err
	}

	defer db.CloseConnectionContext(ctx, connection, repo.logger)

	rows, err := connection.QueryContext(
		ctx,
		`
			SELECT * 
			FROM emails AS e
			WHERE e.user_id = $1
		`,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err = rows.Close(); err != nil {
			logging.LogErrorContext(
				ctx,
				repo.logger,
				"error during closing SQL rows",
				err,
			)
		}
	}()

	var emails []entities.Email
	for rows.Next() {
		email := entities.Email{}
		columns := db.GetEntityColumns(&email) // Only pointer to use rows.Scan() successfully
		err = rows.Scan(columns...)
		if err != nil {
			return nil, err
		}

		emails = append(emails, email)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	span.AddEvent(repo.spanConfig.Events.End.Name, repo.spanConfig.Events.End.Opts...)
	return emails, nil
}
