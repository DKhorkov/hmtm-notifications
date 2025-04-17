package interfaces

import (
	"github.com/DKhorkov/hmtm-sso/api/protobuf/generated/go/sso"
	"github.com/DKhorkov/hmtm-tickets/api/protobuf/generated/go/tickets"
	"github.com/DKhorkov/hmtm-toys/api/protobuf/generated/go/toys"
)

//go:generate mockgen -source=clients.go -destination=../../mocks/clients/sso_client.go -package=mockclients -exclude_interfaces=ToysClient,TicketsClient
type SsoClient interface {
	sso.UsersServiceClient
}

//go:generate mockgen -source=clients.go -destination=../../mocks/clients/tickets_client.go -package=mockclients -exclude_interfaces=ToysClient,SsoClient
type TicketsClient interface {
	tickets.TicketsServiceClient
	tickets.RespondsServiceClient
}

//go:generate mockgen -source=clients.go -destination=../../mocks/clients/toys_client.go -package=mockclients -exclude_interfaces=SsoClient,TicketsClient
type ToysClient interface {
	toys.CategoriesServiceClient
	toys.ToysServiceClient
	toys.TagsServiceClient
	toys.MastersServiceClient
}
