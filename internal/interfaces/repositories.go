package interfaces

import (
	"context"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

//go:generate mockgen -source=repositories.go -destination=../../mocks/repositories/emails_repository.go -exclude_interfaces=ToysRepository,SsoRepository,TicketsRepository -package=mockrepositories
type EmailsRepository interface {
	GetUserCommunications(ctx context.Context, userID uint64, pagination *entities.Pagination) ([]entities.Email, error)
	CountUserCommunications(ctx context.Context, userID uint64) (uint64, error)
	SaveCommunication(ctx context.Context, email entities.Email) (communicationID uint64, err error)
}

//go:generate mockgen -source=repositories.go -destination=../../mocks/repositories/sso_repository.go -exclude_interfaces=ToysRepository,EmailsRepository,TicketsRepository -package=mockrepositories
type SsoRepository interface {
	GetUserByID(ctx context.Context, id uint64) (*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
}

//go:generate mockgen -source=repositories.go -destination=../../mocks/repositories/tickets_repository.go -exclude_interfaces=ToysRepository,EmailsRepository,SsoRepository -package=mockrepositories
type TicketsRepository interface {
	GetTicketByID(ctx context.Context, id uint64) (*entities.RawTicket, error)
	GetAllTickets(ctx context.Context) ([]entities.RawTicket, error)
	GetUserTickets(ctx context.Context, userID uint64) ([]entities.RawTicket, error)
	GetRespondByID(ctx context.Context, id uint64) (*entities.Respond, error)
	GetTicketResponds(ctx context.Context, ticketID uint64) ([]entities.Respond, error)
	GetUserResponds(ctx context.Context, userID uint64) ([]entities.Respond, error)
}

//go:generate mockgen -source=repositories.go -destination=../../mocks/repositories/toys_repository.go -exclude_interfaces=TicketsRepository,EmailsRepository,SsoRepository -package=mockrepositories
type ToysRepository interface {
	GetAllToys(ctx context.Context) ([]entities.Toy, error)
	GetToyByID(ctx context.Context, id uint64) (*entities.Toy, error)
	GetMasterToys(ctx context.Context, masterID uint64) ([]entities.Toy, error)
	GetUserToys(ctx context.Context, userID uint64) ([]entities.Toy, error)
	GetAllMasters(ctx context.Context) ([]entities.Master, error)
	GetMasterByID(ctx context.Context, id uint64) (*entities.Master, error)
	GetAllCategories(ctx context.Context) ([]entities.Category, error)
	GetCategoryByID(ctx context.Context, id uint32) (*entities.Category, error)
	GetAllTags(ctx context.Context) ([]entities.Tag, error)
	GetTagByID(ctx context.Context, id uint32) (*entities.Tag, error)
	GetMasterByUser(ctx context.Context, userID uint64) (*entities.Master, error)
}
