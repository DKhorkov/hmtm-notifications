package usecases

import (
	"context"
	"time"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func New(
	emailsService interfaces.EmailsService,
	ssoService interfaces.SsoService,
	contentBuilders interfaces.ContentBuilders,
	senders interfaces.Senders,
) *UseCases {
	return &UseCases{
		emailsService:   emailsService,
		ssoService:      ssoService,
		contentBuilders: contentBuilders,
		senders:         senders,
	}
}

type UseCases struct {
	emailsService   interfaces.EmailsService
	ssoService      interfaces.SsoService
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

func (useCases *UseCases) SendForgetPasswordCommunication(
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
