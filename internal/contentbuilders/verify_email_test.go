package contentbuilders

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
)

func TestVerifyEmailContentBuilder_Subject(t *testing.T) {
	builder := NewVerifyEmailContentBuilder("http://example.com/verify-email")

	testCases := []struct {
		name     string
		expected string
	}{
		{
			name:     "default subject",
			expected: "Подтверждение адреса электронной почты",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.Subject()
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestVerifyEmailContentBuilder_Body(t *testing.T) {
	builder := NewVerifyEmailContentBuilder("http://example.com/verify-email")

	testCases := []struct {
		name     string
		user     entities.User
		expected string
	}{
		{
			name: "basic user",
			user: entities.User{
				ID:          1,
				DisplayName: "Alice",
			},
			expected: `<p>Добрый день, Alice!</p>
<p>Пожалуйста, перейдите по <a href="http://example.com/verify-email/MQ">ссылке</a>, чтобы подтвердить адрес электронной почты!</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`,
		},
		{
			name: "user with special characters",
			user: entities.User{
				ID:          123,
				DisplayName: "Bob <Test>",
			},
			expected: `<p>Добрый день, Bob <Test>!</p>
<p>Пожалуйста, перейдите по <a href="http://example.com/verify-email/MTIz">ссылке</a>, чтобы подтвердить адрес электронной почты!</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`,
		},
		{
			name: "user with large ID",
			user: entities.User{
				ID:          987654321,
				DisplayName: "Charlie",
			},
			expected: `<p>Добрый день, Charlie!</p>
<p>Пожалуйста, перейдите по <a href="http://example.com/verify-email/OTg3NjU0MzIx">ссылке</a>, чтобы подтвердить адрес электронной почты!</p>
<p>С уважением,<br>
команда Handmade Toys Marketplace.</p>
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.Body(tc.user)
			require.Equal(t, tc.expected, result)
		})
	}
}
