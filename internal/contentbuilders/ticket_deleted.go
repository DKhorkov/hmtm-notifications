package contentbuilders

import (
	"fmt"
	"strconv"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

type TicketDeletedContentBuilder struct {
	ticketDeleteURLBase string
}

func NewTicketDeletedContentBuilder(ticketDeleteURLBase string) *TicketDeletedContentBuilder {
	return &TicketDeletedContentBuilder{
		ticketDeleteURLBase: ticketDeleteURLBase,
	}
}

func (b *TicketDeletedContentBuilder) Subject(ticketData dto.TicketDeletedDTO) string {
	return fmt.Sprintf(
		"Заявка на создание игрушки %s была удалена",
		ticketData.Name,
	)
}

func (b *TicketDeletedContentBuilder) Body(
	ticketData dto.TicketDeletedDTO,
	ticketOwner entities.User,
	respondOwner entities.User,
) string {
	link := fmt.Sprintf(
		"%s/%s",
		b.ticketDeleteURLBase,
		strconv.FormatUint(ticketOwner.ID, 10),
	)

	var priceInfo string
	if ticketData.Price != nil {
		priceInfo = fmt.Sprintf(" на сумму <b>%.2f руб.</b>", *ticketData.Price)
	}

	template := `<p>Добрый день, %s!</p>
<p>Пользователь <a href="%s">%s</a> удалил заявку на создание игрушки <b>%s</b> (<i>%s</i>) 
в количестве <b>%d шт.</b>%s</p>
<p>В связи с этим был удален ваш отклик на создание данной игрушки.</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`

	return fmt.Sprintf(
		template,
		respondOwner.DisplayName,
		link,
		ticketOwner.DisplayName,
		ticketData.Name,
		ticketData.Description,
		ticketData.Quantity,
		priceInfo,
	)
}
