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

//go:generate mockgen -source=content_builders.go -destination=../../mocks/contentbuilders/verify_email_content_builder.go -package=mockcontentbuilders -exclude_interfaces=ForgetPasswordContentBuilder,UpdateTicketContentBuilder,DeleteTicketContentBuilder
type VerifyEmailContentBuilder interface {
	Subject() string
	Body(user entities.User) string
}

//go:generate mockgen -source=content_builders.go -destination=../../mocks/contentbuilders/forget_password_content_builder.go -package=mockcontentbuilders -exclude_interfaces=VerifyEmailContentBuilder,UpdateTicketContentBuilder,DeleteTicketContentBuilder
type ForgetPasswordContentBuilder interface {
	Subject() string
	Body(user entities.User) string
}

//go:generate mockgen -source=content_builders.go -destination=../../mocks/contentbuilders/update_ticket_content_builder.go -package=mockcontentbuilders -exclude_interfaces=ForgetPasswordContentBuilder,VerifyEmailContentBuilder,DeleteTicketContentBuilder
type UpdateTicketContentBuilder interface {
	Subject(ticket entities.RawTicket) string
	Body(ticket entities.RawTicket, respondOwner entities.User) string
}

//go:generate mockgen -source=content_builders.go -destination=../../mocks/contentbuilders/delete_ticket_content_builder.go -package=mockcontentbuilders -exclude_interfaces=ForgetPasswordContentBuilder,UpdateTicketContentBuilder,VerifyEmailContentBuilder
type DeleteTicketContentBuilder interface {
	Subject(ticketData dto.DeleteTicketDTO) string
	Body(ticketData dto.DeleteTicketDTO, ticketOwner, respondOwner entities.User) string
}
