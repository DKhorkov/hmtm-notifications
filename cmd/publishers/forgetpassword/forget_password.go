package main

import (
	"encoding/json"
	"time"

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

	forgetPasswordDTO := dto.ForgetPasswordDTO{
		UserID: 31,
	}

	content, err := json.Marshal(forgetPasswordDTO)
	if err != nil {
		panic(err)
	}

	err = natsPublisher.Publish(settings.NATS.Subjects.ForgetPassword, content)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 2)
}
