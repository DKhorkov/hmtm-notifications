package services

import (
	"context"

	"github.com/DKhorkov/libs/logging"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

type EmailsService struct {
	emailsRepository interfaces.EmailsRepository
	logger           logging.Logger
}

func NewEmailsService(
	emailsRepository interfaces.EmailsRepository,
	logger logging.Logger,
) *EmailsService {
	return &EmailsService{
		emailsRepository: emailsRepository,
		logger:           logger,
	}
}

func (service *EmailsService) GetUserCommunications(
	ctx context.Context,
	userID uint64,
) ([]entities.Email, error) {
	return service.emailsRepository.GetUserCommunications(ctx, userID)
}

func (service *EmailsService) SaveCommunication(
	ctx context.Context,
	email entities.Email,
) (uint64, error) {
	return service.emailsRepository.SaveCommunication(ctx, email)
}
