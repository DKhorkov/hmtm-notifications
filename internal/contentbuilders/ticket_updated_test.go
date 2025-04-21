package contentbuilders

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DKhorkov/libs/pointers"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

func TestTicketUpdatedContentBuilder_Subject(t *testing.T) {
	builder := NewTicketUpdatedContentBuilder("http://example.com/update-ticket")

	testCases := []struct {
		name     string
		ticket   entities.RawTicket
		expected string
	}{
		{
			name: "basic ticket",
			ticket: entities.RawTicket{
				Name: "Teddy Bear",
			},
			expected: "Заявка на создание игрушки Teddy Bear была изменена",
		},
		{
			name: "ticket with special characters",
			ticket: entities.RawTicket{
				Name: "Super <Toy> & Fun",
			},
			expected: "Заявка на создание игрушки Super <Toy> & Fun была изменена",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.Subject(tc.ticket)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestTicketUpdatedContentBuilder_Body(t *testing.T) {
	builder := NewTicketUpdatedContentBuilder("http://example.com/update-ticket")

	testCases := []struct {
		name         string
		ticket       entities.RawTicket
		respondOwner entities.User
		expected     string
	}{
		{
			name: "ticket with price",
			ticket: entities.RawTicket{
				ID:          1,
				Name:        "Teddy Bear",
				Description: "A soft teddy bear",
				Quantity:    5,
				Price:       pointers.New[float32](150.75),
			},
			respondOwner: entities.User{
				DisplayName: "Bob",
			},
			expected: `<p>Добрый день, Bob!</p>
<p>Заявка на создание игрушки <b>Teddy Bear</b> (<i>A soft teddy bear</i>) в количестве <b>5 шт.</b> на сумму <b>150.75 руб.</b> была изменена.</p>
<p>Для большей информации, пожалуйста, перейдите по <a href="http://example.com/update-ticket/1">ссылке</a>.</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`,
		},
		{
			name: "ticket without price",
			ticket: entities.RawTicket{
				ID:          2,
				Name:        "Wooden Car",
				Description: "A wooden toy car",
				Quantity:    1,
				Price:       nil,
			},
			respondOwner: entities.User{
				DisplayName: "Dave",
			},
			expected: `<p>Добрый день, Dave!</p>
<p>Заявка на создание игрушки <b>Wooden Car</b> (<i>A wooden toy car</i>) в количестве <b>1 шт.</b> была изменена.</p>
<p>Для большей информации, пожалуйста, перейдите по <a href="http://example.com/update-ticket/2">ссылке</a>.</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`,
		},
		{
			name: "ticket with special characters",
			ticket: entities.RawTicket{
				ID:          3,
				Name:        "Super <Toy>",
				Description: "Fun & Games",
				Quantity:    3,
				Price:       pointers.New[float32](99.99),
			},
			respondOwner: entities.User{
				DisplayName: "Frank",
			},
			expected: `<p>Добрый день, Frank!</p>
<p>Заявка на создание игрушки <b>Super <Toy></b> (<i>Fun & Games</i>) в количестве <b>3 шт.</b> на сумму <b>99.99 руб.</b> была изменена.</p>
<p>Для большей информации, пожалуйста, перейдите по <a href="http://example.com/update-ticket/3">ссылке</a>.</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.Body(tc.ticket, tc.respondOwner)
			require.Equal(t, tc.expected, result)
		})
	}
}
