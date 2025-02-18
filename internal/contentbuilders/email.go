package contentbuilders

import (
	"fmt"
	"strconv"

	"github.com/DKhorkov/libs/security"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

func NewEmailContentBuilder(verifyEmailURLBase string) *EmailContentBuilder {
	return &EmailContentBuilder{
		verifyEmailURLBase: verifyEmailURLBase,
	}
}

type EmailContentBuilder struct {
	verifyEmailURLBase string
}

func (b *EmailContentBuilder) Subject() string {
	return "Подтверждение адреса электронной почты"
}

func (b *EmailContentBuilder) Body(user entities.User) string {
	link := fmt.Sprintf(
		"%s/%s",
		b.verifyEmailURLBase,
		security.Encode([]byte(strconv.FormatUint(user.ID, 10))),
	)

	template := `<p>Добрый день, %s!</p>
<p>Пожалуйста, перейдите по <a href="%s">ссылке</a>, чтобы подтвердить адрес электронной почты!</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`
	return fmt.Sprintf(
		template,
		user.DisplayName,
		link,
	)
}
