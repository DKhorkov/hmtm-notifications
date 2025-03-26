package services

import (
	"context"
	"fmt"

	"github.com/DKhorkov/libs/logging"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func NewTicketsService(
	ticketsRepository interfaces.TicketsRepository,
	logger logging.Logger,
) *TicketsService {
	return &TicketsService{
		ticketsRepository: ticketsRepository,
		logger:            logger,
	}
}

type TicketsService struct {
	ticketsRepository interfaces.TicketsRepository
	logger            logging.Logger
}

func (service *TicketsService) GetTicketByID(
	ctx context.Context,
	id uint64,
) (*entities.RawTicket, error) {
	ticket, err := service.ticketsRepository.GetTicketByID(ctx, id)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get Ticket with ID=%d", id),
			err,
		)
	}

	return ticket, err
}

func (service *TicketsService) GetAllTickets(ctx context.Context) ([]entities.RawTicket, error) {
	tickets, err := service.ticketsRepository.GetAllTickets(ctx)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			"Error occurred while trying to get all Tickets",
			err,
		)
	}

	return tickets, err
}

func (service *TicketsService) GetUserTickets(
	ctx context.Context,
	userID uint64,
) ([]entities.RawTicket, error) {
	tickets, err := service.ticketsRepository.GetUserTickets(ctx, userID)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get Tickets for User with ID=%d", userID),
			err,
		)
	}

	return tickets, err
}

func (service *TicketsService) GetRespondByID(
	ctx context.Context,
	id uint64,
) (*entities.Respond, error) {
	respond, err := service.ticketsRepository.GetRespondByID(ctx, id)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get Respond with ID=%d", id),
			err,
		)
	}

	return respond, err
}

func (service *TicketsService) GetTicketResponds(
	ctx context.Context,
	ticketID uint64,
) ([]entities.Respond, error) {
	responds, err := service.ticketsRepository.GetTicketResponds(ctx, ticketID)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf(
				"Error occurred while trying to get Responds for Ticket with ID=%d",
				ticketID,
			),
			err,
		)
	}

	return responds, err
}

func (service *TicketsService) GetUserResponds(
	ctx context.Context,
	userID uint64,
) ([]entities.Respond, error) {
	responds, err := service.ticketsRepository.GetUserResponds(ctx, userID)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			service.logger,
			fmt.Sprintf("Error occurred while trying to get Responds for User with ID=%d", userID),
			err,
		)
	}

	return responds, err
}
