package senders

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/DKhorkov/libs/tracing"
	mocktracing "github.com/DKhorkov/libs/tracing/mocks"

	"github.com/DKhorkov/hmtm-notifications/internal/config"
)

func TestEmailSender_Send(t *testing.T) {
	ctrl := gomock.NewController(t)
	traceProvider := mocktracing.NewMockProvider(ctrl)

	// Настройка SMTP конфигурации
	smtpConfig := config.SMTPConfig{
		Host: "smtp.freesmtpservers.com",
		Port: 25,
	}

	// Настройка SpanConfig
	spanConfig := tracing.SpanConfig{
		Events: tracing.SpanEventsConfig{
			Start: tracing.SpanEventConfig{Name: "start_sending_email"},
			End:   tracing.SpanEventConfig{Name: "end_sending_email"},
		},
	}

	testCases := []struct {
		name          string
		subject       string
		body          string
		recipients    []string
		setupMocks    func(traceProvider *mocktracing.MockProvider)
		errorExpected bool
	}{
		{
			name:       "dialer error",
			subject:    "Test Subject",
			body:       "<h1>Test Body</h1>",
			recipients: []string{"recipient1@example.com"},
			setupMocks: func(traceProvider *mocktracing.MockProvider) {
				traceProvider.
					EXPECT().
					Span(gomock.Any(), gomock.Any()).
					Return(context.Background(), mocktracing.NewMockSpan()).
					Times(1)
			},
			errorExpected: true,
		},
	}

	sender := NewEmailSender(
		smtpConfig,
		traceProvider,
		spanConfig,
	)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(traceProvider)
			}

			err := sender.Send(context.Background(), tc.subject, tc.body, tc.recipients)
			if tc.errorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
