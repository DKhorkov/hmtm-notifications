package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
	"github.com/DKhorkov/libs/logging"
)

func NewCommonSsoService(ssoRepository interfaces.SsoRepository, logger *slog.Logger) *CommonSsoService {
	return &CommonSsoService{
		ssoRepository: ssoRepository,
		logger:        logger,
	}
}

type CommonSsoService struct {
	ssoRepository interfaces.SsoRepository
	logger        *slog.Logger
}

func (service *CommonSsoService) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	users, err := service.ssoRepository.GetAllUsers(ctx)
	if err != nil {
		logging.LogErrorContext(ctx, service.logger, "Error occurred while trying to get all Users", err)
	}

	return users, err
}

func (service *CommonSsoService) GetUserByID(ctx context.Context, id uint64) (*entities.User, error) {
	user, err := service.ssoRepository.GetUserByID(ctx, id)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get User with ID=%d", id),
			err,
		)
	}

	return user, err
}
