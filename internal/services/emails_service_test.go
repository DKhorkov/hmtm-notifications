package services_test

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/services"
	mockrepositories "github.com/DKhorkov/hmtm-notifications/mocks/repositories"
)

var (
	userID uint64 = 1
)

func TestEmailsServiceGetUserEmailCommunications(t *testing.T) {
	t.Run("get user email communications with existing email communications", func(t *testing.T) {
		expectedEmailCommunications := []entities.Email{
			{
				ID:      1,
				UserID:  userID,
				Email:   "someTestEmail@gmail.com",
				Content: "some test content",
				SentAt:  time.Now().UTC(),
			},
		}

		mockController := gomock.NewController(t)
		emailsRepository := mockrepositories.NewMockEmailsRepository(mockController)
		emailsRepository.
			EXPECT().
			GetUserCommunications(gomock.Any(), userID).
			Return(expectedEmailCommunications, nil).
			MaxTimes(1)

		logger := slog.New(slog.NewJSONHandler(bytes.NewBuffer(make([]byte, 1000)), nil))
		emailsService := services.NewEmailsService(emailsRepository, logger)
		ctx := context.Background()

		emailCommunications, err := emailsService.GetUserCommunications(ctx, userID)
		require.NoError(t, err)
		assert.Len(t, emailCommunications, len(expectedEmailCommunications))
		assert.Equal(t, expectedEmailCommunications, emailCommunications)
	})

	t.Run("get user email communications with existing email communications", func(t *testing.T) {
		mockController := gomock.NewController(t)
		emailsRepository := mockrepositories.NewMockEmailsRepository(mockController)
		emailsRepository.
			EXPECT().
			GetUserCommunications(gomock.Any(), userID).
			Return([]entities.Email{}, nil).
			MaxTimes(1)

		logger := slog.New(slog.NewJSONHandler(bytes.NewBuffer(make([]byte, 1000)), nil))
		emailsService := services.NewEmailsService(emailsRepository, logger)
		ctx := context.Background()

		emailCommunications, err := emailsService.GetUserCommunications(ctx, userID)
		require.NoError(t, err)
		assert.Empty(t, emailCommunications)
	})

	t.Run("get user email communications fail", func(t *testing.T) {
		mockController := gomock.NewController(t)
		emailsRepository := mockrepositories.NewMockEmailsRepository(mockController)
		emailsRepository.
			EXPECT().
			GetUserCommunications(gomock.Any(), userID).
			Return(nil, errors.New("some error")).
			MaxTimes(1)

		logger := slog.New(slog.NewJSONHandler(bytes.NewBuffer(make([]byte, 1000)), nil))
		emailsService := services.NewEmailsService(emailsRepository, logger)
		ctx := context.Background()

		emailCommunications, err := emailsService.GetUserCommunications(ctx, userID)
		require.Error(t, err)
		assert.Nil(t, emailCommunications)
	})
}
