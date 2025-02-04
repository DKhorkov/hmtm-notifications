package repositories

import (
	"context"
	"log/slog"
	"sync"

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
		mutex:         new(sync.RWMutex),
	}
}

type CommonEmailsRepository struct {
	dbConnector   db.Connector
	logger        *slog.Logger
	traceProvider tracing.TraceProvider
	spanConfig    tracing.SpanConfig
	mutex         *sync.RWMutex
}

func (repo *CommonEmailsRepository) GetUserCommunications(
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

	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
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

func (repo *CommonEmailsRepository) SaveCommunication(ctx context.Context, email entities.Email) (uint64, error) {
	ctx, span := repo.traceProvider.Span(ctx, tracing.CallerName(tracing.DefaultSkipLevel))
	defer span.End()

	span.AddEvent(repo.spanConfig.Events.Start.Name, repo.spanConfig.Events.Start.Opts...)

	connection, err := repo.dbConnector.Connection(ctx)
	if err != nil {
		return 0, err
	}

	defer db.CloseConnectionContext(ctx, connection, repo.logger)

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	var emailCommunicationID uint64
	err = connection.QueryRowContext(
		ctx,
		`
			INSERT INTO emails (user_id, email, content, sent_at) 
			VALUES ($1, $2, $3, $4)
			RETURNING emails.id
		`,
		email.UserID,
		email.Email,
		email.Content,
		email.SentAt,
	).Scan(&emailCommunicationID)

	if err != nil {
		return 0, err
	}

	span.AddEvent(repo.spanConfig.Events.End.Name, repo.spanConfig.Events.End.Opts...)
	return emailCommunicationID, nil
}
