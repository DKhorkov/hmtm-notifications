package builders

import (
	"context"
	"encoding/json"

	"github.com/DKhorkov/libs/logging"
	"github.com/DKhorkov/libs/tracing"
	"github.com/nats-io/nats.go"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
	"github.com/DKhorkov/hmtm-notifications/internal/workers/handlers"
	"github.com/DKhorkov/hmtm-notifications/internal/workers/handlers/helpers"
)

type TicketDeletedBuilder struct {
	useCases      interfaces.UseCases
	traceProvider tracing.Provider
	spanConfig    tracing.SpanConfig
	logger        logging.Logger
}

func NewTicketDeletedBuilder(
	useCases interfaces.UseCases,
	traceProvider tracing.Provider,
	spanConfig tracing.SpanConfig,
	logger logging.Logger,
) *TicketDeletedBuilder {
	return &TicketDeletedBuilder{
		useCases:      useCases,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
		logger:        logger,
	}
}

func (b *TicketDeletedBuilder) MessageHandler() handlers.MessageHandler {
	return func(message *nats.Msg) {
		ctx, span := b.traceProvider.Span(
			context.Background(),
			tracing.CallerName(tracing.DefaultSkipLevel),
		)
		defer span.End()

		span.AddEvent(b.spanConfig.Events.Start.Name, b.spanConfig.Events.Start.Opts...)
		defer span.AddEvent(b.spanConfig.Events.End.Name, b.spanConfig.Events.End.Opts...)

		ctx = helpers.AddTraceIDToContext(ctx, span)

		ticketDeletedDTO := b.natsMessageToDTO(message)
		if ticketDeletedDTO == nil {
			return
		}

		if _, err := b.useCases.SendTicketDeletedEmailCommunication(
			ctx,
			*ticketDeletedDTO,
		); err != nil {
			logging.LogError(
				b.logger,
				"Failed to send delete-ticket message",
				err,
			)
		}
	}
}

func (b *TicketDeletedBuilder) natsMessageToDTO(message *nats.Msg) *dto.TicketDeletedDTO {
	var ticketDeletedDTO dto.TicketDeletedDTO
	if err := json.Unmarshal(message.Data, &ticketDeletedDTO); err != nil {
		logging.LogError(b.logger, "Failed to unmarshal delete-ticket message", err)

		return nil
	}

	return &ticketDeletedDTO
}
