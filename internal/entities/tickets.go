package entities

import "time"

type RawTicket struct {
	ID          uint64             `json:"id"`
	UserID      uint64             `json:"userId"`
	CategoryID  uint32             `json:"categoryId"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       *float32           `json:"price,omitempty"`
	Quantity    uint32             `json:"quantity"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	TagIDs      []uint32           `json:"tagIds"`
	Attachments []TicketAttachment `json:"attachments,omitempty"`
}

type Ticket struct {
	ID          uint64             `json:"id"`
	UserID      uint64             `json:"userId"`
	CategoryID  uint32             `json:"categoryId"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       *float32           `json:"price,omitempty"`
	Quantity    uint32             `json:"quantity"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	Tags        []Tag              `json:"tags,omitempty"`
	Attachments []TicketAttachment `json:"attachments,omitempty"`
}

type TicketAttachment struct {
	ID        uint64    `json:"id"`
	TicketID  uint64    `json:"ticketId"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Respond struct {
	ID        uint64    `json:"id"`
	TicketID  uint64    `json:"ticketId"`
	MasterID  uint64    `json:"masterId"`
	Price     float32   `json:"price"`
	Comment   *string   `json:"comment,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
