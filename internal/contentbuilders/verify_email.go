package contentbuilders

import (
	"fmt"
	"strconv"

	"github.com/DKhorkov/libs/security"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

type VerifyEmailContentBuilder struct {
	verifyEmailURLBase string
}

func NewVerifyEmailContentBuilder(verifyEmailURLBase string) *VerifyEmailContentBuilder {
	return &VerifyEmailContentBuilder{
		verifyEmailURLBase: verifyEmailURLBase,
	}
}

func (b *VerifyEmailContentBuilder) Subject() string {
	return "Подтверждение адреса электронной почты"
}

func (b *VerifyEmailContentBuilder) Body(user entities.User) string {
	link := fmt.Sprintf(
		"%s/%s",
		b.verifyEmailURLBase,
		security.RawEncode([]byte(strconv.FormatUint(user.ID, 10))),
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
