package interfaces

import "context"

type Senders struct {
	Email EmailSender
}

//go:generate mockgen -source=senders.go -destination=../../mocks/senders/email_sender.go -package=mocksenders -exclude_interfaces=
type EmailSender interface {
	Send(ctx context.Context, subject, body string, recipients []string) error
}
