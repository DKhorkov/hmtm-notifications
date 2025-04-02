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
	ToyID     uint64    `json:"toyId"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Toy struct {
	ID          uint64          `json:"id"`
	MasterID    uint64          `json:"masterId"`
	CategoryID  uint32          `json:"categoryId"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       float32         `json:"price"`
	Quantity    uint32          `json:"quantity"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Tags        []Tag           `json:"tags,omitempty"`
	Attachments []ToyAttachment `json:"attachments,omitempty"`
}

type Master struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"userId"`
	Info      *string   `json:"info,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
