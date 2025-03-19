package entities

import "time"

type RawTicket struct {
	ID          uint64             `json:"id"`
	UserID      uint64             `json:"user_id"`
	CategoryID  uint32             `json:"category_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       *float32           `json:"price,omitempty"`
	Quantity    uint32             `json:"quantity"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	TagIDs      []uint32           `json:"tag_ids"`
	Attachments []TicketAttachment `json:"attachments,omitempty"`
}

type Ticket struct {
	ID          uint64             `json:"id"`
	UserID      uint64             `json:"user_id"`
	CategoryID  uint32             `json:"category_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       *float32           `json:"price,omitempty"`
	Quantity    uint32             `json:"quantity"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Tags        []Tag              `json:"tags,omitempty"`
	Attachments []TicketAttachment `json:"attachments,omitempty"`
}

type TicketAttachment struct {
	ID        uint64    `json:"id"`
	TicketID  uint64    `json:"ticket_id"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Respond struct {
	ID        uint64    `json:"id"`
	TicketID  uint64    `json:"ticket_id"`
	MasterID  uint64    `json:"master_id"`
	Price     float32   `json:"price"`
	Comment   *string   `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
