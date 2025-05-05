package emails

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	customgrpc "github.com/DKhorkov/libs/grpc"
	mocklogging "github.com/DKhorkov/libs/logging/mocks"

	"github.com/DKhorkov/hmtm-notifications/api/protobuf/generated/go/notifications"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	mockusecases "github.com/DKhorkov/hmtm-notifications/mocks/usecases"
)

func TestServerAPI_GetUserEmailCommunications(t *testing.T) {
	ctrl := gomock.NewController(t)
	useCases := mockusecases.NewMockUseCases(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	api := &ServerAPI{
		useCases: useCases,
		logger:   logger,
	}

	testCases := []struct {
		name          string
		in            *notifications.GetUserEmailCommunicationsIn
		setupMocks    func(useCases *mockusecases.MockUseCases, logger *mocklogging.MockLogger)
		expectedOut   *notifications.GetUserEmailCommunicationsOut
		expectedErr   error
		errorExpected bool
	}{
		{
			name: "success with emails",
			in:   &notifications.GetUserEmailCommunicationsIn{UserID: 1},
			setupMocks: func(useCases *mockusecases.MockUseCases, logger *mocklogging.MockLogger) {
				emailCommunications := []entities.Email{
					{
						ID:      1,
						UserID:  1,
						Email:   "test1@example.com",
						Content: "Hello, this is email 1",
						SentAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:      2,
						UserID:  1,
						Email:   "test2@example.com",
						Content: "Hello, this is email 2",
						SentAt:  time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
					},
				}
				useCases.
					EXPECT().
					GetUserEmailCommunications(gomock.Any(), uint64(1)).
					Return(emailCommunications, nil).
					Times(1)
			},
			expectedOut: &notifications.GetUserEmailCommunicationsOut{
				Emails: []*notifications.Email{
					{
						ID:      1,
						UserID:  1,
						Email:   "test1@example.com",
						Content: "Hello, this is email 1",
						SentAt:  timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
					},
					{
						ID:      2,
						UserID:  1,
						Email:   "test2@example.com",
						Content: "Hello, this is email 2",
						SentAt:  timestamppb.New(time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
			expectedErr:   nil,
			errorExpected: false,
		},
		{
			name: "success with no emails",
			in:   &notifications.GetUserEmailCommunicationsIn{UserID: 2},
			setupMocks: func(useCases *mockusecases.MockUseCases, logger *mocklogging.MockLogger) {
				useCases.
					EXPECT().
					GetUserEmailCommunications(gomock.Any(), uint64(2)).
					Return([]entities.Email{}, nil).
					Times(1)
			},
			expectedOut: &notifications.GetUserEmailCommunicationsOut{
				Emails: []*notifications.Email{},
			},
			expectedErr:   nil,
			errorExpected: false,
		},
		{
			name: "internal error",
			in:   &notifications.GetUserEmailCommunicationsIn{UserID: 3},
			setupMocks: func(useCases *mockusecases.MockUseCases, logger *mocklogging.MockLogger) {
				useCases.
					EXPECT().
					GetUserEmailCommunications(gomock.Any(), uint64(3)).
					Return(nil, errors.New("internal error")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedErr:   &customgrpc.BaseError{Status: codes.Internal, Message: "internal error"},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(useCases, logger)
			}

			resp, err := api.GetUserEmailCommunications(context.Background(), tc.in)
			if tc.errorExpected {
				require.Error(t, err)
				require.Equal(t, tc.expectedErr, err)
				require.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedOut, resp)
			}
		})
	}
}
