package handlers

import "github.com/nats-io/nats.go"

type MessageHandler func(message *nats.Msg)
