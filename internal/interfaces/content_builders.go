package interfaces

import (
	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

type ContentBuilders struct {
	VerifyEmail    VerifyEmailContentBuilder
	ForgetPassword ForgetPasswordContentBuilder
	UpdateTicket   UpdateTicketContentBuilder
	DeleteTicket   DeleteTicketContentBuilder
}

type VerifyEmailContentBuilder interface {
	Subject() string
	Body(user entities.User) string
}

type ForgetPasswordContentBuilder interface {
	Subject() string
	Body(user entities.User, newPassword string) string
}

type UpdateTicketContentBuilder interface {
	Subject(ticket entities.RawTicket) string
	Body(ticket entities.RawTicket, respondOwner entities.User) string
}

type DeleteTicketContentBuilder interface {
	Subject(ticketData dto.DeleteTicketDTO) string
	Body(ticketData dto.DeleteTicketDTO, ticketOwner, respondOwner entities.User) string
}
