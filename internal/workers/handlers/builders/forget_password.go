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

func NewForgetPasswordBuilder(
	useCases interfaces.UseCases,
	traceProvider tracing.Provider,
	spanConfig tracing.SpanConfig,
	logger logging.Logger,
) *ForgetPasswordBuilder {
	return &ForgetPasswordBuilder{
		useCases:      useCases,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
		logger:        logger,
	}
}

type ForgetPasswordBuilder struct {
	useCases      interfaces.UseCases
	traceProvider tracing.Provider
	spanConfig    tracing.SpanConfig
	logger        logging.Logger
}

func (b *ForgetPasswordBuilder) MessageHandler() handlers.MessageHandler {
	return func(message *nats.Msg) {
		ctx, span := b.traceProvider.Span(
			context.Background(),
			tracing.CallerName(tracing.DefaultSkipLevel),
		)
		defer span.End()

		span.AddEvent(b.spanConfig.Events.Start.Name, b.spanConfig.Events.Start.Opts...)
		defer span.AddEvent(b.spanConfig.Events.End.Name, b.spanConfig.Events.End.Opts...)

		ctx = helpers.AddTraceIDToContext(ctx, span)

		forgetPasswordDTO := b.natsMessageToDTO(message)
		if forgetPasswordDTO == nil {
			return
		}

		if _, err := b.useCases.SendForgetPasswordEmailCommunication(
			ctx,
			forgetPasswordDTO.UserID,
			forgetPasswordDTO.NewPassword,
		); err != nil {
			logging.LogError(
				b.logger,
				fmt.Sprintf(
					"Failed to send forget-password message to User with ID=%d ",
					forgetPasswordDTO.UserID,
				),
				err,
			)
		}
	}
}

func (b *ForgetPasswordBuilder) natsMessageToDTO(message *nats.Msg) *dto.ForgetPasswordDTO {
	var forgetPasswordDTO dto.ForgetPasswordDTO
	if err := json.Unmarshal(message.Data, &forgetPasswordDTO); err != nil {
		logging.LogError(b.logger, "Failed to unmarshal forget-password message", err)
		return nil
	}

	return &forgetPasswordDTO
}
