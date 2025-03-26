package helpers

import (
	"context"

	"github.com/DKhorkov/libs/tracing"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

func AddTraceIDToContext(ctx context.Context, span trace.Span) context.Context {
	traceID := span.SpanContext().TraceID().String()
	ctx = metadata.AppendToOutgoingContext(
		ctx,
		tracing.Key,
		traceID,
	) // setting for cross-service usage
	return ctx
}
