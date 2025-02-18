package services

import (
	"context"
	"log/slog"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func NewEmailsService(
	emailsRepository interfaces.EmailsRepository,
	logger *slog.Logger,
) *EmailsService {
	return &EmailsService{
		emailsRepository: emailsRepository,
		logger:           logger,
	}
}

type EmailsService struct {
	emailsRepository interfaces.EmailsRepository
	logger           *slog.Logger
}

func (service *EmailsService) GetUserCommunications(
	ctx context.Context,
	userID uint64,
) ([]entities.Email, error) {
	return service.emailsRepository.GetUserCommunications(ctx, userID)
}

func (service *EmailsService) SaveCommunication(ctx context.Context, email entities.Email) (uint64, error) {
	return service.emailsRepository.SaveCommunication(ctx, email)
}
