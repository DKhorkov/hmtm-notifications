package contentbuilders

import (
	"fmt"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

func NewForgetPasswordContentBuilder(forgetPasswordURLBase string) *ForgetPasswordContentBuilder {
	return &ForgetPasswordContentBuilder{
		forgetPasswordURLBase: forgetPasswordURLBase,
	}
}

type ForgetPasswordContentBuilder struct {
	forgetPasswordURLBase string
}

func (b *ForgetPasswordContentBuilder) Subject() string {
	return "Восстановление пароля от аккаунта"
}

func (b *ForgetPasswordContentBuilder) Body(user entities.User, newPassword string) string {
	template := `<p>Добрый день, %s!</p>
<p>Ваш новый пароль: <b><i>%s</i></b>.</p>
<p>Пожалуйста, перейдите по <a href="%s">ссылке</a>, чтобы сменить пароль!</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`
	return fmt.Sprintf(
		template,
		user.DisplayName,
		newPassword,
		b.forgetPasswordURLBase,
	)
}
