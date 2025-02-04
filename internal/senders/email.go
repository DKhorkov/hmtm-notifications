package senders

import (
	"context"

	"github.com/DKhorkov/libs/tracing"
	"gopkg.in/gomail.v2"

	"github.com/DKhorkov/hmtm-notifications/internal/config"
)

func NewEmailSender(
	smtpConfig config.EmailSMTPConfig,
	traceProvider tracing.TraceProvider,
	spanConfig tracing.SpanConfig,
) *EmailSender {
	return &EmailSender{
		smtpConfig:    smtpConfig,
		traceProvider: traceProvider,
		spanConfig:    spanConfig,
	}
}

type EmailSender struct {
	smtpConfig    config.EmailSMTPConfig
	traceProvider tracing.TraceProvider
	spanConfig    tracing.SpanConfig
}

func (s *EmailSender) Send(ctx context.Context, subject string, body string, recipients []string) error {
	ctx, span := s.traceProvider.Span(ctx, tracing.CallerName(tracing.DefaultSkipLevel))
	defer span.End()

	span.AddEvent(s.spanConfig.Events.Start.Name, s.spanConfig.Events.Start.Opts...)

	message := gomail.NewMessage()
	message.SetHeader("From", s.smtpConfig.Login)
	message.SetHeader("To", recipients...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	smtpClient := gomail.NewDialer(
		s.smtpConfig.Host,
		s.smtpConfig.Port,
		s.smtpConfig.Login,
		s.smtpConfig.Password,
	)

	err := smtpClient.DialAndSend(message)

	span.AddEvent(s.spanConfig.Events.End.Name, s.spanConfig.Events.End.Opts...)
	return err
}
