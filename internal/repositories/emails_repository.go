package repositories

import (
	"context"
	"fmt"
	"sync"

	"github.com/DKhorkov/libs/db"
	"github.com/DKhorkov/libs/logging"
	"github.com/DKhorkov/libs/tracing"

	sq "github.com/Masterminds/squirrel"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

const (
	selectAllColumns       = "*"
	emailsTableName        = "emails"
	idColumnName           = "id"
	userIDColumnName       = "user_id"
	emailEmailColumnName   = "email"
	emailContentColumnName = "content"
	emailSentAtColumnName  = "sent_at"
	returningIDSuffix      = "RETURNING id"
	DESC                   = "DESC"
	ASC                    = "ASC"
)

type EmailsRepository struct {
	dbConnector   db.Connector
	logger        logging.Logger
	traceProvider tracing.Provider
	spanConfig    tracing.SpanConfig
	mutex         *sync.RWMutex
}

func NewEmailsRepository(
	dbConnector db.Connector,
	logger logging.Logger,
	traceProvider tracing.Provider,
	spanConfig tracing.SpanConfig,
) *EmailsRepository {
	return &EmailsRepository{
		dbConnector:   dbConnector,
		logger:        logger,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
		mutex:         new(sync.RWMutex),
	}
}

func (repo *EmailsRepository) GetUserCommunications(
	ctx context.Context,
	userID uint64,
) ([]entities.Email, error) {
	ctx, span := repo.traceProvider.Span(ctx, tracing.CallerName(tracing.DefaultSkipLevel))
	defer span.End()

	span.AddEvent(repo.spanConfig.Events.Start.Name, repo.spanConfig.Events.Start.Opts...)
	defer span.AddEvent(repo.spanConfig.Events.End.Name, repo.spanConfig.Events.End.Opts...)

	connection, err := repo.dbConnector.Connection(ctx)
	if err != nil {
		return nil, err
	}

	defer db.CloseConnectionContext(ctx, connection, repo.logger)

	stmt, params, err := sq.
		Select(selectAllColumns).
		From(emailsTableName).
		Where(sq.Eq{userIDColumnName: userID}).
		OrderBy(fmt.Sprintf("%s %s", idColumnName, DESC)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	// Using mutex for concurrent-safety purpose of using via workers:
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	rows, err := connection.QueryContext(
		ctx,
		stmt,
		params...,
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

	return emails, nil
}

func (repo *EmailsRepository) SaveCommunication(
	ctx context.Context,
	email entities.Email,
) (uint64, error) {
	ctx, span := repo.traceProvider.Span(ctx, tracing.CallerName(tracing.DefaultSkipLevel))
	defer span.End()

	span.AddEvent(repo.spanConfig.Events.Start.Name, repo.spanConfig.Events.Start.Opts...)
	defer span.AddEvent(repo.spanConfig.Events.End.Name, repo.spanConfig.Events.End.Opts...)

	connection, err := repo.dbConnector.Connection(ctx)
	if err != nil {
		return 0, err
	}

	defer db.CloseConnectionContext(ctx, connection, repo.logger)

	stmt, params, err := sq.
		Insert(emailsTableName).
		Columns(
			userIDColumnName,
			emailEmailColumnName,
			emailContentColumnName,
			emailSentAtColumnName,
		).
		Values(
			email.UserID,
			email.Email,
			email.Content,
			email.SentAt,
		).
		Suffix(returningIDSuffix).
		PlaceholderFormat(sq.Dollar). // pq postgres driver works only with $ placeholders
		ToSql()
	if err != nil {
		return 0, err
	}

	// Using mutex for concurrent-safety purpose of using via workers:
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	var emailCommunicationID uint64
	if err = connection.QueryRowContext(ctx, stmt, params...).Scan(&emailCommunicationID); err != nil {
		return 0, err
	}

	return emailCommunicationID, nil
}
