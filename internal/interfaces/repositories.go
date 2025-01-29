package interfaces

import (
	"context"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

//go:generate mockgen -source=repositories.go -destination=../../mocks/repositories/emails_repository.go -package=mockrepositories
type EmailsRepository interface {
	GetUserEmailCommunications(ctx context.Context, userID uint64) ([]entities.Email, error)
}
