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

type TicketUpdatedBuilder struct {
	useCases      interfaces.UseCases
	traceProvider tracing.Provider
	spanConfig    tracing.SpanConfig
	logger        logging.Logger
}

func NewTicketUpdatedBuilder(
	useCases interfaces.UseCases,
	traceProvider tracing.Provider,
	spanConfig tracing.SpanConfig,
	logger logging.Logger,
) *TicketUpdatedBuilder {
	return &TicketUpdatedBuilder{
		useCases:      useCases,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
		logger:        logger,
	}
}

func (b *TicketUpdatedBuilder) MessageHandler() handlers.MessageHandler {
	return func(message *nats.Msg) {
		ctx, span := b.traceProvider.Span(
			context.Background(),
			tracing.CallerName(tracing.DefaultSkipLevel),
		)
		defer span.End()

		span.AddEvent(b.spanConfig.Events.Start.Name, b.spanConfig.Events.Start.Opts...)
		defer span.AddEvent(b.spanConfig.Events.End.Name, b.spanConfig.Events.End.Opts...)

		ctx = helpers.AddTraceIDToContext(ctx, span)

		ticketUpdatedDTO := b.natsMessageToDTO(message)
		if ticketUpdatedDTO == nil {
			return
		}

		if _, err := b.useCases.SendTicketUpdatedEmailCommunication(
			ctx,
			ticketUpdatedDTO.TicketID,
		); err != nil {
			logging.LogError(
				b.logger,
				fmt.Sprintf(
					"Failed to send update-ticket message for Ticket with ID=%d ",
					ticketUpdatedDTO.TicketID,
				),
				err,
			)
		}
	}
}

func (b *TicketUpdatedBuilder) natsMessageToDTO(message *nats.Msg) *dto.TicketUpdatedDTO {
	var ticketUpdatedDTO dto.TicketUpdatedDTO
	if err := json.Unmarshal(message.Data, &ticketUpdatedDTO); err != nil {
		logging.LogError(b.logger, "Failed to unmarshal update-ticket message", err)

		return nil
	}

	return &ticketUpdatedDTO
}
