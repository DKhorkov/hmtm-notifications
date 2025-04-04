package interfaces

import "context"

type Senders struct {
	Email EmailSender
}

type EmailSender interface {
	Send(ctx context.Context, subject, body string, recipients []string) error
}
