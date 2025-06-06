package builders

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DKhorkov/libs/logging"
	"github.com/DKhorkov/libs/tracing"
	"github.com/nats-io/nats.go"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
	"github.com/DKhorkov/hmtm-notifications/internal/workers/handlers"
	"github.com/DKhorkov/hmtm-notifications/internal/workers/handlers/helpers"
)

type VerifyEmailBuilder struct {
	useCases      interfaces.UseCases
	traceProvider tracing.Provider
	spanConfig    tracing.SpanConfig
	logger        logging.Logger
}

func NewVerifyEmailBuilder(
	useCases interfaces.UseCases,
	traceProvider tracing.Provider,
	spanConfig tracing.SpanConfig,
	logger logging.Logger,
) *VerifyEmailBuilder {
	return &VerifyEmailBuilder{
		useCases:      useCases,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
		logger:        logger,
	}
}

func (b *VerifyEmailBuilder) MessageHandler() handlers.MessageHandler {
	return func(message *nats.Msg) {
		ctx, span := b.traceProvider.Span(
			context.Background(),
			tracing.CallerName(tracing.DefaultSkipLevel),
		)
		defer span.End()

		span.AddEvent(b.spanConfig.Events.Start.Name, b.spanConfig.Events.Start.Opts...)
		defer span.AddEvent(b.spanConfig.Events.End.Name, b.spanConfig.Events.End.Opts...)

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
				fmt.Sprintf(
					"Failed to send verify-email message to User with ID=%d ",
					verifyEmailDTO.UserID,
				),
				err,
			)
		}
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
