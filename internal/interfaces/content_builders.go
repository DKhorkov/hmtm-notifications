package interfaces

import "github.com/DKhorkov/hmtm-notifications/internal/entities"

type ContentBuilders struct {
	Email EmailContentBuilder
}

type EmailContentBuilder interface {
	Subject() string
	Body(user entities.User) string
}
