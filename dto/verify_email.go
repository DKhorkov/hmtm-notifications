package dto

type VerifyEmailDTO struct {
	UserID uint64 `json:"user_id"`
}

type ForgetPasswordDTO struct {
	UserID      uint64 `json:"user_id"`
	NewPassword string `json:"new_password"`
}
