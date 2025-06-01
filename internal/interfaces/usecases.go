package interfaces

import (
	"context"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

//go:generate mockgen -source=usecases.go -destination=../../mocks/usecases/usecases.go -package=mockusecases
type UseCases interface {
	GetUserEmailCommunications(
		ctx context.Context,
		userID uint64,
		pagination *entities.Pagination,
	) ([]entities.Email, error)
	CountUserEmailCommunications(ctx context.Context, userID uint64) (uint64, error)
	SendVerifyEmailCommunication(ctx context.Context, userID uint64) (emailID uint64, err error)
	SendForgetPasswordEmailCommunication(ctx context.Context, userID uint64) (emailID uint64, err error)
	SendTicketUpdatedEmailCommunication(ctx context.Context, ticketID uint64) (emailIDs []uint64, err error)
	SendTicketDeletedEmailCommunication(ctx context.Context, ticketData dto.TicketDeletedDTO) (emailIDs []uint64, err error)
}
