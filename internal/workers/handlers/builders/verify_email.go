package builders

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/nats-io/nats.go"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
	"github.com/DKhorkov/hmtm-notifications/internal/workers/handlers"
	"github.com/DKhorkov/hmtm-notifications/internal/workers/handlers/helpers"
	"github.com/DKhorkov/libs/logging"
	"github.com/DKhorkov/libs/tracing"
)

func NewVerifyEmailBuilder(
	useCases interfaces.UseCases,
	traceProvider tracing.TraceProvider,
	spanConfig tracing.SpanConfig,
	logger *slog.Logger,
) *VerifyEmailBuilder {
	return &VerifyEmailBuilder{
		useCases:      useCases,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
		logger:        logger,
	}
}

type VerifyEmailBuilder struct {
	useCases      interfaces.UseCases
	traceProvider tracing.TraceProvider
	spanConfig    tracing.SpanConfig
	logger        *slog.Logger
}

func (b *VerifyEmailBuilder) MessageHandler() handlers.MessageHandler {
	return func(message *nats.Msg) {
		ctx, span := b.traceProvider.Span(context.Background(), tracing.CallerName(tracing.DefaultSkipLevel))
		defer span.End()

		span.AddEvent(b.spanConfig.Events.Start.Name, b.spanConfig.Events.Start.Opts...)

		ctx = helpers.AddTraceIDToContext(ctx, span)

		verifyEmailDTO := b.natsMessageToDTO(message)
		if verifyEmailDTO == nil {
			return
		}

		if _, err := b.useCases.SendVerifyEmailCommunication(
			ctx,
			verifyEmailDTO.UserID,
		); err != nil {
			logging.LogError(
				b.logger,
				fmt.Sprintf("Failed to send verify-email message to User with ID=%d ", verifyEmailDTO.UserID),
				err,
			)
		}

		span.AddEvent(b.spanConfig.Events.End.Name, b.spanConfig.Events.End.Opts...)
	}
}

func (b *VerifyEmailBuilder) natsMessageToDTO(message *nats.Msg) *dto.VerifyEmailDTO {
	var verifyEmailDTO dto.VerifyEmailDTO
	if err := json.Unmarshal(message.Data, &verifyEmailDTO); err != nil {
		logging.LogError(b.logger, "Failed to unmarshal verify-email message", err)
		return nil
	}

	return &verifyEmailDTO
}
