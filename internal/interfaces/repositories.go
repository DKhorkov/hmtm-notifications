package interfaces

import (
	"context"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

//go:generate mockgen -source=repositories.go -destination=../../mocks/repositories/emails_repository.go -exclude_interfaces=SsoRepository -package=mockrepositories
type EmailsRepository interface {
	GetUserCommunications(ctx context.Context, userID uint64) ([]entities.Email, error)
	SaveCommunication(ctx context.Context, email entities.Email) (communicationID uint64, err error)
}

//go:generate mockgen -source=repositories.go -destination=../../mocks/repositories/sso_repository.go -exclude_interfaces=EmailsRepository -package=mockrepositories
type SsoRepository interface {
	GetUserByID(ctx context.Context, id uint64) (*entities.User, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
}
