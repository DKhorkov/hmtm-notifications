package builders

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"

	"github.com/DKhorkov/libs/logging"
	"github.com/DKhorkov/libs/tracing"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
	"github.com/DKhorkov/hmtm-notifications/internal/workers/handlers"
	"github.com/DKhorkov/hmtm-notifications/internal/workers/handlers/helpers"
)

func NewDeleteTicketBuilder(
	useCases interfaces.UseCases,
	traceProvider tracing.Provider,
	spanConfig tracing.SpanConfig,
	logger logging.Logger,
) *DeleteTicketBuilder {
	return &DeleteTicketBuilder{
		useCases:      useCases,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
		logger:        logger,
	}
}

type DeleteTicketBuilder struct {
	useCases      interfaces.UseCases
	traceProvider tracing.Provider
	spanConfig    tracing.SpanConfig
	logger        logging.Logger
}

func (b *DeleteTicketBuilder) MessageHandler() handlers.MessageHandler {
	return func(message *nats.Msg) {
		ctx, span := b.traceProvider.Span(context.Background(), tracing.CallerName(tracing.DefaultSkipLevel))
		defer span.End()

		span.AddEvent(b.spanConfig.Events.Start.Name, b.spanConfig.Events.Start.Opts...)
		defer span.AddEvent(b.spanConfig.Events.End.Name, b.spanConfig.Events.End.Opts...)

		ctx = helpers.AddTraceIDToContext(ctx, span)

		deleteTicketDTO := b.natsMessageToDTO(message)
		if deleteTicketDTO == nil {
			return
		}

		if _, err := b.useCases.SendDeleteTicketEmailCommunication(
			ctx,
			*deleteTicketDTO,
		); err != nil {
			logging.LogError(
				b.logger,
				"Failed to send delete-ticket message",
				err,
			)
		}
	}
}

func (b *DeleteTicketBuilder) natsMessageToDTO(message *nats.Msg) *dto.DeleteTicketDTO {
	var deleteTicketDTO dto.DeleteTicketDTO
	if err := json.Unmarshal(message.Data, &deleteTicketDTO); err != nil {
		logging.LogError(b.logger, "Failed to unmarshal delete-ticket message", err)
		return nil
	}

	return &deleteTicketDTO
}
