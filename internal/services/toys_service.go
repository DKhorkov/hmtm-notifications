package services

import (
	"context"
	"fmt"

	"github.com/DKhorkov/libs/logging"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func NewToysService(toysRepository interfaces.ToysRepository, logger logging.Logger) *ToysService {
	return &ToysService{
		toysRepository: toysRepository,
		logger:         logger,
	}
}

type ToysService struct {
	toysRepository interfaces.ToysRepository
	logger         logging.Logger
}

func (service *ToysService) GetAllToys(ctx context.Context) ([]entities.Toy, error) {
	toys, err := service.toysRepository.GetAllToys(ctx)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			"Error occurred while trying to get all Toys",
			err,
		)
	}

	return toys, err
}

func (service *ToysService) GetMasterToys(
	ctx context.Context,
	masterID uint64,
) ([]entities.Toy, error) {
	toys, err := service.toysRepository.GetMasterToys(ctx, masterID)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf(
				"Error occurred while trying to get all Toys for Master with ID=%d",
				masterID,
			),
			err,
		)
	}

	return toys, err
}

func (service *ToysService) GetUserToys(
	ctx context.Context,
	userID uint64,
) ([]entities.Toy, error) {
	toys, err := service.toysRepository.GetUserToys(ctx, userID)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get all Toys for User with ID=%d", userID),
			err,
		)
	}

	return toys, err
}

func (service *ToysService) GetToyByID(ctx context.Context, id uint64) (*entities.Toy, error) {
	toy, err := service.toysRepository.GetToyByID(ctx, id)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get Toy with ID=%d", id),
			err,
		)
	}

	return toy, err
}

func (service *ToysService) GetAllMasters(ctx context.Context) ([]entities.Master, error) {
	masters, err := service.toysRepository.GetAllMasters(ctx)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			"Error occurred while trying to get all Masters",
			err,
		)
	}

	return masters, err
}

func (service *ToysService) GetMasterByID(
	ctx context.Context,
	id uint64,
) (*entities.Master, error) {
	master, err := service.toysRepository.GetMasterByID(ctx, id)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get Master with ID=%d", id),
			err,
		)
	}

	return master, err
}

func (service *ToysService) GetAllCategories(ctx context.Context) ([]entities.Category, error) {
	categories, err := service.toysRepository.GetAllCategories(ctx)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			"Error occurred while trying to get all Categories",
			err,
		)
	}

	return categories, err
}

func (service *ToysService) GetCategoryByID(
	ctx context.Context,
	id uint32,
) (*entities.Category, error) {
	category, err := service.toysRepository.GetCategoryByID(ctx, id)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get Category with ID=%d", id),
			err,
		)
	}

	return category, err
}

func (service *ToysService) GetAllTags(ctx context.Context) ([]entities.Tag, error) {
	tags, err := service.toysRepository.GetAllTags(ctx)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			"Error occurred while trying to get all Tags",
			err,
		)
	}

	return tags, err
}

func (service *ToysService) GetTagByID(ctx context.Context, id uint32) (*entities.Tag, error) {
	tag, err := service.toysRepository.GetTagByID(ctx, id)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get Tag with ID=%d", id),
			err,
		)
	}

	return tag, err
}

func (service *ToysService) GetMasterByUser(
	ctx context.Context,
	userID uint64,
) (*entities.Master, error) {
	master, err := service.toysRepository.GetMasterByUser(ctx, userID)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get Master with userID=%d", userID),
			err,
		)
	}

	return master, err
}
