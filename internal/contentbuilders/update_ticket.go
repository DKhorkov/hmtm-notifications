package contentbuilders

import (
	"fmt"
	"strconv"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

func NewUpdateTicketContentBuilder(updateTicketURLBase string) *UpdateTicketContentBuilder {
	return &UpdateTicketContentBuilder{
		updateTicketURLBase: updateTicketURLBase,
	}
}

type UpdateTicketContentBuilder struct {
	updateTicketURLBase string
}

func (b *UpdateTicketContentBuilder) Subject(ticket entities.RawTicket) string {
	return fmt.Sprintf(
		"Заявка на создание игрушки %s была изменена",
		ticket.Name,
	)
}

func (b *UpdateTicketContentBuilder) Body(
	ticket entities.RawTicket,
	respondOwner entities.User,
) string {
	link := fmt.Sprintf(
		"%s/%s",
		b.updateTicketURLBase,
		strconv.FormatUint(ticket.ID, 10),
	)

	var priceInfo string
	if ticket.Price != nil {
		priceInfo = fmt.Sprintf(" на сумму <b>%.2f руб.</b>", *ticket.Price)
	}

	template := `<p>Добрый день, %s!</p>
<p>Заявка на создание игрушки <b>%s</b> (<i>%s</i>) в количестве <b>%d шт.</b>%s была изменена.</p>
<p>Для большей информации, пожалуйста, перейдите по <a href="%s">ссылке</a>.</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`
	return fmt.Sprintf(
		template,
		respondOwner.DisplayName,
		ticket.Name,
		ticket.Description,
		ticket.Quantity,
		priceInfo,
		link,
	)
}
