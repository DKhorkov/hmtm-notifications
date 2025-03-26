package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	loggermock "github.com/DKhorkov/libs/logging/mocks"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/services"
	mockrepositories "github.com/DKhorkov/hmtm-notifications/mocks/repositories"
)

var (
	userID uint64 = 1
)

func TestEmailsService_GetUserEmailCommunications(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name          string
		userID        uint64
		setupMocks    func(emailsRepository *mockrepositories.MockEmailsRepository, logger *loggermock.MockLogger)
		expected      []entities.Email
		errorExpected bool
	}{
		{
			name:   "get user email communications with existing email communications",
			userID: userID,
			expected: []entities.Email{
				{
					ID:      1,
					UserID:  userID,
					Email:   "someTestEmail@gmail.com",
					Content: "some test content",
					SentAt:  now,
				},
			},
			setupMocks: func(emailsRepository *mockrepositories.MockEmailsRepository, _ *loggermock.MockLogger) {
				emailsRepository.
					EXPECT().
					GetUserCommunications(gomock.Any(), userID).
					Return(
						[]entities.Email{
							{
								ID:      1,
								UserID:  userID,
								Email:   "someTestEmail@gmail.com",
								Content: "some test content",
								SentAt:  now,
							},
						},
						nil,
					).
					Times(1)
			},
		},
		{
			name:     "get user email communications without existing email communications",
			userID:   userID,
			expected: []entities.Email{},
			setupMocks: func(emailsRepository *mockrepositories.MockEmailsRepository, _ *loggermock.MockLogger) {
				emailsRepository.
					EXPECT().
					GetUserCommunications(gomock.Any(), userID).
					Return([]entities.Email{}, nil).
					Times(1)
			},
		},
		{
			name:   "get user email communications fail",
			userID: userID,
			setupMocks: func(emailsRepository *mockrepositories.MockEmailsRepository, _ *loggermock.MockLogger) {
				emailsRepository.
					EXPECT().
					GetUserCommunications(gomock.Any(), userID).
					Return(nil, errors.New("some error")).
					Times(1)
			},
			errorExpected: true,
		},
	}

	ctx := context.Background()
	mockController := gomock.NewController(t)
	logger := loggermock.NewMockLogger(mockController)
	emailsRepository := mockrepositories.NewMockEmailsRepository(mockController)
	emailsService := services.NewEmailsService(emailsRepository, logger)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(emailsRepository, logger)
			}

			actual, err := emailsService.GetUserCommunications(ctx, tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tc.expected, actual)
		})
	}
}
