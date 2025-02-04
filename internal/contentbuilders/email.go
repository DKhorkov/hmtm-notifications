package contentbuilders

import (
	"fmt"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

func NewCommonEmailContentBuilder() *CommonEmailContentBuilder {
	return &CommonEmailContentBuilder{}
}

type CommonEmailContentBuilder struct{}

func (b *CommonEmailContentBuilder) Subject() string {
	return "Подтверждение адреса электронной почты"
}

func (b *CommonEmailContentBuilder) Body(user entities.User) string {
	template := `<p>Добрый день, %s!</p>
<p>Пожалуйста, перейдите по ссылке, чтобы подтвердить адрес электронной почты!</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`
	return fmt.Sprintf(
		template,
		user.DisplayName,
	)
}
