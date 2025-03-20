package usecases

import (
	"context"
	"time"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func New(
	emailsService interfaces.EmailsService,
	ssoService interfaces.SsoService,
	toysService interfaces.ToysService,
	ticketsService interfaces.TicketsService,
	contentBuilders interfaces.ContentBuilders,
	senders interfaces.Senders,
) *UseCases {
	return &UseCases{
		emailsService:   emailsService,
		ssoService:      ssoService,
		toysService:     toysService,
		ticketsService:  ticketsService,
		contentBuilders: contentBuilders,
		senders:         senders,
	}
}

type UseCases struct {
	emailsService   interfaces.EmailsService
	ssoService      interfaces.SsoService
	toysService     interfaces.ToysService
	ticketsService  interfaces.TicketsService
	contentBuilders interfaces.ContentBuilders
	senders         interfaces.Senders
}

func (useCases *UseCases) GetUserEmailCommunications(
	ctx context.Context,
	userID uint64,
) ([]entities.Email, error) {
	return useCases.emailsService.GetUserCommunications(ctx, userID)
}

func (useCases *UseCases) SendVerifyEmailCommunication(
	ctx context.Context,
	userID uint64,
) (uint64, error) {
	user, err := useCases.ssoService.GetUserByID(ctx, userID)
	if err != nil {
		return 0, err
	}

	if err = useCases.senders.Email.Send(
		ctx,
		useCases.contentBuilders.VerifyEmail.Subject(),
		useCases.contentBuilders.VerifyEmail.Body(*user),
		[]string{user.Email},
	); err != nil {
		return 0, err
	}

	emailCommunication := entities.Email{
		UserID:  user.ID,
		Email:   user.Email,
		Content: useCases.contentBuilders.VerifyEmail.Body(*user),
		SentAt:  time.Now().UTC(),
	}

	return useCases.emailsService.SaveCommunication(ctx, emailCommunication)
}

func (useCases *UseCases) SendForgetPasswordEmailCommunication(
	ctx context.Context,
	userID uint64,
	newPassword string,
) (uint64, error) {
	user, err := useCases.ssoService.GetUserByID(ctx, userID)
	if err != nil {
		return 0, err
	}

	if err = useCases.senders.Email.Send(
		ctx,
		useCases.contentBuilders.ForgetPassword.Subject(),
		useCases.contentBuilders.ForgetPassword.Body(*user, newPassword),
		[]string{user.Email},
	); err != nil {
		return 0, err
	}

	emailCommunication := entities.Email{
		UserID:  user.ID,
		Email:   user.Email,
		Content: useCases.contentBuilders.ForgetPassword.Body(*user, newPassword),
		SentAt:  time.Now().UTC(),
	}

	return useCases.emailsService.SaveCommunication(ctx, emailCommunication)
}

func (useCases *UseCases) SendUpdateTicketEmailCommunication(ctx context.Context, ticketID uint64) ([]uint64, error) {
	rawTicket, err := useCases.ticketsService.GetTicketByID(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	responds, err := useCases.ticketsService.GetTicketResponds(ctx, rawTicket.ID)
	if err != nil {
		return nil, err
	}

	var emailIDs []uint64
	for _, respond := range responds {
		master, err := useCases.toysService.GetMasterByID(ctx, respond.MasterID)
		if err != nil {
			return nil, err
		}

		respondOwner, err := useCases.ssoService.GetUserByID(ctx, master.UserID)
		if err != nil {
			return nil, err
		}

		if err = useCases.senders.Email.Send(
			ctx,
			useCases.contentBuilders.UpdateTicket.Subject(*rawTicket),
			useCases.contentBuilders.UpdateTicket.Body(*rawTicket, *respondOwner),
			[]string{respondOwner.Email},
		); err != nil {
			return nil, err
		}

		emailCommunication := entities.Email{
			UserID:  respondOwner.ID,
			Email:   respondOwner.Email,
			Content: useCases.contentBuilders.UpdateTicket.Body(*rawTicket, *respondOwner),
			SentAt:  time.Now().UTC(),
		}

		emailID, err := useCases.emailsService.SaveCommunication(ctx, emailCommunication)
		if err != nil {
			return nil, err
		}

		emailIDs = append(emailIDs, emailID)
	}

	return emailIDs, nil
}

func (useCases *UseCases) SendDeleteTicketEmailCommunication(
	ctx context.Context,
	ticketData dto.DeleteTicketDTO,
) ([]uint64, error) {
	ticketOwner, err := useCases.ssoService.GetUserByID(ctx, ticketData.TicketOwnerID)
	if err != nil {
		return nil, err
	}

	var emailIDs []uint64
	for _, masterID := range ticketData.RespondedMastersIDs {
		master, err := useCases.toysService.GetMasterByID(ctx, masterID)
		if err != nil {
			return nil, err
		}

		respondOwner, err := useCases.ssoService.GetUserByID(ctx, master.UserID)
		if err != nil {
			return nil, err
		}

		if err = useCases.senders.Email.Send(
			ctx,
			useCases.contentBuilders.DeleteTicket.Subject(ticketData),
			useCases.contentBuilders.DeleteTicket.Body(ticketData, *ticketOwner, *respondOwner),
			[]string{respondOwner.Email},
		); err != nil {
			return nil, err
		}

		emailCommunication := entities.Email{
			UserID:  respondOwner.ID,
			Email:   respondOwner.Email,
			Content: useCases.contentBuilders.DeleteTicket.Body(ticketData, *ticketOwner, *respondOwner),
			SentAt:  time.Now().UTC(),
		}

		emailID, err := useCases.emailsService.SaveCommunication(ctx, emailCommunication)
		if err != nil {
			return nil, err
		}

		emailIDs = append(emailIDs, emailID)
	}

	return emailIDs, nil
}
