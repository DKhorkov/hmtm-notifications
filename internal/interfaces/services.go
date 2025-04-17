package interfaces

//go:generate mockgen -source=services.go -destination=../../mocks/services/email_service.go -package=mockservices -exclude_interfaces=ToysService,TicketsService,SsoService
type EmailsService interface {
	EmailsRepository
}

//go:generate mockgen -source=services.go -destination=../../mocks/services/sso_service.go -package=mockservices -exclude_interfaces=ToysService,EmailsService,TicketsService
type SsoService interface {
	SsoRepository
}

//go:generate mockgen -source=services.go -destination=../../mocks/services/tickets_service.go -package=mockservices -exclude_interfaces=ToysService,EmailsService,SsoService
type TicketsService interface {
	TicketsRepository
}

//go:generate mockgen -source=services.go -destination=../../mocks/services/toys_service.go -package=mockservices -exclude_interfaces=SsoService,EmailsService,TicketsService
type ToysService interface {
	ToysRepository
}
