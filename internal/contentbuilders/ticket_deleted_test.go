package contentbuilders

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DKhorkov/libs/pointers"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

func TestTicketDeletedContentBuilder_Subject(t *testing.T) {
	builder := NewTicketDeletedContentBuilder("http://example.com/delete-ticket")

	testCases := []struct {
		name       string
		ticketData dto.TicketDeletedDTO
		expected   string
	}{
		{
			name: "basic ticket",
			ticketData: dto.TicketDeletedDTO{
				Name: "Teddy Bear",
			},
			expected: "Заявка на создание игрушки Teddy Bear была удалена",
		},
		{
			name: "ticket with special characters",
			ticketData: dto.TicketDeletedDTO{
				Name: "Super <Toy> & Fun",
			},
			expected: "Заявка на создание игрушки Super <Toy> & Fun была удалена",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.Subject(tc.ticketData)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestTicketDeletedContentBuilder_Body(t *testing.T) {
	builder := NewTicketDeletedContentBuilder("http://example.com/delete-ticket")

	testCases := []struct {
		name         string
		ticketData   dto.TicketDeletedDTO
		ticketOwner  entities.User
		respondOwner entities.User
		expected     string
	}{
		{
			name: "ticket with price",
			ticketData: dto.TicketDeletedDTO{
				Name:        "Teddy Bear",
				Description: "A soft teddy bear",
				Quantity:    5,
				Price:       pointers.New[float32](150.75),
			},
			ticketOwner: entities.User{
				ID:          1,
				DisplayName: "Alice",
			},
			respondOwner: entities.User{
				DisplayName: "Bob",
			},
			expected: `<p>Добрый день, Bob!</p>
<p>Пользователь <a href="http://example.com/delete-ticket/1">Alice</a> удалил заявку на создание игрушки <b>Teddy Bear</b> (<i>A soft teddy bear</i>) 
в количестве <b>5 шт.</b> на сумму <b>150.75 руб.</b></p>
<p>В связи с этим был удален ваш отклик на создание данной игрушки.</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`,
		},
		{
			name: "ticket without price",
			ticketData: dto.TicketDeletedDTO{
				Name:        "Wooden Car",
				Description: "A wooden toy car",
				Quantity:    1,
				Price:       nil,
			},
			ticketOwner: entities.User{
				ID:          2,
				DisplayName: "Charlie",
			},
			respondOwner: entities.User{
				DisplayName: "Dave",
			},
			expected: `<p>Добрый день, Dave!</p>
<p>Пользователь <a href="http://example.com/delete-ticket/2">Charlie</a> удалил заявку на создание игрушки <b>Wooden Car</b> (<i>A wooden toy car</i>) 
в количестве <b>1 шт.</b></p>
<p>В связи с этим был удален ваш отклик на создание данной игрушки.</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`,
		},
		{
			name: "ticket with special characters",
			ticketData: dto.TicketDeletedDTO{
				Name:        "Super <Toy>",
				Description: "Fun & Games",
				Quantity:    3,
				Price:       pointers.New[float32](99.99),
			},
			ticketOwner: entities.User{
				ID:          3,
				DisplayName: "Eve <Test>",
			},
			respondOwner: entities.User{
				DisplayName: "Frank",
			},
			expected: `<p>Добрый день, Frank!</p>
<p>Пользователь <a href="http://example.com/delete-ticket/3">Eve <Test></a> удалил заявку на создание игрушки <b>Super <Toy></b> (<i>Fun & Games</i>) 
в количестве <b>3 шт.</b> на сумму <b>99.99 руб.</b></p>
<p>В связи с этим был удален ваш отклик на создание данной игрушки.</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.Body(tc.ticketData, tc.ticketOwner, tc.respondOwner)
			require.Equal(t, tc.expected, result)
		})
	}
}
