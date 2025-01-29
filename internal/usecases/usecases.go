package usecases

import (
	"context"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func NewCommonUseCases(
	emailsService interfaces.EmailsService,
) *CommonUseCases {
	return &CommonUseCases{
		emailsService: emailsService,
	}
}

type CommonUseCases struct {
	emailsService interfaces.EmailsService
}

func (useCases *CommonUseCases) GetUserEmailCommunications(
	ctx context.Context,
	userID uint64,
) ([]entities.Email, error) {
	return useCases.emailsService.GetUserEmailCommunications(ctx, userID)
}
