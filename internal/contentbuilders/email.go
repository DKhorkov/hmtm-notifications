package contentbuilders

import (
	"fmt"
	"strconv"

	"github.com/DKhorkov/libs/security"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

func NewCommonEmailContentBuilder(verifyEmailURLBase string) *CommonEmailContentBuilder {
	return &CommonEmailContentBuilder{
		verifyEmailURLBase: verifyEmailURLBase,
	}
}

type CommonEmailContentBuilder struct {
	verifyEmailURLBase string
}

func (b *CommonEmailContentBuilder) Subject() string {
	return "Подтверждение адреса электронной почты"
}

func (b *CommonEmailContentBuilder) Body(user entities.User) string {
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
