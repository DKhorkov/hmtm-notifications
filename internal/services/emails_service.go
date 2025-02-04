package services

import (
	"context"
	"log/slog"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func NewCommonEmailsService(
	emailsRepository interfaces.EmailsRepository,
	logger *slog.Logger,
) *CommonEmailsService {
	return &CommonEmailsService{
		emailsRepository: emailsRepository,
		logger:           logger,
	}
}

type CommonEmailsService struct {
	emailsRepository interfaces.EmailsRepository
	logger           *slog.Logger
}

func (service *CommonEmailsService) GetUserCommunications(
	ctx context.Context,
	userID uint64,
) ([]entities.Email, error) {
	return service.emailsRepository.GetUserCommunications(ctx, userID)
}

func (service *CommonEmailsService) SaveCommunication(ctx context.Context, email entities.Email) (uint64, error) {
	return service.emailsRepository.SaveCommunication(ctx, email)
}
