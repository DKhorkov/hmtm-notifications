package helpers

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"

	"github.com/DKhorkov/libs/tracing"
)

func AddTraceIDToContext(ctx context.Context, span trace.Span) context.Context {
	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(ctx, tracing.Key, traceID) // setting for cross-service usage
	return ctx
}
