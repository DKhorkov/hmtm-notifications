package builders

import (
	"context"
	"errors"
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	mocklogging "github.com/DKhorkov/libs/logging/mocks"
	"github.com/DKhorkov/libs/pointers"
	"github.com/DKhorkov/libs/tracing"
	mocktracing "github.com/DKhorkov/libs/tracing/mocks"

	"github.com/DKhorkov/hmtm-notifications/dto"
	mockusecases "github.com/DKhorkov/hmtm-notifications/mocks/usecases"
)

func TestTicketDeletedBuilder_MessageHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	useCases := mockusecases.NewMockUseCases(ctrl)
	traceProvider := mocktracing.NewMockProvider(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	spanConfig := tracing.SpanConfig{}
	builder := NewTicketDeletedBuilder(
		useCases,
		traceProvider,
		spanConfig,
		logger,
	)

	testCases := []struct {
		name       string
		message    *nats.Msg
		setupMocks func(useCases *mockusecases.MockUseCases, traceProvider *mocktracing.MockProvider, logger *mocklogging.MockLogger)
	}{
		{
			name: "successful processing",
			message: &nats.Msg{
				Data: []byte(`{"name":"Teddy Bear","description":"Soft toy","quantity":5,"price":150.75}`),
			},
			setupMocks: func(useCases *mockusecases.MockUseCases, traceProvider *mocktracing.MockProvider, logger *mocklogging.MockLogger) {
				traceProvider.
					EXPECT().
					Span(gomock.Any(), gomock.Any()).
					Return(context.Background(), mocktracing.NewMockSpan()).
					Times(1)

				ticketDeletedDTO := dto.TicketDeletedDTO{
					Name:        "Teddy Bear",
					Description: "Soft toy",
					Quantity:    5,
					Price:       pointers.New[float32](150.75),
				}

				useCases.
					EXPECT().
					SendTicketDeletedEmailCommunication(gomock.Any(), ticketDeletedDTO).
					Return(nil, nil).
					Times(1)
			},
		},
		{
			name: "invalid message data",
			message: &nats.Msg{
				Data: []byte(`{invalid json}`),
			},
			setupMocks: func(useCases *mockusecases.MockUseCases, traceProvider *mocktracing.MockProvider, logger *mocklogging.MockLogger) {
				traceProvider.
					EXPECT().
					Span(gomock.Any(), gomock.Any()).
					Return(context.Background(), mocktracing.NewMockSpan()).
					Times(1)

				logger.
					EXPECT().
					Error(gomock.Any(), gomock.Any()).
					Times(1)
			},
		},
		{
			name: "use case error",
			message: &nats.Msg{
				Data: []byte(`{"name":"Wooden Car","description":"Toy car","quantity":1}`),
			},
			setupMocks: func(useCases *mockusecases.MockUseCases, traceProvider *mocktracing.MockProvider, logger *mocklogging.MockLogger) {
				traceProvider.
					EXPECT().
					Span(gomock.Any(), gomock.Any()).
					Return(context.Background(), mocktracing.NewMockSpan()).
					Times(1)

				ticketDeletedDTO := dto.TicketDeletedDTO{
					Name:        "Wooden Car",
					Description: "Toy car",
					Quantity:    1,
					Price:       nil,
				}

				useCases.
					EXPECT().
					SendTicketDeletedEmailCommunication(gomock.Any(), ticketDeletedDTO).
					Return(nil, errors.New("test")).
					Times(1)

				logger.
					EXPECT().
					Error(gomock.Any(), gomock.Any()).
					Times(1)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks(useCases, traceProvider, logger)
			handler := builder.MessageHandler()
			handler(tc.message)
		})
	}
}

func TestTicketDeletedBuilder_natsMessageToDTO(t *testing.T) {
	ctrl := gomock.NewController(t)
	useCases := mockusecases.NewMockUseCases(ctrl)
	traceProvider := mocktracing.NewMockProvider(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	spanConfig := tracing.SpanConfig{}
	builder := NewTicketDeletedBuilder(
		useCases,
		traceProvider,
		spanConfig,
		logger,
	)

	testCases := []struct {
		name        string
		message     *nats.Msg
		expectedDTO *dto.TicketDeletedDTO
		setupMocks  func(logger *mocklogging.MockLogger)
	}{
		{
			name: "valid message",
			message: &nats.Msg{
				Data: []byte(`{"name":"Teddy Bear","description":"Soft toy","quantity":5,"price":150.75}`),
			},
			expectedDTO: &dto.TicketDeletedDTO{
				Name:        "Teddy Bear",
				Description: "Soft toy",
				Quantity:    5,
				Price:       pointers.New[float32](150.75),
			},
			setupMocks: func(logger *mocklogging.MockLogger) {},
		},
		{
			name: "invalid message",
			message: &nats.Msg{
				Data: []byte(`{invalid json}`),
			},
			expectedDTO: nil,
			setupMocks: func(logger *mocklogging.MockLogger) {
				logger.
					EXPECT().
					Error(gomock.Any(), gomock.Any()).
					Times(1)
			},
		},
		{
			name: "empty message",
			message: &nats.Msg{
				Data: []byte(``),
			},
			expectedDTO: nil,
			setupMocks: func(logger *mocklogging.MockLogger) {
				logger.
					EXPECT().
					Error(gomock.Any(), gomock.Any()).
					Times(1)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks(logger)
			result := builder.natsMessageToDTO(tc.message)
			require.Equal(t, tc.expectedDTO, result)
		})
	}
}
