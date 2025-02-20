package interfaces

import "github.com/DKhorkov/hmtm-notifications/internal/entities"

type ContentBuilders struct {
	VerifyEmail    VerifyEmailContentBuilder
	ForgetPassword ForgetPasswordContentBuilder
}

type VerifyEmailContentBuilder interface {
	Subject() string
	Body(user entities.User) string
}

type ForgetPasswordContentBuilder interface {
	Subject() string
	Body(user entities.User, newPassword string) string
}
