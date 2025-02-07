package usecases

import (
	"context"
	"time"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func NewCommonUseCases(
	emailsService interfaces.EmailsService,
	ssoService interfaces.SsoService,
	contentBuilders interfaces.ContentBuilders,
	senders interfaces.Senders,
) *CommonUseCases {
	return &CommonUseCases{
		emailsService:   emailsService,
		ssoService:      ssoService,
		contentBuilders: contentBuilders,
		senders:         senders,
	}
}

type CommonUseCases struct {
	emailsService   interfaces.EmailsService
	ssoService      interfaces.SsoService
	contentBuilders interfaces.ContentBuilders
	senders         interfaces.Senders
}

func (useCases *CommonUseCases) GetUserEmailCommunications(
	ctx context.Context,
	userID uint64,
) ([]entities.Email, error) {
	return useCases.emailsService.GetUserCommunications(ctx, userID)
}

func (useCases *CommonUseCases) SendVerifyEmailCommunication(
	ctx context.Context,
	userID uint64,
) (uint64, error) {
	user, err := useCases.ssoService.GetUserByID(ctx, userID)
	if err != nil {
		return 0, err
	}

	if err = useCases.senders.Email.Send(
		ctx,
		useCases.contentBuilders.Email.Subject(),
		useCases.contentBuilders.Email.Body(*user),
		[]string{user.Email},
	); err != nil {
		return 0, err
	}

	emailCommunication := entities.Email{
		UserID:  user.ID,
		Email:   user.Email,
		Content: useCases.contentBuilders.Email.Body(*user),
		SentAt:  time.Now().UTC(),
	}

	return useCases.emailsService.SaveCommunication(ctx, emailCommunication)
}
