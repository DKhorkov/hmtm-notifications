package entities

import (
	"time"
)

type Category struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

type ToyAttachment struct {
	ID        uint64    `json:"id"`
	ToyID     uint64    `json:"toy_id"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Toy struct {
	ID          uint64          `json:"id"`
	MasterID    uint64          `json:"master_id"`
	CategoryID  uint32          `json:"category_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       float32         `json:"price"`
	Quantity    uint32          `json:"quantity"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Tags        []Tag           `json:"tags,omitempty"`
	Attachments []ToyAttachment `json:"attachments,omitempty"`
}

type Master struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Info      *string   `json:"info,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
