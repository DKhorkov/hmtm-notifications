package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	mocklogging "github.com/DKhorkov/libs/logging/mocks"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/services"
	mockrepositories "github.com/DKhorkov/hmtm-notifications/mocks/repositories"
)

var (
	userID uint64 = 1
	now           = time.Now()
)

func TestEmailsService_GetUserEmailCommunications(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name          string
		userID        uint64
		setupMocks    func(emailsRepository *mockrepositories.MockEmailsRepository, logger *mocklogging.MockLogger)
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
			setupMocks: func(emailsRepository *mockrepositories.MockEmailsRepository, _ *mocklogging.MockLogger) {
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
			setupMocks: func(emailsRepository *mockrepositories.MockEmailsRepository, _ *mocklogging.MockLogger) {
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
			setupMocks: func(emailsRepository *mockrepositories.MockEmailsRepository, _ *mocklogging.MockLogger) {
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
	ctrl := gomock.NewController(t)
	logger := mocklogging.NewMockLogger(ctrl)
	emailsRepository := mockrepositories.NewMockEmailsRepository(ctrl)
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

func TestEmailsService_SaveCommunication(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := mocklogging.NewMockLogger(ctrl)
	emailsRepository := mockrepositories.NewMockEmailsRepository(ctrl)
	emailsService := services.NewEmailsService(emailsRepository, logger)

	testCases := []struct {
		name          string
		email         entities.Email
		setupMocks    func(emailsRepository *mockrepositories.MockEmailsRepository)
		expectedID    uint64
		errorExpected bool
	}{
		{
			name: "success",
			email: entities.Email{
				// Здесь предполагается структура Email, заполняем минимально для примера
				ID:      0, // ID обычно 0 для новой записи
				UserID:  userID,
				Content: "Test Subject",
				SentAt:  now,
			},
			setupMocks: func(emailsRepository *mockrepositories.MockEmailsRepository) {
				emailsRepository.
					EXPECT().
					SaveCommunication(gomock.Any(), entities.Email{
						ID:      0,
						UserID:  userID,
						Content: "Test Subject",
						SentAt:  now,
					}).
					Return(uint64(1), nil).
					Times(1)
			},
			expectedID:    1,
			errorExpected: false,
		},
		{
			name: "error",
			email: entities.Email{
				ID:      0,
				UserID:  userID,
				Content: "Test Subject",
				SentAt:  now,
			},
			setupMocks: func(emailsRepository *mockrepositories.MockEmailsRepository) {
				emailsRepository.
					EXPECT().
					SaveCommunication(gomock.Any(), entities.Email{
						ID:      0,
						UserID:  userID,
						Content: "Test Subject",
						SentAt:  now,
					}).
					Return(uint64(0), errors.New("save failed")).
					Times(1)
			},
			expectedID:    0,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(emailsRepository)
			}

			emailID, err := emailsService.SaveCommunication(context.Background(), tc.email)
			if tc.errorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expectedID, emailID)
		})
	}
}
