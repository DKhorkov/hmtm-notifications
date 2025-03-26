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

func NewUpdateTicketBuilder(
	useCases interfaces.UseCases,
	traceProvider tracing.Provider,
	spanConfig tracing.SpanConfig,
	logger logging.Logger,
) *UpdateTicketBuilder {
	return &UpdateTicketBuilder{
		useCases:      useCases,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
		logger:        logger,
	}
}

type UpdateTicketBuilder struct {
	useCases      interfaces.UseCases
	traceProvider tracing.Provider
	spanConfig    tracing.SpanConfig
	logger        logging.Logger
}

func (b *UpdateTicketBuilder) MessageHandler() handlers.MessageHandler {
	return func(message *nats.Msg) {
		ctx, span := b.traceProvider.Span(
			context.Background(),
			tracing.CallerName(tracing.DefaultSkipLevel),
		)
		defer span.End()

		span.AddEvent(b.spanConfig.Events.Start.Name, b.spanConfig.Events.Start.Opts...)
		defer span.AddEvent(b.spanConfig.Events.End.Name, b.spanConfig.Events.End.Opts...)

		ctx = helpers.AddTraceIDToContext(ctx, span)

		updateTicketDTO := b.natsMessageToDTO(message)
		if updateTicketDTO == nil {
			return
		}

		if _, err := b.useCases.SendUpdateTicketEmailCommunication(
			ctx,
			updateTicketDTO.TicketID,
		); err != nil {
			logging.LogError(
				b.logger,
				fmt.Sprintf(
					"Failed to send update-ticket message for Ticket with ID=%d ",
					updateTicketDTO.TicketID,
				),
				err,
			)
		}
	}
}

func (b *UpdateTicketBuilder) natsMessageToDTO(message *nats.Msg) *dto.UpdateTicketDTO {
	var updateTicketDTO dto.UpdateTicketDTO
	if err := json.Unmarshal(message.Data, &updateTicketDTO); err != nil {
		logging.LogError(b.logger, "Failed to unmarshal update-ticket message", err)
		return nil
	}

	return &updateTicketDTO
}
