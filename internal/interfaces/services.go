package interfaces

type EmailsService interface {
	EmailsRepository
}

type SsoService interface {
	SsoRepository
}

type TicketsService interface {
	TicketsRepository
}

type ToysService interface {
	ToysRepository
}
