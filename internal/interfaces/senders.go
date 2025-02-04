package interfaces

import "context"

type Senders struct {
	Email EmailSender
}

type EmailSender interface {
	Send(ctx context.Context, subject string, body string, recipients []string) error
}
