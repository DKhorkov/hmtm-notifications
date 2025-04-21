package main

import (
	"encoding/json"
	"time"

	"github.com/DKhorkov/libs/pointers"
	"github.com/nats-io/nats.go"

	customnats "github.com/DKhorkov/libs/nats"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/config"
)

func main() {
	settings := config.New()

	natsPublisher, err := customnats.NewPublisher(
		settings.NATS.ClientURL,
		nats.Name("hmtm-notifications-test"),
	)
	if err != nil {
		panic(err)
	}

	ticketDeletedDTO := dto.TicketDeletedDTO{
		TicketOwnerID:       1,
		Name:                "test ticket",
		Description:         "test description",
		Price:               pointers.New[float32](112),
		Quantity:            1,
		RespondedMastersIDs: []uint64{1},
	}

	content, err := json.Marshal(ticketDeletedDTO)
	if err != nil {
		panic(err)
	}

	err = natsPublisher.Publish(settings.NATS.Subjects.TicketDeleted, content)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 2)
}
