package repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	mockclients "github.com/DKhorkov/hmtm-notifications/mocks/clients"
	"github.com/DKhorkov/hmtm-sso/api/protobuf/generated/go/sso"
)

func TestSsoRepository_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	ssoClient := mockclients.NewMockSsoClient(ctrl)
	repo := NewSsoRepository(ssoClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name          string
		id            uint64
		setupMocks    func(ssoClient *mockclients.MockSsoClient)
		expectedUser  *entities.User
		errorExpected bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(ssoClient *mockclients.MockSsoClient) {
				ssoClient.
					EXPECT().
					GetUser(
						gomock.Any(),
						&sso.GetUserIn{ID: 1},
					).
					Return(&sso.GetUserOut{
						ID:          1,
						DisplayName: "Test User",
						Email:       "test@example.com",
						CreatedAt:   timestamppb.New(now),
						UpdatedAt:   timestamppb.New(now),
					}, nil).
					Times(1)
			},
			expectedUser: &entities.User{
				ID:          1,
				DisplayName: "Test User",
				Email:       "test@example.com",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(ssoClient *mockclients.MockSsoClient) {
				ssoClient.
					EXPECT().
					GetUser(
						gomock.Any(),
						&sso.GetUserIn{ID: 1},
					).
					Return(nil, errors.New("get user failed")).
					Times(1)
			},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ssoClient)
			}

			user, err := repo.GetUserByID(context.Background(), tc.id)
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

func TestSsoRepository_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	ssoClient := mockclients.NewMockSsoClient(ctrl)
	repo := NewSsoRepository(ssoClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name          string
		email         string
		setupMocks    func(ssoClient *mockclients.MockSsoClient)
		expectedUser  *entities.User
		errorExpected bool
	}{
		{
			name:  "success",
			email: "test@example.com",
			setupMocks: func(ssoClient *mockclients.MockSsoClient) {
				ssoClient.
					EXPECT().
					GetUserByEmail(
						gomock.Any(),
						&sso.GetUserByEmailIn{Email: "test@example.com"},
					).
					Return(&sso.GetUserOut{
						ID:          1,
						DisplayName: "Test User",
						Email:       "test@example.com",
						CreatedAt:   timestamppb.New(now),
						UpdatedAt:   timestamppb.New(now),
					}, nil).
					Times(1)
			},
			expectedUser: &entities.User{
				ID:          1,
				DisplayName: "Test User",
				Email:       "test@example.com",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			errorExpected: false,
		},
		{
			name:  "error",
			email: "test@example.com",
			setupMocks: func(ssoClient *mockclients.MockSsoClient) {
				ssoClient.
					EXPECT().
					GetUserByEmail(
						gomock.Any(),
						&sso.GetUserByEmailIn{Email: "test@example.com"},
					).
					Return(nil, errors.New("get user failed")).
					Times(1)
			},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ssoClient)
			}

			user, err := repo.GetUserByEmail(context.Background(), tc.email)
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
