package interfaces

import (
	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

type ContentBuilders struct {
	VerifyEmail    VerifyEmailContentBuilder
	ForgetPassword ForgetPasswordContentBuilder
	TicketUpdated  TicketUpdatedContentBuilder
	TicketDeleted  TicketDeletedContentBuilder
}

//go:generate mockgen -source=content_builders.go -destination=../../mocks/contentbuilders/verify_email_content_builder.go -package=mockcontentbuilders -exclude_interfaces=ForgetPasswordContentBuilder,TicketUpdatedContentBuilder,TicketDeletedContentBuilder
type VerifyEmailContentBuilder interface {
	Subject() string
	Body(user entities.User) string
}

//go:generate mockgen -source=content_builders.go -destination=../../mocks/contentbuilders/forget_password_content_builder.go -package=mockcontentbuilders -exclude_interfaces=VerifyEmailContentBuilder,TicketUpdatedContentBuilder,TicketDeletedContentBuilder
type ForgetPasswordContentBuilder interface {
	Subject() string
	Body(user entities.User) string
}

//go:generate mockgen -source=content_builders.go -destination=../../mocks/contentbuilders/ticket_updated_content_builder.go -package=mockcontentbuilders -exclude_interfaces=ForgetPasswordContentBuilder,VerifyEmailContentBuilder,TicketDeletedContentBuilder
type TicketUpdatedContentBuilder interface {
	Subject(ticket entities.RawTicket) string
	Body(ticket entities.RawTicket, respondOwner entities.User) string
}

//go:generate mockgen -source=content_builders.go -destination=../../mocks/contentbuilders/ticket_deleted_content_builder.go -package=mockcontentbuilders -exclude_interfaces=ForgetPasswordContentBuilder,TicketUpdatedContentBuilder,VerifyEmailContentBuilder
type TicketDeletedContentBuilder interface {
	Subject(ticketData dto.TicketDeletedDTO) string
	Body(ticketData dto.TicketDeletedDTO, ticketOwner, respondOwner entities.User) string
}
