package services

import (
	"context"
	"fmt"

	"github.com/DKhorkov/libs/logging"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func NewSsoService(ssoRepository interfaces.SsoRepository, logger logging.Logger) *SsoService {
	return &SsoService{
		ssoRepository: ssoRepository,
		logger:        logger,
	}
}

type SsoService struct {
	ssoRepository interfaces.SsoRepository
	logger        logging.Logger
}

func (service *SsoService) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	users, err := service.ssoRepository.GetAllUsers(ctx)
	if err != nil {
		logging.LogErrorContext(ctx, service.logger, "Error occurred while trying to get all Users", err)
	}

	return users, err
}

func (service *SsoService) GetUserByID(ctx context.Context, id uint64) (*entities.User, error) {
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
