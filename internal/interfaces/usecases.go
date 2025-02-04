package interfaces

import (
	"context"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

type UseCases interface {
	GetUserEmailCommunications(ctx context.Context, userID uint64) ([]entities.Email, error)
	SendVerifyEmailCommunication(ctx context.Context, userID uint64) (emailID uint64, err error)
}
