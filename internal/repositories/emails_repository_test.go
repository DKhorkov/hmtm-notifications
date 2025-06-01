//go:build integration

package repositories_test

import (
	"context"
	"database/sql"
	"os"
	"path"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3" // Must be imported for correct work

	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/DKhorkov/libs/db"
	mocklogging "github.com/DKhorkov/libs/logging/mocks"
	"github.com/DKhorkov/libs/pointers"
	"github.com/DKhorkov/libs/tracing"
	mocktracing "github.com/DKhorkov/libs/tracing/mocks"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/repositories"
)

const (
	driver = "sqlite3"
	//dsn    = "file::memory:?cache=shared"
	dsn              = "../../test.db"
	migrationsDir    = "/migrations"
	gooseZeroVersion = 0
)

func TestEmailsRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(EmailsRepositoryTestSuite))
}

type EmailsRepositoryTestSuite struct {
	suite.Suite

	cwd              string
	ctx              context.Context
	dbConnector      db.Connector
	connection       *sql.Conn
	emailsRepository *repositories.EmailsRepository
	logger           *mocklogging.MockLogger
	traceProvider    *mocktracing.MockProvider
	spanConfig       tracing.SpanConfig
}

func (s *EmailsRepositoryTestSuite) SetupSuite() {
	s.NoError(goose.SetDialect(driver))

	ctrl := gomock.NewController(s.T())
	s.ctx = context.Background()
	s.logger = mocklogging.NewMockLogger(ctrl)
	dbConnector, err := db.New(dsn, driver, s.logger)
	s.NoError(err)

	cwd, err := os.Getwd()
	s.NoError(err)

	s.cwd = cwd
	s.dbConnector = dbConnector
	s.traceProvider = mocktracing.NewMockProvider(ctrl)
	s.spanConfig = tracing.SpanConfig{}
	s.emailsRepository = repositories.NewEmailsRepository(s.dbConnector, s.logger, s.traceProvider, s.spanConfig)
}

func (s *EmailsRepositoryTestSuite) SetupTest() {
	s.NoError(
		goose.Up(
			s.dbConnector.Pool(),
			path.Dir(
				path.Dir(s.cwd),
			)+migrationsDir,
		),
	)

	connection, err := s.dbConnector.Connection(s.ctx)
	s.NoError(err)

	s.connection = connection
}

func (s *EmailsRepositoryTestSuite) TearDownTest() {
	s.NoError(
		goose.DownTo(
			s.dbConnector.Pool(),
			path.Dir(
				path.Dir(s.cwd),
			)+migrationsDir,
			gooseZeroVersion,
		),
	)

	s.NoError(s.connection.Close())
}

func (s *EmailsRepositoryTestSuite) TearDownSuite() {
	s.NoError(s.dbConnector.Close())
}

func (s *EmailsRepositoryTestSuite) TestGetUserCommunicationsWithExistingEmails() {
	s.traceProvider.
		EXPECT().
		Span(gomock.Any(), gomock.Any()).
		Return(context.Background(), mocktracing.NewMockSpan()).
		Times(1)

	id := 1
	userID := uint64(1)
	sentAt := time.Now().UTC()
	_, err := s.connection.ExecContext(
		s.ctx,
		`
			INSERT INTO emails (id, user_id, email, content, sent_at) 
			VALUES ($1, $2, $3, $4, $5)
		`,
		id,
		userID,
		"test@example.com",
		"Test email content",
		sentAt,
	)
	s.NoError(err)

	emails, err := s.emailsRepository.GetUserCommunications(s.ctx, userID, nil)
	s.NoError(err)
	s.NotEmpty(emails)
	s.Equal(1, len(emails))
	s.Equal(userID, emails[0].UserID)
	s.Equal("test@example.com", emails[0].Email)
	s.Equal("Test email content", emails[0].Content)
	s.WithinDuration(sentAt, emails[0].SentAt, time.Second)
}

func (s *EmailsRepositoryTestSuite) TestCountUserCommunications() {
	s.traceProvider.
		EXPECT().
		Span(gomock.Any(), gomock.Any()).
		Return(context.Background(), mocktracing.NewMockSpan()).
		Times(1)

	id := 1
	userID := uint64(1)
	sentAt := time.Now().UTC()
	_, err := s.connection.ExecContext(
		s.ctx,
		`
			INSERT INTO emails (id, user_id, email, content, sent_at) 
			VALUES ($1, $2, $3, $4, $5)
		`,
		id,
		userID,
		"test@example.com",
		"Test email content",
		sentAt,
	)
	s.NoError(err)

	count, err := s.emailsRepository.CountUserCommunications(s.ctx, userID)
	s.NoError(err)
	s.NotZero(count)
	s.Equal(uint64(1), count)
}

func (s *EmailsRepositoryTestSuite) TestGetUserCommunicationsWithExistingEmailsAndPagination() {
	s.traceProvider.
		EXPECT().
		Span(gomock.Any(), gomock.Any()).
		Return(context.Background(), mocktracing.NewMockSpan()).
		Times(1)

	userID := uint64(1)
	sentAt := time.Now().UTC()
	_, err := s.connection.ExecContext(
		s.ctx,
		`
			INSERT INTO emails (id, user_id, email, content, sent_at) 
			VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10), ($11, $12, $13, $14, $15)
		`,
		1, userID, "test@example.com", "Test email content 1", sentAt,
		2, userID, "test@example.com", "Test email content 2", sentAt,
		3, userID, "test@example.com", "Test email content 3", sentAt,
	)
	s.NoError(err)

	pagination := &entities.Pagination{
		Limit:  pointers.New[uint64](1),
		Offset: pointers.New[uint64](1),
	}

	emails, err := s.emailsRepository.GetUserCommunications(s.ctx, userID, pagination)
	s.NoError(err)
	s.NotEmpty(emails)
	s.Equal(1, len(emails))
	s.Equal(userID, emails[0].UserID)
	s.Equal("test@example.com", emails[0].Email)
	s.Equal("Test email content 2", emails[0].Content)
	s.WithinDuration(sentAt, emails[0].SentAt, time.Second)
}

func (s *EmailsRepositoryTestSuite) TestGetUserCommunicationsWithoutExistingEmails() {
	s.traceProvider.
		EXPECT().
		Span(gomock.Any(), gomock.Any()).
		Return(context.Background(), mocktracing.NewMockSpan()).
		Times(1)

	userID := uint64(2)
	emails, err := s.emailsRepository.GetUserCommunications(s.ctx, userID, nil)
	s.NoError(err)
	s.Empty(emails)
}

func (s *EmailsRepositoryTestSuite) TestSaveCommunicationSuccess() {
	s.traceProvider.
		EXPECT().
		Span(gomock.Any(), gomock.Any()).
		Return(context.Background(), mocktracing.NewMockSpan()).
		Times(1)

	email := entities.Email{
		UserID:  3,
		Email:   "new@example.com",
		Content: "New email content",
		SentAt:  time.Now().UTC(),
	}

	// Error and zero id due to returning nil ID after insert operation
	// SQLite inner realization without AUTO_INCREMENT for SERIAL PRIMARY KEY
	id, err := s.emailsRepository.SaveCommunication(s.ctx, email)
	s.Error(err)
	s.Zero(id)
}

func (s *EmailsRepositoryTestSuite) TestSaveCommunicationError() {
	s.traceProvider.
		EXPECT().
		Span(gomock.Any(), gomock.Any()).
		Return(context.Background(), mocktracing.NewMockSpan()).
		Times(1)

	email := entities.Email{
		UserID:  4,
		Email:   "error@example.com",
		Content: "Error case",
		SentAt:  time.Now().UTC(),
	}

	id, err := s.emailsRepository.SaveCommunication(s.ctx, email)
	s.Error(err)
	s.Zero(id)
}
