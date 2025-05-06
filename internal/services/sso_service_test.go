package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	mockrepositories "github.com/DKhorkov/hmtm-notifications/mocks/repositories"
	mocklogging "github.com/DKhorkov/libs/logging/mocks"
)

func TestSsoService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	ssoRepository := mockrepositories.NewMockSsoRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewSsoService(ssoRepository, logger)

	testCases := []struct {
		name          string
		id            uint64
		setupMocks    func(ssoRepository *mockrepositories.MockSsoRepository, logger *mocklogging.MockLogger)
		expectedUser  *entities.User
		errorExpected bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(ssoRepository *mockrepositories.MockSsoRepository, logger *mocklogging.MockLogger) {
				ssoRepository.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&entities.User{ID: 1, Email: "user@example.com"}, nil).
					Times(1)
			},
			expectedUser:  &entities.User{ID: 1, Email: "user@example.com"},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(ssoRepository *mockrepositories.MockSsoRepository, logger *mocklogging.MockLogger) {
				ssoRepository.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedUser:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ssoRepository, logger)
			}

			user, err := service.GetUserByID(context.Background(), tc.id)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, user)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedUser, user)
			}
		})
	}
}

func TestSsoService_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	ssoRepository := mockrepositories.NewMockSsoRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewSsoService(ssoRepository, logger)

	testCases := []struct {
		name          string
		email         string
		setupMocks    func(ssoRepository *mockrepositories.MockSsoRepository, logger *mocklogging.MockLogger)
		expectedUser  *entities.User
		errorExpected bool
	}{
		{
			name:  "success",
			email: "user@example.com",
			setupMocks: func(ssoRepository *mockrepositories.MockSsoRepository, logger *mocklogging.MockLogger) {
				ssoRepository.
					EXPECT().
					GetUserByEmail(gomock.Any(), "user@example.com").
					Return(&entities.User{ID: 1, Email: "user@example.com"}, nil).
					Times(1)
			},
			expectedUser:  &entities.User{ID: 1, Email: "user@example.com"},
			errorExpected: false,
		},
		{
			name:  "error",
			email: "user@example.com",
			setupMocks: func(ssoRepository *mockrepositories.MockSsoRepository, logger *mocklogging.MockLogger) {
				ssoRepository.
					EXPECT().
					GetUserByEmail(gomock.Any(), "user@example.com").
					Return(nil, errors.New("not found")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedUser:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ssoRepository, logger)
			}

			user, err := service.GetUserByEmail(context.Background(), tc.email)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, user)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedUser, user)
			}
		})
	}
}
