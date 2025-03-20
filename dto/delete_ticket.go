package dto

type DeleteTicketDTO struct {
	TicketOwnerID       uint64   `json:"ticket_owner_id"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	Price               *float32 `json:"price,omitempty"`
	Quantity            uint32   `json:"quantity"`
	RespondedMastersIDs []uint64 `json:"responded_masters_ids"`
}
