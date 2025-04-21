package dto

type TicketDeletedDTO struct {
	TicketOwnerID       uint64   `json:"ticketOwnerId"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	Price               *float32 `json:"price,omitempty"`
	Quantity            uint32   `json:"quantity"`
	RespondedMastersIDs []uint64 `json:"respondedMastersIds"`
}
