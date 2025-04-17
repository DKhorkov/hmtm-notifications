package helpers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"

	"github.com/DKhorkov/libs/tracing"
	mocktracing "github.com/DKhorkov/libs/tracing/mocks"
)

func TestAddTraceIDToContext(t *testing.T) {
	testCases := []struct {
		name          string
		initialCtx    context.Context
		traceID       string
		expectedMDKey string
		expectedMDVal string
	}{
		{
			name:          "empty context",
			initialCtx:    context.Background(),
			traceID:       "1234567890abcdef1234567890abcdef",
			expectedMDKey: tracing.Key,
			expectedMDVal: "1234567890abcdef1234567890abcdef",
		},
		{
			name:          "context with existing metadata",
			initialCtx:    metadata.NewOutgoingContext(context.Background(), metadata.Pairs("existing-key", "value")),
			traceID:       "abcdef1234567890abcdef1234567890",
			expectedMDKey: tracing.Key,
			expectedMDVal: "abcdef1234567890abcdef1234567890",
		},
	}

	// Создаем мок для Span
	span := mocktracing.NewMockSpan()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Вызываем тестируемую функцию
			resultCtx := AddTraceIDToContext(tc.initialCtx, span)

			// Извлекаем metadata из контекста
			md, ok := metadata.FromOutgoingContext(resultCtx)
			require.True(t, ok, "metadata should be present in the context")

			// Проверяем, что TraceID добавлен в metadata
			traceIDValues, exists := md[tc.expectedMDKey]
			require.True(t, exists, "TraceID key should exist in metadata")
			require.Len(t, traceIDValues, 1, "TraceID should have exactly one value")
		})
	}
}
