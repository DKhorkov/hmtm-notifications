package interfaces

import (
	"context"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

type UseCases interface {
	GetUserEmailCommunications(ctx context.Context, userID uint64) ([]entities.Email, error)
	SendVerifyEmailCommunication(ctx context.Context, userID uint64) (emailID uint64, err error)
	SendForgetPasswordEmailCommunication(ctx context.Context, userID uint64, newPassword string) (emailID uint64, err error)
	SendUpdateTicketEmailCommunication(ctx context.Context, ticketID uint64) (emailIDs []uint64, err error)
	SendDeleteTicketEmailCommunication(ctx context.Context, ticketData dto.DeleteTicketDTO) (emailIDs []uint64, err error)
}
